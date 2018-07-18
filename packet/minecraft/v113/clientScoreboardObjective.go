package v113

import (
	"errors"
	"fmt"
	"github.com/LilyPad/GoLilyPad/packet"
	minecraft "github.com/LilyPad/GoLilyPad/packet/minecraft"
	"io"
	"strconv"
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
		// 1.13 // We store the Type as a formatted string for compatibility
		var valueType int
		valueType, err = packet.ReadVarInt(reader)
		if err != nil {
			return
		}
		packetClientScoreboardObjective.Type = strconv.FormatInt(int64(valueType), 10)
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
		// 1.13 // We store the Type as a formatted string for compatibility
		var valueType int64
		valueType, err = strconv.ParseInt(packetClientScoreboardObjective.Type, 10, 64)
		if err != nil {
			return
		}
		err = packet.WriteVarInt(writer, int(valueType))
	case minecraft.PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_REMOVE:
		// no payload
	default:
		err = errors.New(fmt.Sprintf("Encode, PacketClientScoreboardObjective action is not valid: %d", packetClientScoreboardObjective.Action))
	}
	return
}
