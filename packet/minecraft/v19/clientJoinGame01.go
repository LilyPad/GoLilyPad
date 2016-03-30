package v19

import (
	"github.com/LilyPad/GoLilyPad/packet"
	minecraft "github.com/LilyPad/GoLilyPad/packet/minecraft"
	"io"
)

type CodecClientJoinGame01 struct {
	IdMap *minecraft.IdMap
}

func (this *CodecClientJoinGame01) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientJoinGame := new(minecraft.PacketClientJoinGame)
	packetClientJoinGame.IdFrom(this.IdMap)
	packetClientJoinGame.EntityId, err = packet.ReadInt32(reader)
	if err != nil {
		return
	}
	packetClientJoinGame.Gamemode, err = packet.ReadInt8(reader)
	if err != nil {
		return
	}
	dimension32, err := packet.ReadInt32(reader)
	if err != nil {
		return
	}
	packetClientJoinGame.Dimension = int8(dimension32)
	packetClientJoinGame.Difficulty, err = packet.ReadInt8(reader)
	if err != nil {
		return
	}
	packetClientJoinGame.MaxPlayers, err = packet.ReadInt8(reader)
	if err != nil {
		return
	}
	packetClientJoinGame.LevelType, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	packetClientJoinGame.ReducedDebugInfo, err = packet.ReadBool(reader)
	if err != nil {
		return
	}
	decode = packetClientJoinGame
	return
}

func (this *CodecClientJoinGame01) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientJoinGame := encode.(*minecraft.PacketClientJoinGame)
	err = packet.WriteInt32(writer, packetClientJoinGame.EntityId)
	if err != nil {
		return
	}
	err = packet.WriteInt8(writer, packetClientJoinGame.Gamemode)
	if err != nil {
		return
	}
	err = packet.WriteInt32(writer, int32(packetClientJoinGame.Dimension))
	if err != nil {
		return
	}
	err = packet.WriteInt8(writer, packetClientJoinGame.Difficulty)
	if err != nil {
		return
	}
	err = packet.WriteInt8(writer, packetClientJoinGame.MaxPlayers)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, packetClientJoinGame.LevelType)
	if err != nil {
		return
	}
	err = packet.WriteBool(writer, packetClientJoinGame.ReducedDebugInfo)
	return
}
