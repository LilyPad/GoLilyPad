package connect

import (
	"github.com/LilyPad/GoLilyPad/packet/connect"
)

type EventServer struct {
	Add bool
	Server string
	SecurityKey string
	Address string
	Port uint16
}

func WrapEventServer(packet *connect.PacketServerEvent) (this *EventServer) {
	this = new(EventServer)
	this.Add = packet.Add
	this.Server = packet.Server
	this.SecurityKey = packet.SecurityKey
	this.Address = packet.Address
	this.Port = packet.Port
	return
}
