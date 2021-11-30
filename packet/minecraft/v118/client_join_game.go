package v118

import (
	"github.com/LilyPad/GoLilyPad/packet"
	"github.com/LilyPad/GoLilyPad/packet/minecraft"
	"github.com/LilyPad/GoLilyPad/packet/minecraft/nbt"
	"io"
)

type CodecClientJoinGame struct {
	IdMap *minecraft.IdMap
}

func (this *CodecClientJoinGame) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientJoinGame := new(minecraft.PacketClientJoinGame)
	packetClientJoinGame.IdFrom(this.IdMap)
	packetClientJoinGame.EntityId, err = packet.ReadInt32(reader)
	if err != nil {
		return
	}
	packetClientJoinGame.Hardcore, err = packet.ReadBool(reader)
	if err != nil {
		return
	}
	packetClientJoinGame.Gamemode, err = packet.ReadInt8(reader)
	if err != nil {
		return
	}
	packetClientJoinGame.PreviousGamemode, err = packet.ReadInt8(reader)
	if err != nil {
		return
	}
	worldNameLen, err := packet.ReadVarInt(reader)
	if err != nil {
		return
	}
	packetClientJoinGame.WorldNames = make([]string, worldNameLen)
	for i := range packetClientJoinGame.WorldNames {
		packetClientJoinGame.WorldNames[i], err = packet.ReadString(reader)
		if err != nil {
			return
		}
	}
	packetClientJoinGame.DimensionCodec, err = nbt.ReadNbt(reader)
	if err != nil {
		return
	}
	packetClientJoinGame.DimensionNBT, err = nbt.ReadNbt(reader)
	if err != nil {
		return
	}
	packetClientJoinGame.WorldName, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	packetClientJoinGame.HashedSeed, err = packet.ReadInt64(reader)
	if err != nil {
		return
	}
	packetClientJoinGame.MaxPlayers, err = packet.ReadVarInt(reader)
	if err != nil {
		return
	}
	packetClientJoinGame.ViewDistance, err = packet.ReadVarInt(reader)
	if err != nil {
		return
	}
	packetClientJoinGame.SimulationDistance, err = packet.ReadVarInt(reader)
	if err != nil {
		return
	}
	packetClientJoinGame.ReducedDebugInfo, err = packet.ReadBool(reader)
	if err != nil {
		return
	}
	packetClientJoinGame.EnableRespawnScreen, err = packet.ReadBool(reader)
	if err != nil {
		return
	}
	packetClientJoinGame.IsDebug, err = packet.ReadBool(reader)
	if err != nil {
		return
	}
	packetClientJoinGame.IsFlat, err = packet.ReadBool(reader)
	if err != nil {
		return
	}
	decode = packetClientJoinGame
	return
}

func (this *CodecClientJoinGame) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientJoinGame := encode.(*minecraft.PacketClientJoinGame)
	err = packet.WriteInt32(writer, packetClientJoinGame.EntityId)
	if err != nil {
		return
	}
	err = packet.WriteBool(writer, packetClientJoinGame.Hardcore)
	if err != nil {
		return
	}
	err = packet.WriteInt8(writer, packetClientJoinGame.Gamemode)
	if err != nil {
		return
	}
	err = packet.WriteInt8(writer, packetClientJoinGame.PreviousGamemode)
	if err != nil {
		return
	}
	err = packet.WriteVarInt(writer, len(packetClientJoinGame.WorldNames))
	if err != nil {
		return
	}
	for _, worldName := range packetClientJoinGame.WorldNames {
		err = packet.WriteString(writer, worldName)
		if err != nil {
			return
		}
	}
	err = nbt.WriteNbt(writer, packetClientJoinGame.DimensionCodec)
	if err != nil {
		return
	}
	err = nbt.WriteNbt(writer, packetClientJoinGame.DimensionNBT)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, packetClientJoinGame.WorldName)
	if err != nil {
		return
	}
	err = packet.WriteInt64(writer, packetClientJoinGame.HashedSeed)
	if err != nil {
		return
	}
	err = packet.WriteVarInt(writer, packetClientJoinGame.MaxPlayers)
	if err != nil {
		return
	}
	err = packet.WriteVarInt(writer, packetClientJoinGame.ViewDistance)
	if err != nil {
		return
	}
	err = packet.WriteVarInt(writer, packetClientJoinGame.SimulationDistance)
	if err != nil {
		return
	}
	err = packet.WriteBool(writer, packetClientJoinGame.ReducedDebugInfo)
	if err != nil {
		return
	}
	err = packet.WriteBool(writer, packetClientJoinGame.EnableRespawnScreen)
	if err != nil {
		return
	}
	err = packet.WriteBool(writer, packetClientJoinGame.IsDebug)
	if err != nil {
		return
	}
	err = packet.WriteBool(writer, packetClientJoinGame.IsFlat)
	return
}
