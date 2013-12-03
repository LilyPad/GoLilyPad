package minecraft

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type PacketClientScoreboardObjective struct {
	Name string
	Value string
	Action int8
}

func (this *PacketClientScoreboardObjective) Id() int {
	return PACKET_CLIENT_SCOREBOARD_OBJECTIVE
}

type PacketClientScoreboardObjectiveCodec struct {
	
}

func (this *PacketClientScoreboardObjectiveCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetClientScoreboardObjective := &PacketClientScoreboardObjective{}
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
	return packetClientScoreboardObjective, nil
}

func (this *PacketClientScoreboardObjectiveCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
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
