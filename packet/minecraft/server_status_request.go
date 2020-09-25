package minecraft

import (
	"github.com/LilyPad/GoLilyPad/packet"
	"io"
)

type PacketServerStatusRequest struct {
}

func NewPacketServerStatusRequest() (this *PacketServerStatusRequest) {
	this = new(PacketServerStatusRequest)
	return
}

func (this *PacketServerStatusRequest) Id() int {
	return PACKET_SERVER_STATUS_REQUEST
}

type CodecServerStatusRequest struct {
}

func (this *CodecServerStatusRequest) Decode(reader io.Reader) (decode packet.Packet, err error) {
	decode = new(PacketServerStatusRequest)
	return
}

func (this *CodecServerStatusRequest) Encode(writer io.Writer, encode packet.Packet) (err error) {
	return
}
