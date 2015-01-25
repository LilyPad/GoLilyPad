package minecraft

import (
	"errors"
	"fmt"
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

const (
	PACKET_CLIENT_TEAMS_ACTION_ADD = int8(0)
	PACKET_CLIENT_TEAMS_ACTION_REMOVE = int8(1)
	PACKET_CLIENT_TEAMS_ACTION_INFO_UPDATE = int8(2)
	PACKET_CLIENT_TEAMS_ACTION_PLAYERS_ADD = int8(3)
	PACKET_CLIENT_TEAMS_ACTION_PLAYERS_REMOVE = int8(4)
)

type PacketClientTeams struct {
	Name string
	Action int8
	DisplayName string
	Prefix string
	Suffix string
	FriendlyFire int8
	NameTagVisibility string
	Color int8
	Players []string
}

func NewPacketClientTeamsAdd(name string, displayName string, prefix string, suffix string, friendlyFire int8, nameTagVisibility string, color int8, players []string) (this *PacketClientTeams) {
	this = new(PacketClientTeams)
	this.Name = name
	this.Action = PACKET_CLIENT_TEAMS_ACTION_ADD
	this.DisplayName = displayName
	this.Prefix = prefix
	this.Suffix = suffix
	this.FriendlyFire = friendlyFire
	this.NameTagVisibility = nameTagVisibility
	this.Color = color
	this.Players = players
	return
}

func NewPacketClientTeamsRemove(name string) (this *PacketClientTeams) {
	this = new(PacketClientTeams)
	this.Name = name
	this.Action = PACKET_CLIENT_TEAMS_ACTION_REMOVE
	return
}

func NewPacketClientTeamsInfoUpdate(name string, displayName string, prefix string, suffix string, friendlyFire int8, nameTagVisibility string, color int8) (this *PacketClientTeams) {
	this = new(PacketClientTeams)
	this.Name = name
	this.Action = PACKET_CLIENT_TEAMS_ACTION_INFO_UPDATE
	this.DisplayName = displayName
	this.Prefix = prefix
	this.Suffix = suffix
	this.FriendlyFire = friendlyFire
	this.NameTagVisibility = nameTagVisibility
	this.Color = color
	return
}

func NewPacketClientTeamsPlayersAdd(name string, players []string) (this *PacketClientTeams) {
	this = new(PacketClientTeams)
	this.Name = name
	this.Action = PACKET_CLIENT_TEAMS_ACTION_PLAYERS_ADD
	this.Players = players
	return
}

func NewPacketClientTeamsPlayersRemove(name string, players []string) (this *PacketClientTeams) {
	this = new(PacketClientTeams)
	this.Name = name
	this.Action = PACKET_CLIENT_TEAMS_ACTION_PLAYERS_REMOVE
	this.Players = players
	return
}

func (this *PacketClientTeams) Id() int {
	return PACKET_CLIENT_TEAMS
}

type packetClientTeamsCodec struct {

}

func (this *packetClientTeamsCodec) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientTeams := new(PacketClientTeams)
	packetClientTeams.Name, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	packetClientTeams.Action, err = packet.ReadInt8(reader)
	if err != nil {
		return
	}
	if packetClientTeams.Action == 0 || packetClientTeams.Action == 2 {
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
		packetClientTeams.Color, err = packet.ReadInt8(reader)
		if err != nil {
			return
		}
	}
	if packetClientTeams.Action == 0 || packetClientTeams.Action == 3 || packetClientTeams.Action == 4 {
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

func (this *packetClientTeamsCodec) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientTeams := encode.(*PacketClientTeams)
	err = packet.WriteString(writer, packetClientTeams.Name)
	if err != nil {
		return
	}
	err = packet.WriteInt8(writer, packetClientTeams.Action)
	if packetClientTeams.Action == 0 || packetClientTeams.Action == 2 {
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
		err = packet.WriteInt8(writer, packetClientTeams.Color)
	}
	if packetClientTeams.Action == 0 || packetClientTeams.Action == 3 || packetClientTeams.Action == 4 {
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
