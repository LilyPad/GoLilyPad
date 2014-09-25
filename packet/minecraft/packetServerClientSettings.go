package minecraft

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type PacketServerClientSettings struct {
	Locale string
	ViewDistance byte
	ChatFlags byte
	ChatColours bool
	difficulty17 byte
	SkinParts byte
}

func NewPacketServerClientSettings(locale string, viewDistance byte, chatFlags byte, chatColours bool, skinParts byte) (this *PacketServerClientSettings) {
	this = new(PacketServerClientSettings)
	this.Locale = locale
	this.ViewDistance = viewDistance
	this.ChatFlags = chatFlags
	this.ChatColours = chatColours
	this.SkinParts = skinParts
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
	packetServerClientSettings.ChatColours, err = packet.ReadBool(reader, util)
	if err != nil {
		return
	}
	packetServerClientSettings.SkinParts, err = packet.ReadUint8(reader, util)
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
	err = packet.WriteBool(writer, util, packetServerClientSettings.ChatColours)
	if err != nil {
		return
	}
	err = packet.WriteUint8(writer, util, packetServerClientSettings.SkinParts)
	return
}
