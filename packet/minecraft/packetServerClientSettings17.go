package minecraft

import (
	"github.com/LilyPad/GoLilyPad/packet"
	"io"
)

type packetServerClientSettingsCodec17 struct {
}

func (this *packetServerClientSettingsCodec17) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetServerClientSettings := new(PacketServerClientSettings)
	packetServerClientSettings.Locale, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	packetServerClientSettings.ViewDistance, err = packet.ReadUint8(reader)
	if err != nil {
		return
	}
	packetServerClientSettings.ChatFlags, err = packet.ReadUint8(reader)
	if err != nil {
		return
	}
	packetServerClientSettings.ChatColours, err = packet.ReadBool(reader)
	if err != nil {
		return
	}
	packetServerClientSettings.difficulty17, err = packet.ReadUint8(reader)
	if err != nil {
		return
	}
	packetServerClientSettings.SkinParts, err = packet.ReadUint8(reader)
	if err != nil {
		return
	}
	decode = packetServerClientSettings
	return
}

func (this *packetServerClientSettingsCodec17) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetServerClientSettings := encode.(*PacketServerClientSettings)
	err = packet.WriteString(writer, packetServerClientSettings.Locale)
	if err != nil {
		return
	}
	err = packet.WriteUint8(writer, packetServerClientSettings.ViewDistance)
	if err != nil {
		return
	}
	err = packet.WriteUint8(writer, packetServerClientSettings.ChatFlags)
	if err != nil {
		return
	}
	err = packet.WriteBool(writer, packetServerClientSettings.ChatColours)
	if err != nil {
		return
	}
	err = packet.WriteUint8(writer, packetServerClientSettings.difficulty17)
	if err != nil {
		return
	}
	err = packet.WriteUint8(writer, packetServerClientSettings.SkinParts)
	return
}
