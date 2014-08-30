package connect

import (
	"github.com/LilyPad/GoLilyPad/packet/connect"
)

type EventMessage struct {
	Sender string
	Channel string
	Payload []byte
}

func WrapEventMessage(packet *connect.PacketMessageEvent) (this *EventMessage) {
	this = new(EventMessage)
	this.Sender = packet.Sender
	this.Channel = packet.Channel
	this.Payload = packet.Payload
	return
}
