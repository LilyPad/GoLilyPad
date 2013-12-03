package connect

import "github.com/LilyPad/GoLilyPad/packet/connect"

type EventMessage struct {
	Sender string
	Channel string
	Payload []byte
}

func WrapEventMessage(packet *connect.PacketMessageEvent) *EventMessage {
	return &EventMessage{
		Sender: packet.Sender,
		Channel: packet.Channel,
		Payload: packet.Payload,
	}
}