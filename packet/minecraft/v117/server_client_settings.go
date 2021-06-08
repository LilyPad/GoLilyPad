package v117

import (
	"github.com/LilyPad/GoLilyPad/packet"
	minecraft "github.com/LilyPad/GoLilyPad/packet/minecraft"
	"io"
)

type CodecServerClientSettings struct {
	IdMap *minecraft.IdMap
}

func (this *CodecServerClientSettings) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetServerClientSettings := new(minecraft.PacketServerClientSettings)
	packetServerClientSettings.IdFrom(this.IdMap)
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
	packetServerClientSettings.SkinParts, err = packet.ReadUint8(reader)
	if err != nil {
		return
	}
	packetServerClientSettings.MainHand, err = packet.ReadVarInt(reader)
	if err != nil {
		return
	}
	packetServerClientSettings.DisableTextFiltering, err = packet.ReadBool(reader)
	if err != nil {
		return
	}
	decode = packetServerClientSettings
	return
}

func (this *CodecServerClientSettings) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetServerClientSettings := encode.(*minecraft.PacketServerClientSettings)
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
	err = packet.WriteUint8(writer, packetServerClientSettings.SkinParts)
	if err != nil {
		return
	}
	err = packet.WriteVarInt(writer, packetServerClientSettings.MainHand)
	if err != nil {
		return
	}
	err = packet.WriteBool(writer, packetServerClientSettings.DisableTextFiltering)
	return
}
