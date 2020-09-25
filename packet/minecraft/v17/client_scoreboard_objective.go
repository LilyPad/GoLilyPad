package v17

import (
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
	packetClientScoreboardObjective.Value, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	packetClientScoreboardObjective.Action, err = packet.ReadInt8(reader)
	if err != nil {
		return
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
	err = packet.WriteString(writer, packetClientScoreboardObjective.Value)
	if err != nil {
		return
	}
	err = packet.WriteInt8(writer, packetClientScoreboardObjective.Action)
	return
}
