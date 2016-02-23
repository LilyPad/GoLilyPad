package minecraft

import (
	"github.com/LilyPad/GoLilyPad/packet"
	"io"
)

type PacketClientStatusPing struct {
	Time int64
}

func NewPacketClientStatusPing(time int64) (this *PacketClientStatusPing) {
	this = new(PacketClientStatusPing)
	this.Time = time
	return
}

func (this *PacketClientStatusPing) Id() int {
	return PACKET_CLIENT_STATUS_PING
}

type CodecClientStatusPing struct {
}

func (this *CodecClientStatusPing) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientStatusPing := new(PacketClientStatusPing)
	packetClientStatusPing.Time, err = packet.ReadInt64(reader)
	if err != nil {
		return
	}
	decode = packetClientStatusPing
	return
}

func (this *CodecClientStatusPing) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientStatusPing := encode.(*PacketClientStatusPing)
	err = packet.WriteInt64(writer, packetClientStatusPing.Time)
	return
}
