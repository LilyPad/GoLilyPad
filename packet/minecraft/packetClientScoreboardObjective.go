package minecraft

import (
	"errors"
	"fmt"
	"github.com/LilyPad/GoLilyPad/packet"
	"io"
)

const (
	PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_ADD    = int8(0)
	PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_REMOVE = int8(1)
	PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_UPDATE = int8(2)
)

type PacketClientScoreboardObjective struct {
	Name   string
	Action int8
	Value  string
	Type   string
}

func NewPacketClientScoreboardObjectiveAdd(name string, value string, stype string) (this *PacketClientScoreboardObjective) {
	this = new(PacketClientScoreboardObjective)
	this.Name = name
	this.Action = PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_ADD
	this.Value = value
	this.Type = stype
	return
}

func NewPacketClientScoreboardObjectiveRemove(name string) (this *PacketClientScoreboardObjective) {
	this = new(PacketClientScoreboardObjective)
	this.Name = name
	this.Action = PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_REMOVE
	return
}

func NewPacketClientScoreboardObjectiveUpdate(name string, value string, stype string) (this *PacketClientScoreboardObjective) {
	this = new(PacketClientScoreboardObjective)
	this.Name = name
	this.Value = value
	this.Action = PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_UPDATE
	this.Type = stype
	return
}

func (this *PacketClientScoreboardObjective) Id() int {
	return PACKET_CLIENT_SCOREBOARD_OBJECTIVE
}

type packetClientScoreboardObjectiveCodec struct {
}

func (this *packetClientScoreboardObjectiveCodec) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientScoreboardObjective := new(PacketClientScoreboardObjective)
	packetClientScoreboardObjective.Name, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	packetClientScoreboardObjective.Action, err = packet.ReadInt8(reader)
	if err != nil {
		return
	}
	switch packetClientScoreboardObjective.Action {
	case PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_ADD:
		fallthrough
	case PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_UPDATE:
		packetClientScoreboardObjective.Value, err = packet.ReadString(reader)
		if err != nil {
			return
		}
		packetClientScoreboardObjective.Type, err = packet.ReadString(reader)
		if err != nil {
			return
		}
	case PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_REMOVE:
		// no payload
	default:
		err = errors.New(fmt.Sprintf("Decode, PacketClientScoreboardObjective action is not valid: %d", packetClientScoreboardObjective.Action))
	}
	decode = packetClientScoreboardObjective
	return
}

func (this *packetClientScoreboardObjectiveCodec) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientScoreboardObjective := encode.(*PacketClientScoreboardObjective)
	err = packet.WriteString(writer, packetClientScoreboardObjective.Name)
	if err != nil {
		return
	}
	err = packet.WriteInt8(writer, packetClientScoreboardObjective.Action)
	if err != nil {
		return
	}
	switch packetClientScoreboardObjective.Action {
	case PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_ADD:
		fallthrough
	case PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_UPDATE:
		err = packet.WriteString(writer, packetClientScoreboardObjective.Value)
		if err != nil {
			return
		}
		err = packet.WriteString(writer, packetClientScoreboardObjective.Type)
	case PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_REMOVE:
		// no payload
	default:
		err = errors.New(fmt.Sprintf("Encode, PacketClientScoreboardObjective action is not valid: %d", packetClientScoreboardObjective.Action))
	}
	return
}
