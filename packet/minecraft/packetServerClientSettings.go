package minecraft

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type PacketServerClientSettings struct {
	Locale string
	ViewDistance byte
	ChatFlags byte
	Unused bool
	Difficulty byte
	ShowCape bool
}

func (this *PacketServerClientSettings) Id() int {
	return PACKET_SERVER_CLIENT_SETTINGS
}

type PacketServerClientSettingsCodec struct {
	
}

func (this *PacketServerClientSettingsCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetServerClientSettings := &PacketServerClientSettings{}
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
	return packetServerClientSettings, nil
}

func (this *PacketServerClientSettingsCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
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
