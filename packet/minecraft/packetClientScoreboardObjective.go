package minecraft

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type PacketClientScoreboardObjective struct {
	Name string
	Value string
	Action int8
}

func NewPacketClientScoreboardObjectiveAdd(name string, value string) (this *PacketClientScoreboardObjective) {
	this = new(PacketClientScoreboardObjective)
	this.Name = name
	this.Value = value
	this.Action = 0
	return
}

func NewPacketClientScoreboardObjectiveRemove(name string, value string) (this *PacketClientScoreboardObjective) {
	this = new(PacketClientScoreboardObjective)
	this.Name = name
	this.Value = value
	this.Action = 1
	return
}

func NewPacketClientScoreboardObjectiveUpdate(name string, value string) (this *PacketClientScoreboardObjective) {
	this = new(PacketClientScoreboardObjective)
	this.Name = name
	this.Value = value
	this.Action = 2
	return
}

func (this *PacketClientScoreboardObjective) Id() int {
	return PACKET_CLIENT_SCOREBOARD_OBJECTIVE
}

type packetClientScoreboardObjectiveCodec struct {

}

func (this *packetClientScoreboardObjectiveCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetClientScoreboardObjective := new(PacketClientScoreboardObjective)
	packetClientScoreboardObjective.Name, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	packetClientScoreboardObjective.Value, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	packetClientScoreboardObjective.Action, err = packet.ReadInt8(reader, util)
	if err != nil {
		return
	}
	decode = packetClientScoreboardObjective
	return
}

func (this *packetClientScoreboardObjectiveCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetClientScoreboardObjective := encode.(*PacketClientScoreboardObjective)
	err = packet.WriteString(writer, util, packetClientScoreboardObjective.Name)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, util, packetClientScoreboardObjective.Value)
	if err != nil {
		return
	}
	err = packet.WriteInt8(writer, util, packetClientScoreboardObjective.Action)
	return
}
