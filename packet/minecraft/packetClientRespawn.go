package minecraft

import (
	"github.com/LilyPad/GoLilyPad/packet"
	"io"
)

type PacketClientRespawn struct {
	Dimension  int32
	Difficulty int8
	Gamemode   int8
	LevelType  string
}

func NewPacketClientRespawn(dimension int32, difficulty int8, gamemode int8, levelType string) (this *PacketClientRespawn) {
	this = new(PacketClientRespawn)
	this.Dimension = dimension
	this.Difficulty = difficulty
	this.Gamemode = gamemode
	this.LevelType = levelType
	return
}

func (this *PacketClientRespawn) Id() int {
	return PACKET_CLIENT_RESPAWN
}

type packetClientRespawnCodec struct {
}

func (this *packetClientRespawnCodec) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientRespawn := new(PacketClientRespawn)
	packetClientRespawn.Dimension, err = packet.ReadInt32(reader)
	if err != nil {
		return
	}
	packetClientRespawn.Difficulty, err = packet.ReadInt8(reader)
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

func (this *packetClientRespawnCodec) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientRespawn := encode.(*PacketClientRespawn)
	err = packet.WriteInt32(writer, packetClientRespawn.Dimension)
	if err != nil {
		return
	}
	err = packet.WriteInt8(writer, packetClientRespawn.Difficulty)
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
