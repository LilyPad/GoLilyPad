package minecraft

import (
	"github.com/LilyPad/GoLilyPad/packet"
)

type Version struct {
	Name             string
	NameLatest       string
	LoginClientCodec packet.PacketPipelineChild
	LoginServerCodec packet.PacketPipelineChild
	PlayClientCodec  packet.PacketPipelineChild
	PlayServerCodec  packet.PacketPipelineChild
	IdMap            *IdMap
	Id               []int
	IdLatest         int
}
