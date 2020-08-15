package connect

import (
	"github.com/LilyPad/GoLilyPad/packet/connect"
	"github.com/satori/go.uuid"
)

type EventPlayer struct {
	Joining    bool
	PlayerName string
	PlayerUUID uuid.UUID
}

func WrapEventPlayer(packet *connect.PacketPlayerEvent) (this *EventPlayer) {
	this = new(EventPlayer)
	this.Joining = packet.Joining
	this.PlayerName = packet.PlayerName
	this.PlayerUUID = packet.PlayerUUID
	return
}
