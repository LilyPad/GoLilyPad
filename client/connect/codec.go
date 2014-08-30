package connect

import (
	"github.com/LilyPad/GoLilyPad/packet"
	"github.com/LilyPad/GoLilyPad/packet/connect"
)

func NewCodec(connectClient Connect) (packetCodec packet.PacketCodec) {
	packetCodecs := make([]packet.PacketCodec, len(connect.PacketCodecs))
	copy(packetCodecs, connect.PacketCodecs)
	packetCodecs[connect.PACKET_RESULT] = connect.NewPacketResultCodec(connectClient)
	packetCodec = packet.NewPacketCodecVarIntLength(packet.NewPacketCodecRegistry(packetCodecs))
	return
}
