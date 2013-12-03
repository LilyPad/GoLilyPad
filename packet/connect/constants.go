package connect

import "github.com/LilyPad/GoLilyPad/packet"

const (
	STATUS_ERROR_ROLE = 0x02
	STATUS_ERROR_GENERIC = 0x01
	STATUS_SUCCESS = 0x00

	REQUEST_AUTHENTICATE = 0x00
	REQUEST_AS_SERVER = 0x01
	REQUEST_AS_PROXY = 0x02
	REQUEST_GET_SALT = 0x03
	REQUEST_GET_WHOAMI = 0x04
	REQUEST_MESSAGE = 0x10
	REQUEST_REDIRECT = 0x11
	REQUEST_GET_PLAYERS = 0x20
	REQUEST_NOTIFY_PLAYER = 0x21
	REQUEST_GET_DETAILS = 0x22

	PACKET_KEEPALIVE = 0x00
	PACKET_REQUEST = 0x01
	PACKET_RESULT = 0x02
	PACKET_MESSAGE_EVENT = 0x03
	PACKET_REDIRECT_EVENT = 0x04
	PACKET_SERVER_EVENT = 0x05
)

var requestCodecs = []RequestCodec {
	REQUEST_AUTHENTICATE: &RequestAuthenticateCodec{},
	REQUEST_AS_SERVER: &RequestAsServerCodec{},
	REQUEST_AS_PROXY: &RequestAsProxyCodec{},
	REQUEST_GET_SALT: &RequestGetSaltCodec{},
	REQUEST_GET_WHOAMI: &RequestGetWhoamiCodec{},
	REQUEST_MESSAGE: &RequestMessageCodec{},
	REQUEST_REDIRECT: &RequestRedirectCodec{},
	REQUEST_GET_PLAYERS: &RequestGetPlayersCodec{},
	REQUEST_NOTIFY_PLAYER: &RequestNotifyPlayerCodec{},
	REQUEST_GET_DETAILS: &RequestGetDetailsCodec{},
}

var resultCodecs = []ResultCodec {
	REQUEST_AUTHENTICATE: &ResultAuthenticateCodec{},
	REQUEST_AS_SERVER: &ResultAsServerCodec{},
	REQUEST_AS_PROXY: &ResultAsProxyCodec{},
	REQUEST_GET_SALT: &ResultGetSaltCodec{},
	REQUEST_GET_WHOAMI: &ResultGetWhoamiCodec{},
	REQUEST_MESSAGE: &ResultMessageCodec{},
	REQUEST_REDIRECT: &ResultRedirectCodec{},
	REQUEST_GET_PLAYERS: &ResultGetPlayersCodec{},
	REQUEST_NOTIFY_PLAYER: &ResultNotifyPlayerCodec{},
	REQUEST_GET_DETAILS: &ResultGetDetailsCodec{},
}

var PacketCodecs = []packet.PacketCodec {
	PACKET_KEEPALIVE: &PacketKeepaliveCodec{},
	PACKET_REQUEST: &PacketRequestCodec{},
	PACKET_RESULT: &PacketResultCodec{nil},
	PACKET_MESSAGE_EVENT: &PacketMessageEventCodec{},
	PACKET_REDIRECT_EVENT: &PacketRedirectEventCodec{},
	PACKET_SERVER_EVENT: &PacketServerEventCodec{},
}
var PacketCodec = packet.NewPacketCodecVarIntLength(packet.NewPacketCodecRegistry(PacketCodecs))
