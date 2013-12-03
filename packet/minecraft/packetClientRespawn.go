package minecraft

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type PacketClientRespawn struct {
	Dimension int32
	Difficulty int8
	Gamemode int8
	LevelType string
}

func (this *PacketClientRespawn) Id() int {
	return PACKET_CLIENT_RESPAWN
}

type PacketClientRespawnCodec struct {
	
}

func (this *PacketClientRespawnCodec) Decode(reader io.Reader, util []byte) (decode packet.Packet, err error) {
	packetClientRespawn := &PacketClientRespawn{}
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
	return packetClientRespawn, nil
}

func (this *PacketClientRespawnCodec) Encode(writer io.Writer, util []byte, encode packet.Packet) (err error) {
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