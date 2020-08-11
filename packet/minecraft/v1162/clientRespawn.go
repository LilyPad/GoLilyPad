package v1162

import (
	"github.com/LilyPad/GoLilyPad/packet"
	minecraft "github.com/LilyPad/GoLilyPad/packet/minecraft"
	"github.com/LilyPad/GoLilyPad/packet/minecraft/nbt"
	"io"
)

type CodecClientRespawn struct {
	IdMap *minecraft.IdMap
}

func (this *CodecClientRespawn) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientRespawn := new(minecraft.PacketClientRespawn)
	packetClientRespawn.IdFrom(this.IdMap)
	packetClientRespawn.DimensionNBT, err = nbt.ReadNbt(reader)
	if err != nil {
		return
	}
	packetClientRespawn.WorldName, err = packet.ReadString(reader)
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
	packetClientRespawn.PreviousGamemode, err = packet.ReadInt8(reader)
	if err != nil {
		return
	}
	packetClientRespawn.IsDebug, err = packet.ReadBool(reader)
	if err != nil {
		return
	}
	packetClientRespawn.IsFlat, err = packet.ReadBool(reader)
	if err != nil {
		return
	}
	packetClientRespawn.CopyMetadata, err = packet.ReadBool(reader)
	if err != nil {
		return
	}
	decode = packetClientRespawn
	return
}

func (this *CodecClientRespawn) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientRespawn := encode.(*minecraft.PacketClientRespawn)
	err = nbt.WriteNbt(writer, packetClientRespawn.DimensionNBT)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, packetClientRespawn.WorldName)
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
	err = packet.WriteInt8(writer, packetClientRespawn.PreviousGamemode)
	if err != nil {
		return
	}
	err = packet.WriteBool(writer, packetClientRespawn.IsDebug)
	if err != nil {
		return
	}
	err = packet.WriteBool(writer, packetClientRespawn.IsFlat)
	if err != nil {
		return
	}
	err = packet.WriteBool(writer, packetClientRespawn.CopyMetadata)
	return
}
