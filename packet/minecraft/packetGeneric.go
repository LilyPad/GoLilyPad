package minecraft

import "bytes"
import "encoding/binary"
import "io"
import "io/ioutil"
import "github.com/LilyPad/GoLilyPad/packet"

type PacketGeneric struct {
	id int
	Bytes []byte
}

func (this *PacketGeneric) SwapEntities(a int32, b int32, clientServer bool) {
	if a == b {
		return
	}
	if this.id == PACKET_CLIENT_SPAWN_OBJECT && clientServer {
		buffer := bytes.NewBuffer(this.Bytes)
		bufferUtil := make([]byte, packet.UTIL_BUFFER_LENGTH)
		_, err := packet.ReadVarInt(buffer, bufferUtil)
		varIntLength := len(this.Bytes) - buffer.Len()
		if err == nil && len(this.Bytes) > varIntLength + 19 {
			objectType := this.Bytes[varIntLength]
			if objectType == 60 || objectType == 61 || objectType == 62 || objectType == 63 || objectType == 64 || objectType == 65 || objectType == 66 || objectType == 90 {
				id := int32(binary.BigEndian.Uint32(this.Bytes[varIntLength+15:varIntLength+19]))
				if id == a {
					binary.BigEndian.PutUint32(this.Bytes[varIntLength+15:varIntLength+19], uint32(b))
				} else if id == b {
					binary.BigEndian.PutUint32(this.Bytes[varIntLength+15:varIntLength+19], uint32(a))
				}
			}
		}
	} else if this.id == PACKET_CLIENT_DESTROY_ENTITIES && clientServer {
		// TODO
	}
	this.SwapEntitiesInt(a, b, clientServer)
	this.SwapEntitiesVarInt(a, b, clientServer)
}

func (this *PacketGeneric) SwapEntitiesInt(a int32, b int32, clientServer bool) {
	var positions [][]int
	if clientServer {
		positions = PlayPacketClientEntityIntPositions
	} else {
		positions = PlayPacketServerEntityIntPositions
	}
	if this.id < 0 {
		return
	}
	if this.id >= len(positions) {
		return
	}
	idPositions := positions[this.id]
	if idPositions == nil {
		return
	}
	var id int32
	for _, position := range idPositions {
		if len(this.Bytes) < position + 4 {
			continue
		}
		id = int32(binary.BigEndian.Uint32(this.Bytes[position:position+4]))
		if id == a {
			binary.BigEndian.PutUint32(this.Bytes[position:position+4], uint32(b))
		} else if id == b {
			binary.BigEndian.PutUint32(this.Bytes[position:position+4], uint32(a))
		}
	}
}

func (this *PacketGeneric) SwapEntitiesVarInt(a int32, b int32, clientServer bool) {
	var positions []bool
	if clientServer {
		positions = PlayPacketClientEntityVarIntPositions
	} else {
		positions = PlayPacketServerEntityVarIntPositions
	}
	if this.id < 0 {
		return
	}
	if this.id >= len(positions) {
		return
	}
	if positions[this.id] == false {
		return
	}
	// Read the old Id
	buffer := bytes.NewBuffer(this.Bytes)
	bufferUtil := make([]byte, packet.UTIL_BUFFER_LENGTH)
	id, err := packet.ReadVarInt(buffer, bufferUtil)
	if err != nil {
		return
	}
	// Check the Id
	var newId int
	if id == int(a) {
		newId = int(b)
	} else if id == int(b) {
		newId = int(a)
	} else {
		return
	}
	// Apply the new Id
	newBuffer := &bytes.Buffer{}
	err = packet.WriteVarInt(newBuffer, bufferUtil, newId)
	if err != nil {
		return
	}
	buffer.WriteTo(newBuffer)
	this.Bytes = newBuffer.Bytes()
}

func (this *PacketGeneric) Id() int {
	return this.id
}

type PacketGenericCodec struct {
	Id int
}

func (this *PacketGenericCodec) Decode(reader io.Reader, util []byte) (packet packet.Packet, err error) {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return
	}
	packet = &PacketGeneric{this.Id, bytes}
	return
}

func (this *PacketGenericCodec) Encode(writer io.Writer, util []byte, packet packet.Packet) (err error) {
	_, err = writer.Write(packet.(*PacketGeneric).Bytes)
	return
}
