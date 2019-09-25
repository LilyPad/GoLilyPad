package v19

import (
	"errors"
	"fmt"
	"github.com/LilyPad/GoLilyPad/packet"
	"github.com/LilyPad/GoLilyPad/packet/minecraft"
	uuid "github.com/satori/go.uuid"
	"io"
)

const (
	PACKET_CLIENT_BOSS_BAR_ACTION_ADD           = int(0)
	PACKET_CLIENT_BOSS_BAR_ACTION_REMOVE        = int(1)
	PACKET_CLIENT_BOSS_BAR_ACTION_UPDATE_HEALTH = int(2)
	PACKET_CLIENT_BOSS_BAR_ACTION_UPDATE_TITLE  = int(3)
	PACKET_CLIENT_BOSS_BAR_ACTION_UPDATE_STYLE  = int(4)
	PACKET_CLIENT_BOSS_BAR_ACTION_UPDATE_FLAGS  = int(5)
)

type PacketClientBossBar struct {
	minecraft.IdMapPacket
	UUID     uuid.UUID
	Action   int
	Title    string
	Health   float32
	Color    int
	Division int
	Flags    uint8
}

func NewPacketClientBossBarRemove(idMap *minecraft.IdMap, UUID uuid.UUID) (this *PacketClientBossBar) {
	this = new(PacketClientBossBar)
	this.IdFrom(idMap)
	this.UUID = UUID
	this.Action = PACKET_CLIENT_BOSS_BAR_ACTION_REMOVE
	return
}

func (this *PacketClientBossBar) IdFrom(idMap *minecraft.IdMap) {
	this.IdSet(idMap.PacketClientBossBar)
}

type CodecClientBossBar struct {
	IdMap *minecraft.IdMap
}

func (this *CodecClientBossBar) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetClientBossBar := new(PacketClientBossBar)
	packetClientBossBar.IdFrom(this.IdMap)
	packetClientBossBar.UUID, err = packet.ReadUUID(reader)
	if err != nil {
		return
	}
	packetClientBossBar.Action, err = packet.ReadVarInt(reader)
	if err != nil {
		return
	}
	switch packetClientBossBar.Action {
	case PACKET_CLIENT_BOSS_BAR_ACTION_ADD:
		packetClientBossBar.Title, err = packet.ReadString(reader)
		if err != nil {
			return
		}
		packetClientBossBar.Health, err = packet.ReadFloat32(reader)
		if err != nil {
			return
		}
		packetClientBossBar.Color, err = packet.ReadVarInt(reader)
		if err != nil {
			return
		}
		packetClientBossBar.Division, err = packet.ReadVarInt(reader)
		if err != nil {
			return
		}
		packetClientBossBar.Flags, err = packet.ReadUint8(reader)
	case PACKET_CLIENT_BOSS_BAR_ACTION_REMOVE:
	case PACKET_CLIENT_BOSS_BAR_ACTION_UPDATE_HEALTH:
		packetClientBossBar.Health, err = packet.ReadFloat32(reader)
	case PACKET_CLIENT_BOSS_BAR_ACTION_UPDATE_TITLE:
		packetClientBossBar.Title, err = packet.ReadString(reader)
	case PACKET_CLIENT_BOSS_BAR_ACTION_UPDATE_STYLE:
		packetClientBossBar.Color, err = packet.ReadVarInt(reader)
		if err != nil {
			return
		}
		packetClientBossBar.Division, err = packet.ReadVarInt(reader)
	case PACKET_CLIENT_BOSS_BAR_ACTION_UPDATE_FLAGS:
		packetClientBossBar.Flags, err = packet.ReadUint8(reader)
	default:
		err = errors.New(fmt.Sprintf("Decode, PacketClientBossBar action is not valid: %d", packetClientBossBar.Action))
	}
	if err != nil {
		return
	}
	decode = packetClientBossBar
	return
}

func (this *CodecClientBossBar) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetClientBossBar := encode.(*PacketClientBossBar)
	err = packet.WriteUUID(writer, packetClientBossBar.UUID)
	if err != nil {
		return
	}
	err = packet.WriteVarInt(writer, packetClientBossBar.Action)
	if err != nil {
		return
	}
	switch packetClientBossBar.Action {
	case PACKET_CLIENT_BOSS_BAR_ACTION_ADD:
		err = packet.WriteString(writer, packetClientBossBar.Title)
		if err != nil {
			return
		}
		err = packet.WriteFloat32(writer, packetClientBossBar.Health)
		if err != nil {
			return
		}
		err = packet.WriteVarInt(writer, packetClientBossBar.Color)
		if err != nil {
			return
		}
		err = packet.WriteVarInt(writer, packetClientBossBar.Division)
		if err != nil {
			return
		}
		err = packet.WriteUint8(writer, packetClientBossBar.Flags)
	case PACKET_CLIENT_BOSS_BAR_ACTION_REMOVE:
	case PACKET_CLIENT_BOSS_BAR_ACTION_UPDATE_HEALTH:
		err = packet.WriteFloat32(writer, packetClientBossBar.Health)
	case PACKET_CLIENT_BOSS_BAR_ACTION_UPDATE_TITLE:
		err = packet.WriteString(writer, packetClientBossBar.Title)
	case PACKET_CLIENT_BOSS_BAR_ACTION_UPDATE_STYLE:
		err = packet.WriteVarInt(writer, packetClientBossBar.Color)
		if err != nil {
			return
		}
		err = packet.WriteVarInt(writer, packetClientBossBar.Division)
	case PACKET_CLIENT_BOSS_BAR_ACTION_UPDATE_FLAGS:
		err = packet.WriteUint8(writer, packetClientBossBar.Flags)
	default:
		err = errors.New(fmt.Sprintf("Encode, PacketClientBossBar action is not valid: %d", packetClientBossBar.Action))
	}
	return
}
