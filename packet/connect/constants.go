package connect

import (
	"github.com/LilyPad/GoLilyPad/packet"
)

const (
	STATUS_ERROR_ROLE    = 0x02
	STATUS_ERROR_GENERIC = 0x01
	STATUS_SUCCESS       = 0x00

	REQUEST_AUTHENTICATE  = 0x00
	REQUEST_AS_SERVER     = 0x01
	REQUEST_AS_PROXY      = 0x02
	REQUEST_GET_SALT      = 0x03
	REQUEST_GET_WHOAMI    = 0x04
	REQUEST_MESSAGE       = 0x10
	REQUEST_REDIRECT      = 0x11
	REQUEST_GET_PLAYERS   = 0x20
	REQUEST_NOTIFY_PLAYER = 0x21
	REQUEST_GET_DETAILS   = 0x22

	PACKET_KEEPALIVE      = 0x00
	PACKET_REQUEST        = 0x01
	PACKET_RESULT         = 0x02
	PACKET_MESSAGE_EVENT  = 0x03
	PACKET_REDIRECT_EVENT = 0x04
	PACKET_SERVER_EVENT   = 0x05
	PACKET_PLAYER_EVENT   = 0x06
)

var requestCodecs = []RequestCodec{
	REQUEST_AUTHENTICATE:  new(requestAuthenticateCodec),
	REQUEST_AS_SERVER:     new(requestAsServerCodec),
	REQUEST_AS_PROXY:      new(requestAsProxyCodec),
	REQUEST_GET_SALT:      new(requestGetSaltCodec),
	REQUEST_GET_WHOAMI:    new(requestGetWhoamiCodec),
	REQUEST_MESSAGE:       new(requestMessageCodec),
	REQUEST_REDIRECT:      new(requestRedirectCodec),
	REQUEST_GET_PLAYERS:   new(requestGetPlayersCodec),
	REQUEST_NOTIFY_PLAYER: new(requestNotifyPlayerCodec),
	REQUEST_GET_DETAILS:   new(requestGetDetailsCodec),
}

var resultCodecs = []ResultCodec{
	REQUEST_AUTHENTICATE:  new(resultAuthenticateCodec),
	REQUEST_AS_SERVER:     new(resultAsServerCodec),
	REQUEST_AS_PROXY:      new(resultAsProxyCodec),
	REQUEST_GET_SALT:      new(resultGetSaltCodec),
	REQUEST_GET_WHOAMI:    new(resultGetWhoamiCodec),
	REQUEST_MESSAGE:       new(resultMessageCodec),
	REQUEST_REDIRECT:      new(resultRedirectCodec),
	REQUEST_GET_PLAYERS:   new(resultGetPlayersCodec),
	REQUEST_NOTIFY_PLAYER: new(resultNotifyPlayerCodec),
	REQUEST_GET_DETAILS:   new(resultGetDetailsCodec),
}

var PacketCodec = packet.NewPacketCodecRegistry([]packet.PacketCodec{
	PACKET_KEEPALIVE:      new(packetKeepaliveCodec),
	PACKET_REQUEST:        new(packetRequestCodec),
	PACKET_RESULT:         new(packetResultCodec),
	PACKET_MESSAGE_EVENT:  new(packetMessageEventCodec),
	PACKET_REDIRECT_EVENT: new(packetRedirectEventCodec),
	PACKET_SERVER_EVENT:   new(packetServerEventCodec),
	PACKET_PLAYER_EVENT:   new(packetPlayerEventCodec),
})
