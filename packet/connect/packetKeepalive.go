package connect

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type PacketKeepalive struct {
	Random int32
}

func NewPacketKeepalive(random int32) (this *PacketKeepalive) {
	this = new(PacketKeepalive)
	this.Random = random
	return
}

func (this *PacketKeepalive) Id() int {
	return PACKET_KEEPALIVE
}

type packetKeepaliveCodec struct {

}

func (this *packetKeepaliveCodec) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetKeepalive := new(PacketKeepalive)
	packetKeepalive.Random, err = packet.ReadInt32(reader)
	if err != nil {
		return
	}
	decode = packetKeepalive
	return
}

func (this *packetKeepaliveCodec) Encode(writer io.Writer, encode packet.Packet) (err error) {
	err = packet.WriteInt32(writer, encode.(*PacketKeepalive).Random)
	return
}
