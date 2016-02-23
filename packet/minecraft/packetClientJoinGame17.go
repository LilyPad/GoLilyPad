package minecraft

import (
	"github.com/LilyPad/GoLilyPad/packet"
	"io"
)

type packetClientJoinGameCodec17 struct {
}

func (this *packetClientJoinGameCodec17) Decode(reader io.Reader) (decode packet.Packet, err error) {
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
	decode = packetClientJoinGame
	return
}

func (this *packetClientJoinGameCodec17) Encode(writer io.Writer, encode packet.Packet) (err error) {
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
	return
}
