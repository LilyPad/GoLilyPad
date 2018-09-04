package main

import (
	"fmt"
	"github.com/LilyPad/GoLilyPad/packet/minecraft"
	"github.com/LilyPad/GoLilyPad/server/proxy/api"
)

var Plugin Example

type Example string

func (this *Example) Init(context api.Context) {
	fmt.Println("example plugin loaded, motd:", context.Config().Motd())
	context.EventBus().HandleSessionPacket(func(eventRaw api.EventSession) {
		event := eventRaw.(api.EventSessionPacket)
		packet := event.Packet()
		if packetHandshake, ok := packet.(*minecraft.PacketServerHandshake); ok {
			fmt.Println("handshake:", packetHandshake.ProtocolVersion, packetHandshake.ServerAddress, packetHandshake.ServerPort, packetHandshake.State)
		}
	}, api.PacketStagePre, api.PacketSubjectClient, api.PacketDirectionRead, api.SessionStateDisconnected)
}
