package v116

import (
	"github.com/LilyPad/GoLilyPad/packet"
	minecraft "github.com/LilyPad/GoLilyPad/packet/minecraft"
	uuidlib "github.com/satori/go.uuid"
	"io"
)

type CodecClientLoginSuccess struct {
	IdMap *minecraft.IdMap
}

func (this *CodecClientLoginSuccess) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientLoginSuccess := new(minecraft.PacketClientLoginSuccess)
	packetClientLoginSuccess.IdFrom(this.IdMap)
	uuid, err := packet.ReadUUID(reader)
	if err != nil {
		return
	}
	packetClientLoginSuccess.UUID = uuid.String()
	packetClientLoginSuccess.Name, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	decode = packetClientLoginSuccess
	return
}

func (this *CodecClientLoginSuccess) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientLoginSuccess := encode.(*minecraft.PacketClientLoginSuccess)
	uuid, err := uuidlib.FromString(packetClientLoginSuccess.UUID)
	if err != nil {
		return
	}
	err = packet.WriteUUID(writer, uuid)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, packetClientLoginSuccess.Name)
	return
}
