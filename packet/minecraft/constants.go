package minecraft

import (
	"github.com/LilyPad/GoLilyPad/packet"
)

const (
	STRING_VERSION = string("1.12.2")
	MAGIC          = string("§")

	PACKET_SERVER_HANDSHAKE = 0x00

	PACKET_CLIENT_STATUS_RESPONSE = 0x00
	PACKET_CLIENT_STATUS_PING     = 0x01
	PACKET_SERVER_STATUS_REQUEST  = 0x00
	PACKET_SERVER_STATUS_PING     = 0x01
)

var Versions = []int{340, 338, 335, 316, 315, 210, 110, 109, 108, 107, 47, 5, 4}

var HandshakePacketServerCodec = packet.NewPacketCodecRegistryDual([]packet.PacketCodec{}, []packet.PacketCodec{
	PACKET_SERVER_HANDSHAKE: new(CodecServerHandshake),
})

var HandshakePacketClientCodec = HandshakePacketServerCodec.Flip()

var StatusPacketServerCodec = packet.NewPacketCodecRegistryDual([]packet.PacketCodec{
	PACKET_CLIENT_STATUS_RESPONSE: new(CodecClientStatusResponse),
	PACKET_CLIENT_STATUS_PING:     new(CodecClientStatusPing),
}, []packet.PacketCodec{
	PACKET_SERVER_STATUS_REQUEST: new(CodecServerStatusRequest),
	PACKET_SERVER_STATUS_PING:    new(CodecServerStatusPing),
})

var StatusPacketClientCodec = StatusPacketServerCodec.Flip()
