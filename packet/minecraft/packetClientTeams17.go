package minecraft

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type packetClientTeamsCodec17 struct {

}

func (this *packetClientTeamsCodec17) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetClientTeams := new(PacketClientTeams)
	packetClientTeams.Name, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	packetClientTeams.Action, err = packet.ReadInt8(reader, util)
	if err != nil {
		return
	}
	if packetClientTeams.Action == 0 || packetClientTeams.Action == 2 {
		packetClientTeams.DisplayName, err = packet.ReadString(reader, util)
		if err != nil {
			return
		}
		packetClientTeams.Prefix, err = packet.ReadString(reader, util)
		if err != nil {
			return
		}
		packetClientTeams.Suffix, err = packet.ReadString(reader, util)
		if err != nil {
			return
		}
		packetClientTeams.FriendlyFire, err = packet.ReadInt8(reader, util)
		if err != nil {
			return
		}
	}
	if packetClientTeams.Action == 0 || packetClientTeams.Action == 3 || packetClientTeams.Action == 4 {
		var playersLength uint16
		playersLength, err = packet.ReadUint16(reader, util)
		if err != nil {
			return
		}
		packetClientTeams.Players = make([]string, playersLength)
		var i uint16
		for i = 0; i < playersLength; i++ {
			packetClientTeams.Players[i], err = packet.ReadString(reader, util)
			if err != nil {
				return
			}
		}
	}
	decode = packetClientTeams
	return
}

func (this *packetClientTeamsCodec17) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetClientTeams := encode.(*PacketClientTeams)
	err = packet.WriteString(writer, util, packetClientTeams.Name)
	if err != nil {
		return
	}
	err = packet.WriteInt8(writer, util, packetClientTeams.Action)
	if packetClientTeams.Action == 0 || packetClientTeams.Action == 2 {
		if err != nil {
			return
		}
		err = packet.WriteString(writer, util, packetClientTeams.DisplayName)
		if err != nil {
			return
		}
		err = packet.WriteString(writer, util, packetClientTeams.Prefix)
		if err != nil {
			return
		}
		err = packet.WriteString(writer, util, packetClientTeams.Suffix)
		if err != nil {
			return
		}
		err = packet.WriteInt8(writer, util, packetClientTeams.FriendlyFire)
	}
	if packetClientTeams.Action == 0 || packetClientTeams.Action == 3 || packetClientTeams.Action == 4 {
		if err != nil {
			return
		}
		err = packet.WriteUint16(writer, util, uint16(len(packetClientTeams.Players)))
		for i := 0; i < len(packetClientTeams.Players); i++ {
			if err != nil {
				return
			}
			err = packet.WriteString(writer, util, packetClientTeams.Players[i])
		}
	}
	return
}
