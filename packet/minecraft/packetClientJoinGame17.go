package minecraft

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type packetClientJoinGameCodec17 struct {

}

func (this *packetClientJoinGameCodec17) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
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

func (this *packetClientJoinGameCodec17) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
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