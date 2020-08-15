package connect

import (
	"github.com/LilyPad/GoLilyPad/packet"
	"io"
	"github.com/satori/go.uuid"
)

type PacketPlayerEvent struct {
	Joining    bool
	PlayerName string
	PlayerUUID uuid.UUID
}

func NewPacketPlayerEventJoin(playerName string, playerUUID uuid.UUID) (this *PacketPlayerEvent) {
	this = new(PacketPlayerEvent)
	this.Joining = true
	this.PlayerName = playerName
	this.PlayerUUID = playerUUID
	return
}

func NewPacketPlayerEventLeave(playerName string, playerUUID uuid.UUID) (this *PacketPlayerEvent) {
	this = new(PacketPlayerEvent)
	this.Joining = false
	this.PlayerName = playerName
	this.PlayerUUID = playerUUID
	return
}

func (this *PacketPlayerEvent) Id() int {
	return PACKET_PLAYER_EVENT
}

type packetPlayerEventCodec struct {
}

func (this *packetPlayerEventCodec) Decode(reader io.Reader) (decode packet.Packet, err error) {
	packetPlayerEvent := new(PacketPlayerEvent)
	packetPlayerEvent.Joining, err = packet.ReadBool(reader)
	if err != nil {
		return
	}
	packetPlayerEvent.PlayerName, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	packetPlayerEvent.PlayerUUID, err = packet.ReadUUID(reader)
	if err != nil {
		return
	}
	decode = packetPlayerEvent
	return
}

func (this *packetPlayerEventCodec) Encode(writer io.Writer, encode packet.Packet) (err error) {
	packetPlayerEvent := encode.(*PacketPlayerEvent)
	err = packet.WriteBool(writer, packetPlayerEvent.Joining)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, packetPlayerEvent.PlayerName)
	if err != nil {
		return
	}
	err = packet.WriteUUID(writer, packetPlayerEvent.PlayerUUID)
	return
	return
}
