package v119

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
	packetClientRespawn.DimensionName, err = packet.ReadString(reader)
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
	packetClientRespawn.HasLastDeathPosition, err = packet.ReadBool(reader)
	if err != nil {
		return
	}
	if packetClientRespawn.HasLastDeathPosition {
		packetClientRespawn.LastDeathPositionWorld, err = packet.ReadString(reader)
		if err != nil {
			return
		}
		packetClientRespawn.LastDeathPosition, err = packet.ReadUint64(reader)
		if err != nil {
			return
		}
	}
	decode = packetClientRespawn
	return
}

func (this *CodecClientRespawn) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientRespawn := encode.(*minecraft.PacketClientRespawn)
	err = packet.WriteString(writer, packetClientRespawn.DimensionName)
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
	if err != nil {
		return
	}
	err = packet.WriteBool(writer, packetClientRespawn.HasLastDeathPosition)
	if err != nil {
		return err
	}
	if packetClientRespawn.HasLastDeathPosition {
		err = packet.WriteString(writer, packetClientRespawn.LastDeathPositionWorld)
		if err != nil {
			return
		}
		err = packet.WriteUint64(writer, packetClientRespawn.LastDeathPosition)
		if err != nil {
			return
		}
	}
	return
}
