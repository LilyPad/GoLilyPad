package v18

import (
	"errors"
	"fmt"
	"github.com/LilyPad/GoLilyPad/packet"
	minecraft "github.com/LilyPad/GoLilyPad/packet/minecraft"
	"io"
)

type CodecClientScoreboardObjective struct {
	IdMap *minecraft.IdMap
}

func (this *CodecClientScoreboardObjective) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientScoreboardObjective := new(minecraft.PacketClientScoreboardObjective)
	packetClientScoreboardObjective.IdFrom(this.IdMap)
	packetClientScoreboardObjective.Name, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	packetClientScoreboardObjective.Action, err = packet.ReadInt8(reader)
	if err != nil {
		return
	}
	switch packetClientScoreboardObjective.Action {
	case minecraft.PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_ADD:
		fallthrough
	case minecraft.PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_UPDATE:
		packetClientScoreboardObjective.Value, err = packet.ReadString(reader)
		if err != nil {
			return
		}
		packetClientScoreboardObjective.Type, err = packet.ReadString(reader)
		if err != nil {
			return
		}
	case minecraft.PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_REMOVE:
		// no payload
	default:
		err = errors.New(fmt.Sprintf("Decode, PacketClientScoreboardObjective action is not valid: %d", packetClientScoreboardObjective.Action))
	}
	decode = packetClientScoreboardObjective
	return
}

func (this *CodecClientScoreboardObjective) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientScoreboardObjective := encode.(*minecraft.PacketClientScoreboardObjective)
	err = packet.WriteString(writer, packetClientScoreboardObjective.Name)
	if err != nil {
		return
	}
	err = packet.WriteInt8(writer, packetClientScoreboardObjective.Action)
	if err != nil {
		return
	}
	switch packetClientScoreboardObjective.Action {
	case minecraft.PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_ADD:
		fallthrough
	case minecraft.PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_UPDATE:
		err = packet.WriteString(writer, packetClientScoreboardObjective.Value)
		if err != nil {
			return
		}
		err = packet.WriteString(writer, packetClientScoreboardObjective.Type)
	case minecraft.PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_REMOVE:
		// no payload
	default:
		err = errors.New(fmt.Sprintf("Encode, PacketClientScoreboardObjective action is not valid: %d", packetClientScoreboardObjective.Action))
	}
	return
}
