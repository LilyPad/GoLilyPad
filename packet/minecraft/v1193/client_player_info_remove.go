package v1193

import (
	"github.com/LilyPad/GoLilyPad/packet"
	"github.com/LilyPad/GoLilyPad/packet/minecraft"
	uuid "github.com/satori/go.uuid"
	"io"
)

type CodecClientPlayerInfoRemove struct {
	IdMap *minecraft.IdMap
}

func (this *CodecClientPlayerInfoRemove) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientPlayerInfoRemove := new(minecraft.PacketClientPlayerInfoRemove)
	packetClientPlayerInfoRemove.IdFrom(this.IdMap)
	uuidsLength, err := packet.ReadVarInt(reader)
	if err != nil {
		return
	}
	packetClientPlayerInfoRemove.UUIDs = make([]uuid.UUID, uuidsLength)
	for i := 0; i < uuidsLength; i++ {
		packetClientPlayerInfoRemove.UUIDs[i], err = packet.ReadUUID(reader)
		if err != nil {
			return
		}
	}
	decode = packetClientPlayerInfoRemove
	return
}

func (this *CodecClientPlayerInfoRemove) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientPlayerInfoRemove := encode.(*minecraft.PacketClientPlayerInfoRemove)
	err = packet.WriteVarInt(writer, len(packetClientPlayerInfoRemove.UUIDs))
	if err != nil {
		return
	}
	for _, uuid := range packetClientPlayerInfoRemove.UUIDs {
		err = packet.WriteUUID(writer, uuid)
		if err != nil {
			return
		}
	}
	return
}
