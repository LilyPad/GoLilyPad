package minecraft

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type packetClientScoreboardObjectiveCodec17 struct {

}

func (this *packetClientScoreboardObjectiveCodec17) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
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

func (this *packetClientScoreboardObjectiveCodec17) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
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
