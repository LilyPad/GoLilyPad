package v19

import (
	"errors"
	"fmt"
	"github.com/LilyPad/GoLilyPad/packet"
	minecraft "github.com/LilyPad/GoLilyPad/packet/minecraft"
	"io"
)

type CodecClientTeams struct {
	IdMap *minecraft.IdMap
}

func (this *CodecClientTeams) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientTeams := new(minecraft.PacketClientTeams)
	packetClientTeams.IdFrom(this.IdMap)
	packetClientTeams.Name, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	packetClientTeams.Action, err = packet.ReadInt8(reader)
	if err != nil {
		return
	}
	if packetClientTeams.Action == minecraft.PACKET_CLIENT_TEAMS_ACTION_ADD || packetClientTeams.Action == minecraft.PACKET_CLIENT_TEAMS_ACTION_INFO_UPDATE {
		packetClientTeams.DisplayName, err = packet.ReadString(reader)
		if err != nil {
			return
		}
		packetClientTeams.Prefix, err = packet.ReadString(reader)
		if err != nil {
			return
		}
		packetClientTeams.Suffix, err = packet.ReadString(reader)
		if err != nil {
			return
		}
		packetClientTeams.FriendlyFire, err = packet.ReadInt8(reader)
		if err != nil {
			return
		}
		packetClientTeams.NameTagVisibility, err = packet.ReadString(reader)
		if err != nil {
			return
		}
		packetClientTeams.CollisionRule, err = packet.ReadString(reader)
		if err != nil {
			return
		}
		var teamColor int8
		teamColor, err = packet.ReadInt8(reader)
		if err != nil {
			return
		}
		packetClientTeams.Color = int(teamColor)
	}
	if packetClientTeams.Action == minecraft.PACKET_CLIENT_TEAMS_ACTION_ADD || packetClientTeams.Action == minecraft.PACKET_CLIENT_TEAMS_ACTION_PLAYERS_ADD || packetClientTeams.Action == minecraft.PACKET_CLIENT_TEAMS_ACTION_PLAYERS_REMOVE {
		var playersLength int
		playersLength, err = packet.ReadVarInt(reader)
		if err != nil {
			return
		}
		if playersLength < 0 {
			err = errors.New(fmt.Sprintf("Decode, Players length is below zero: %d", playersLength))
			return
		}
		if playersLength > 65535 {
			err = errors.New(fmt.Sprintf("Decode, Players length is above maximum: %d", playersLength))
			return
		}
		packetClientTeams.Players = make([]string, playersLength)
		for i := 0; i < playersLength; i++ {
			packetClientTeams.Players[i], err = packet.ReadString(reader)
			if err != nil {
				return
			}
		}
	}
	decode = packetClientTeams
	return
}

func (this *CodecClientTeams) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientTeams := encode.(*minecraft.PacketClientTeams)
	err = packet.WriteString(writer, packetClientTeams.Name)
	if err != nil {
		return
	}
	err = packet.WriteInt8(writer, packetClientTeams.Action)
	if packetClientTeams.Action == minecraft.PACKET_CLIENT_TEAMS_ACTION_ADD || packetClientTeams.Action == minecraft.PACKET_CLIENT_TEAMS_ACTION_INFO_UPDATE {
		if err != nil {
			return
		}
		err = packet.WriteString(writer, packetClientTeams.DisplayName)
		if err != nil {
			return
		}
		err = packet.WriteString(writer, packetClientTeams.Prefix)
		if err != nil {
			return
		}
		err = packet.WriteString(writer, packetClientTeams.Suffix)
		if err != nil {
			return
		}
		err = packet.WriteInt8(writer, packetClientTeams.FriendlyFire)
		if err != nil {
			return
		}
		err = packet.WriteString(writer, packetClientTeams.NameTagVisibility)
		if err != nil {
			return
		}
		if len(packetClientTeams.CollisionRule) == 0 {
			err = packet.WriteString(writer, "never")
		} else {
			err = packet.WriteString(writer, packetClientTeams.CollisionRule)
		}
		if err != nil {
			return
		}
		err = packet.WriteInt8(writer, int8(packetClientTeams.Color))
	}
	if packetClientTeams.Action == minecraft.PACKET_CLIENT_TEAMS_ACTION_ADD || packetClientTeams.Action == minecraft.PACKET_CLIENT_TEAMS_ACTION_PLAYERS_ADD || packetClientTeams.Action == minecraft.PACKET_CLIENT_TEAMS_ACTION_PLAYERS_REMOVE {
		if err != nil {
			return
		}
		err = packet.WriteVarInt(writer, len(packetClientTeams.Players))
		for i := 0; i < len(packetClientTeams.Players); i++ {
			if err != nil {
				return
			}
			err = packet.WriteString(writer, packetClientTeams.Players[i])
		}
	}
	return
}
