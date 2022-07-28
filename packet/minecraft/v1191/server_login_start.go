package v119

import (
	"github.com/LilyPad/GoLilyPad/packet"
	minecraft "github.com/LilyPad/GoLilyPad/packet/minecraft"
	"github.com/satori/go.uuid"
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
	publicKeyPresent, err := packet.ReadBool(reader)
	if err != nil {
		return
	}
	if publicKeyPresent {
		packetServerLoginStart.PublicKey, err = minecraft.ReadProfilePublicKey(reader)
		if err != nil {
			return
		}
	}
	uuidPresent, err := packet.ReadBool(reader)
	if err != nil {
		return
	}
	if uuidPresent {
		var uuid uuid.UUID
		uuid, err = packet.ReadUUID(reader)
		if err != nil {
			return
		}
		packetServerLoginStart.UUID = &uuid
	}
	decode = packetServerLoginStart
	return
}

func (this *CodecServerLoginStart) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetServerLoginStart := encode.(*minecraft.PacketServerLoginStart)
	err = packet.WriteString(writer, packetServerLoginStart.Name)
	if err != nil {
		return
	}
	if packetServerLoginStart.PublicKey == nil {
		err = packet.WriteBool(writer, false)
		if err != nil {
			return
		}
	} else {
		err = packet.WriteBool(writer, true)
		if err != nil {
			return
		}
		err = minecraft.WriteProfilePublicKey(writer, packetServerLoginStart.PublicKey)
	}
	if packetServerLoginStart.UUID == nil {
		err = packet.WriteBool(writer, false)
	} else {
		err = packet.WriteBool(writer, true)
		if err != nil {
			return
		}
		err = packet.WriteUUID(writer, *packetServerLoginStart.UUID)
	}
	return
}
