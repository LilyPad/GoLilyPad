package v18

import (
	"github.com/LilyPad/GoLilyPad/packet"
	minecraft "github.com/LilyPad/GoLilyPad/packet/minecraft"
	"io"
)

type CodecServerLoginStart struct {
	IdMap *minecraft.IdMap
}

func (this *CodecServerLoginStart) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetServerLoginStart := new(minecraft.PacketServerLoginStart)
	packetServerLoginStart.IdFrom(this.IdMap)
	packetServerLoginStart.Name, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	decode = packetServerLoginStart
	return
}

func (this *CodecServerLoginStart) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetServerLoginStart := encode.(*minecraft.PacketServerLoginStart)
	err = packet.WriteString(writer, packetServerLoginStart.Name)
	return
}
