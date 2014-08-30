package minecraft

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type PacketClientRespawn struct {
	Dimension int32
	Difficulty int8
	Gamemode int8
	LevelType string
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

func (this *packetClientRespawnCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetClientRespawn := new(PacketClientRespawn)
	packetClientRespawn.Dimension, err = packet.ReadInt32(reader, util)
	if err != nil {
		return
	}
	packetClientRespawn.Difficulty, err = packet.ReadInt8(reader, util)
	if err != nil {
		return
	}
	packetClientRespawn.Gamemode, err = packet.ReadInt8(reader, util)
	if err != nil {
		return
	}
	packetClientRespawn.LevelType, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	decode = packetClientRespawn
	return
}

func (this *packetClientRespawnCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
	packetClientRespawn := encode.(*PacketClientRespawn)
	err = packet.WriteInt32(writer, util, packetClientRespawn.Dimension)
	if err != nil {
		return
	}
	err = packet.WriteInt8(writer, util, packetClientRespawn.Difficulty)
	if err != nil {
		return
	}
	err = packet.WriteInt8(writer, util, packetClientRespawn.Gamemode)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, util, packetClientRespawn.LevelType)
	return
}
