package connect

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type PacketRedirectEvent struct {
	Server string
	Player string
}

func NewPacketRedirectEvent(server string, player string) (this *PacketRedirectEvent) {
	this = new(PacketRedirectEvent)
	this.Server = server
	this.Player = player
	return
}

func (this *PacketRedirectEvent) Id() int {
	return PACKET_REDIRECT_EVENT
}

type packetRedirectEventCodec struct {

}

func (this *packetRedirectEventCodec) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetRedirectEvent := new(PacketRedirectEvent)
	packetRedirectEvent.Server, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	packetRedirectEvent.Player, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	decode = packetRedirectEvent
	return
}

func (this *packetRedirectEventCodec) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetRedirectEvent := encode.(*PacketRedirectEvent)
	err = packet.WriteString(writer, packetRedirectEvent.Server)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, packetRedirectEvent.Player)
	return
}
