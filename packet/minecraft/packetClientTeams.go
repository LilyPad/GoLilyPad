package minecraft

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type PacketClientTeams struct {
	Name string
	Action int8
	DisplayName string
	Prefix string
	Suffix string
	FriendlyFire int8
	Players []string
}

func (this *PacketClientTeams) Id() int {
	return PACKET_CLIENT_TEAMS
}

type PacketClientTeamsCodec struct {
	
}

func (this *PacketClientTeamsCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetClientTeams := &PacketClientTeams{}
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
		var playersSize uint16
		playersSize, err = packet.ReadUint16(reader, util)
		if err != nil {
			return
		}
		packetClientTeams.Players = make([]string, playersSize)
		var i uint16
		for i = 0; i < playersSize; i++ {
			packetClientTeams.Players[i], err = packet.ReadString(reader, util)
			if err != nil {
				return
			}
		}
	}
	return packetClientTeams, nil
}

func (this *PacketClientTeamsCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
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