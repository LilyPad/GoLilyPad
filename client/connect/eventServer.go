package connect

import "github.com/LilyPad/GoLilyPad/packet/connect"

type EventServer struct {
	Add bool
	Server string
	SecurityKey string
	Address string
	Port uint16
}

func WrapEventServer(packet *connect.PacketServerEvent) *EventServer {
	return &EventServer{
		Add: packet.Add,
		Server: packet.Server,
		SecurityKey: packet.SecurityKey,
		Address: packet.Address,
		Port: packet.Port,
	}
}