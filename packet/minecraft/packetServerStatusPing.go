package minecraft

import (
	"io"
	"github.com/suedadam/GoLilyPad/packet"
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

func (this *packetServerStatusPingCodec) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetServerStatusPing := new(PacketServerStatusPing)
	packetServerStatusPing.Time, err = packet.ReadInt64(reader)
	if err != nil {
		return
	}
	decode = packetServerStatusPing
	return
}

func (this *packetServerStatusPingCodec) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetServerStatusPing := encode.(*PacketServerStatusPing)
	err = packet.WriteInt64(writer, packetServerStatusPing.Time)
	return
}
