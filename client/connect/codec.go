package connect

import "github.com/LilyPad/GoLilyPad/packet"
import "github.com/LilyPad/GoLilyPad/packet/connect"

func NewCodec(connectClient Connect) packet.PacketCodec {
	packetCodecs := make([]packet.PacketCodec, len(connect.PacketCodecs))
	copy(packetCodecs, connect.PacketCodecs)
	packetCodecs[connect.PACKET_RESULT] = connect.NewPacketResultCodec(connectClient)
	return packet.NewPacketCodecVarIntLength(packet.NewPacketCodecRegistry(packetCodecs))
}
