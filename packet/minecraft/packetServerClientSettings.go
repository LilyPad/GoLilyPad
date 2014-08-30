package minecraft

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type PacketServerClientSettings struct {
	Locale string
	ViewDistance byte
	ChatFlags byte
	Unused bool
	Difficulty byte
	ShowCape bool
}

func NewPacketServerClientSettings(locale string, viewDistance byte, chatFlags byte, unused bool, difficulty byte, showCape bool) (this *PacketServerClientSettings) {
	this = new(PacketServerClientSettings)
	this.Locale = locale
	this.ViewDistance = viewDistance
	this.ChatFlags = chatFlags
	this.Unused = unused
	this.Difficulty = difficulty
	this.ShowCape = showCape
	return
}

func (this *PacketServerClientSettings) Id() int {
	return PACKET_SERVER_CLIENT_SETTINGS
}

type packetServerClientSettingsCodec struct {

}

func (this *packetServerClientSettingsCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetServerClientSettings := new(PacketServerClientSettings)
	packetServerClientSettings.Locale, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	packetServerClientSettings.ViewDistance, err = packet.ReadUint8(reader, util)
	if err != nil {
		return
	}
	packetServerClientSettings.ChatFlags, err = packet.ReadUint8(reader, util)
	if err != nil {
		return
	}
	packetServerClientSettings.Unused, err = packet.ReadBool(reader, util)
	if err != nil {
		return
	}
	packetServerClientSettings.Difficulty, err = packet.ReadUint8(reader, util)
	if err != nil {
		return
	}
	packetServerClientSettings.ShowCape, err = packet.ReadBool(reader, util)
	if err != nil {
		return
	}
	decode = packetServerClientSettings
	return
}

func (this *packetServerClientSettingsCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetServerClientSettings := encode.(*PacketServerClientSettings)
	err = packet.WriteString(writer, util, packetServerClientSettings.Locale)
	if err != nil {
		return
	}
	err = packet.WriteUint8(writer, util, packetServerClientSettings.ViewDistance)
	if err != nil {
		return
	}
	err = packet.WriteUint8(writer, util, packetServerClientSettings.ChatFlags)
	if err != nil {
		return
	}
	err = packet.WriteBool(writer, util, packetServerClientSettings.Unused)
	if err != nil {
		return
	}
	err = packet.WriteUint8(writer, util, packetServerClientSettings.Difficulty)
	if err != nil {
		return
	}
	err = packet.WriteBool(writer, util, packetServerClientSettings.ShowCape)
	return
}
