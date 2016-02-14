package minecraft

import (
	"io"
	"github.com/suedadam/GoLilyPad/packet"
)

type PacketClientJoinGame struct {
	EntityId int32
	Gamemode int8
	Dimension int8
	Difficulty int8
	MaxPlayers int8
	LevelType string
	ReducedDebugInfo bool
}

func NewPacketClientJoinGame(entityId int32, gamemode int8, dimension int8, difficulty int8, maxPlayers int8, levelType string, reducedDebugInfo bool) (this *PacketClientJoinGame) {
	this = new(PacketClientJoinGame)
	this.EntityId = entityId
	this.Gamemode = gamemode
	this.Dimension = dimension
	this.Difficulty = difficulty
	this.MaxPlayers = maxPlayers
	this.LevelType = levelType
	this.ReducedDebugInfo = reducedDebugInfo
	return
}

func (this *PacketClientJoinGame) Id() int {
	return PACKET_CLIENT_JOIN_GAME
}

type packetClientJoinGameCodec struct {

}

func (this *packetClientJoinGameCodec) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientJoinGame := new(PacketClientJoinGame)
	packetClientJoinGame.EntityId, err = packet.ReadInt32(reader)
	if err != nil {
		return
	}
	packetClientJoinGame.Gamemode, err = packet.ReadInt8(reader)
	if err != nil {
		return
	}
	packetClientJoinGame.Dimension, err = packet.ReadInt8(reader)
	if err != nil {
		return
	}
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

func (this *packetClientJoinGameCodec) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientJoinGame := encode.(*PacketClientJoinGame)
	err = packet.WriteInt32(writer, packetClientJoinGame.EntityId)
	if err != nil {
		return
	}
	err = packet.WriteInt8(writer, packetClientJoinGame.Gamemode)
	if err != nil {
		return
	}
	err = packet.WriteInt8(writer, packetClientJoinGame.Dimension)
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
