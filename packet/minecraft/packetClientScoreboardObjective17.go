package minecraft

import (
	"io"
	"github.com/suedadam/GoLilyPad/packet"
)

type packetClientScoreboardObjectiveCodec17 struct {

}

func (this *packetClientScoreboardObjectiveCodec17) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientScoreboardObjective := new(PacketClientScoreboardObjective)
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

func (this *packetClientScoreboardObjectiveCodec17) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientScoreboardObjective := encode.(*PacketClientScoreboardObjective)
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
