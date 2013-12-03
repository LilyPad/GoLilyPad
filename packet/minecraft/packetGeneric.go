package minecraft

import "encoding/binary"
import "io"
import "io/ioutil"
import "github.com/LilyPad/GoLilyPad/packet"

type PacketGeneric struct {
	id int
	Bytes []byte
}

func (this *PacketGeneric) SwapEntities(a int32, b int32) {
	if a == b {
		return
	}
	if this.id < 0 {
		return
	}
	if this.id >= len(PlayPacketClientEntityPositions) {
		return
	}
	positions := PlayPacketClientEntityPositions[this.id]
	if positions == nil {
		return
	}
	var id int32
	for _, position := range positions {
		if len(this.Bytes) >= position + 4 {
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