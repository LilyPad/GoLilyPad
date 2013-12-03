package connect

import "github.com/LilyPad/GoLilyPad/packet/connect"

type EventRedirect struct {
	Server string
	Player string
}

func WrapEventRedirect(packet *connect.PacketRedirectEvent) *EventRedirect {
	return &EventRedirect{
		Server: packet.Server,
		Player: packet.Player,
	}
}
