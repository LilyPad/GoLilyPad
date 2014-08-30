package minecraft

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type PacketClientJoinGame struct {
	EntityId int32
	Gamemode int8
	Dimension int8
	Difficulty int8
	MaxPlayers int8
	LevelType string
}

func NewPacketClientJoinGame(entityId int32, gamemode int8, dimension int8, difficulty int8, maxPlayers int8, levelType string) (this *PacketClientJoinGame) {
	this = new(PacketClientJoinGame)
	this.EntityId = entityId
	this.Gamemode = gamemode
	this.Dimension = dimension
	this.Difficulty = difficulty
	this.MaxPlayers = maxPlayers
	this.LevelType = levelType
	return
}

func (this *PacketClientJoinGame) Id() int {
	return PACKET_CLIENT_JOIN_GAME
}

type packetClientJoinGameCodec struct {

}

func (this *packetClientJoinGameCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetClientJoinGame := new(PacketClientJoinGame)
	packetClientJoinGame.EntityId, err = packet.ReadInt32(reader, util)
	if err != nil {
		return
	}
	packetClientJoinGame.Gamemode, err = packet.ReadInt8(reader, util)
	if err != nil {
		return
	}
	packetClientJoinGame.Dimension, err = packet.ReadInt8(reader, util)
	if err != nil {
		return
	}
	packetClientJoinGame.Difficulty, err = packet.ReadInt8(reader, util)
	if err != nil {
		return
	}
	packetClientJoinGame.MaxPlayers, err = packet.ReadInt8(reader, util)
	if err != nil {
		return
	}
	packetClientJoinGame.LevelType, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	decode = packetClientJoinGame
	return
}

func (this *packetClientJoinGameCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetClientJoinGame := encode.(*PacketClientJoinGame)
	err = packet.WriteInt32(writer, util, packetClientJoinGame.EntityId)
	if err != nil {
		return
	}
	err = packet.WriteInt8(writer, util, packetClientJoinGame.Gamemode)
	if err != nil {
		return
	}
	err = packet.WriteInt8(writer, util, packetClientJoinGame.Dimension)
	if err != nil {
		return
	}
	err = packet.WriteInt8(writer, util, packetClientJoinGame.Difficulty)
	if err != nil {
		return
	}
	err = packet.WriteInt8(writer, util, packetClientJoinGame.MaxPlayers)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, util, packetClientJoinGame.LevelType)
	return
}
