package v115

import (
	"github.com/LilyPad/GoLilyPad/packet"
	minecraft "github.com/LilyPad/GoLilyPad/packet/minecraft"
	"io"
)

type CodecClientRespawn struct {
	IdMap *minecraft.IdMap
}

func (this *CodecClientRespawn) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientRespawn := new(minecraft.PacketClientRespawn)
	packetClientRespawn.IdFrom(this.IdMap)
	packetClientRespawn.Dimension, err = packet.ReadInt32(reader)
	if err != nil {
		return
	}
	packetClientRespawn.HashedSeed, err = packet.ReadInt64(reader)
	if err != nil {
		return
	}
	packetClientRespawn.Gamemode, err = packet.ReadInt8(reader)
	if err != nil {
		return
	}
	packetClientRespawn.LevelType, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	decode = packetClientRespawn
	return
}

func (this *CodecClientRespawn) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientRespawn := encode.(*minecraft.PacketClientRespawn)
	err = packet.WriteInt32(writer, packetClientRespawn.Dimension)
	if err != nil {
		return
	}
	err = packet.WriteInt64(writer, packetClientRespawn.HashedSeed)
	if err != nil {
		return
	}
	err = packet.WriteInt8(writer, packetClientRespawn.Gamemode)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, packetClientRespawn.LevelType)
	return
}
