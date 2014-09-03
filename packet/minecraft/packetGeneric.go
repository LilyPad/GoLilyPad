package minecraft

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"io"
	"io/ioutil"
	"github.com/LilyPad/GoLilyPad/packet"
)

type PacketGeneric struct {
	id int
	Bytes []byte
	compressed bool
}

func NewPacketGeneric(id int, bytes []byte, compressed bool) (this *PacketGeneric) {
	this = new(PacketGeneric)
	this.id = id
	this.Bytes = bytes
	this.compressed = compressed
	return
}

func (this *PacketGeneric) Decompress() (err error) {
	if !this.compressed {
		return
	}
	buffer := bytes.NewReader(this.Bytes)
	bufferUtil := make([]byte, packet.UTIL_BUFFER_LENGTH)
	zlibReader, err := zlib.NewReader(buffer)
	if err != nil {
		return
	}
	_, err = packet.ReadVarInt(zlibReader, bufferUtil) // compression length
	if err != nil {
		return
	}
	_, err = packet.ReadVarInt(zlibReader, bufferUtil) // id
	if err != nil {
		return
	}
	bytes, err := ioutil.ReadAll(zlibReader)
	if err != nil {
		return
	}
	this.Bytes = bytes
	this.compressed = false
	return
}

func (this *PacketGeneric) SwapEntities(a int32, b int32, clientServer bool, protocol17 bool) {
	if a == b {
		return
	}
	if this.id == PACKET_CLIENT_SPAWN_OBJECT && clientServer {
		this.Decompress()
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
		this.Decompress()
		// TODO
	}
	this.swapEntitiesInt(a, b, clientServer, protocol17)
	this.swapEntitiesVarInt(a, b, clientServer, protocol17)
}

func (this *PacketGeneric) swapEntitiesInt(a int32, b int32, clientServer bool, protocol17 bool) {
	var positions [][]int
	if clientServer {
		if protocol17 {
			positions = PlayPacketClientEntityIntPositions17
		} else {
			positions = PlayPacketClientEntityIntPositions
		}
	} else {
		if protocol17 {
			positions = PlayPacketServerEntityIntPositions17
		} else {
			positions = PlayPacketServerEntityIntPositions
		}
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
	this.Decompress()
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

func (this *PacketGeneric) swapEntitiesVarInt(a int32, b int32, clientServer bool, protocol17 bool) {
	var positions []bool
	if clientServer {
		if protocol17 {
			positions = PlayPacketClientEntityVarIntPositions17
		} else {
			positions = PlayPacketClientEntityVarIntPositions
		}
	} else {
		if protocol17 {
			positions = PlayPacketServerEntityVarIntPositions17
		} else {
			positions = PlayPacketServerEntityVarIntPositions
		}
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
	this.Decompress()
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
	newBuffer := new(bytes.Buffer)
	err = packet.WriteVarInt(newBuffer, bufferUtil, newId)
	if err != nil {
		return
	}
	buffer.WriteTo(newBuffer)
	this.Bytes = newBuffer.Bytes()
}

func (this *PacketGeneric) Raw() bool {
	return this.compressed
}

func (this *PacketGeneric) Id() int {
	return this.id
}

type packetGenericCodec struct {
	Id int
}

func NewPacketGenericCodec(id int) (this *packetGenericCodec) {
	this = new(packetGenericCodec)
	this.Id = id
	return
}

func (this *packetGenericCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetGeneric := new(PacketGeneric)
	packetGeneric.id = this.Id
	if zlibReader, ok := reader.(*packet.ZlibToggleReader); ok {
		zlibReader.SetRaw(false)
		packetGeneric.compressed = true
	}
	packetGeneric.Bytes, err = ioutil.ReadAll(reader)
	if err != nil {
		return
	}
	decode = packetGeneric
	return
}

func (this *packetGenericCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	_, err = writer.Write(encode.(*PacketGeneric).Bytes)
	return
}
