package minecraft

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type PacketServerStatusPing struct {
	Time int64
}

func NewPacketServerStatusPing(time int64) (this *PacketServerStatusPing) {
	this = new(PacketServerStatusPing)
	this.Time = time
	return
}

func (this *PacketServerStatusPing) Id() int {
	return PACKET_SERVER_STATUS_PING
}

type packetServerStatusPingCodec struct {

}

func (this *packetServerStatusPingCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetServerStatusPing := new(PacketServerStatusPing)
	packetServerStatusPing.Time, err = packet.ReadInt64(reader, util)
	if err != nil {
		return
	}
	decode = packetServerStatusPing
	return
}

func (this *packetServerStatusPingCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetServerStatusPing := encode.(*PacketServerStatusPing)
	err = packet.WriteInt64(writer, util, packetServerStatusPing.Time)
	return
}
