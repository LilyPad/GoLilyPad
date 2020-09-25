package v18

import (
	"github.com/LilyPad/GoLilyPad/packet"
	minecraft "github.com/LilyPad/GoLilyPad/packet/minecraft"
	"io"
)

type CodecClientLoginSuccess struct {
	IdMap *minecraft.IdMap
}

func (this *CodecClientLoginSuccess) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientLoginSuccess := new(minecraft.PacketClientLoginSuccess)
	packetClientLoginSuccess.IdFrom(this.IdMap)
	packetClientLoginSuccess.UUID, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	packetClientLoginSuccess.Name, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	decode = packetClientLoginSuccess
	return
}

func (this *CodecClientLoginSuccess) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientLoginSuccess := encode.(*minecraft.PacketClientLoginSuccess)
	err = packet.WriteString(writer, packetClientLoginSuccess.UUID)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, packetClientLoginSuccess.Name)
	return
}
