package minecraft

import (
	uuid "github.com/satori/go.uuid"
)

const (
	PACKET_CLIENT_PLAYER_LIST_ACTION_ADD                 = 0
	PACKET_CLIENT_PLAYER_LIST_ACTION_UPDATE_GAMEMODE     = 1
	PACKET_CLIENT_PLAYER_LIST_ACTION_UPDATE_LATENCY      = 2
	PACKET_CLIENT_PLAYER_LIST_ACTION_UPDATE_DISPLAY_NAME = 3
	PACKET_CLIENT_PLAYER_LIST_ACTION_REMOVE              = 4
)

type PacketClientPlayerList struct {
	IdMapPacket
	Action int
	Items  []PacketClientPlayerListItem
}

type PacketClientPlayerListItem struct {
	UUID uuid.UUID
	Info interface{}
}

type PacketClientPlayerListAddPlayer struct {
	Name        string
	Properties  []PacketClientPlayerListAddPlayerProperty
	Gamemode    int
	Latency     int
	DisplayName string
}

type PacketClientPlayerListAddPlayerProperty struct {
	Name      string
	Value     string
	Signature string
}

type PacketClientPlayerListUpdateGamemode struct {
	Gamemode int
}

type PacketClientPlayerListUpdateLatency struct {
	Latency int
}

type PacketClientPlayerListUpdateDisplayName struct {
	DisplayName string
}

func NewPacketClientPlayerList(idMap *IdMap, action int, items []PacketClientPlayerListItem) (this *PacketClientPlayerList) {
	this = new(PacketClientPlayerList)
	this.IdFrom(idMap)
	this.Action = action
	this.Items = items
	return
}

func (this *PacketClientPlayerList) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketClientPlayerList
}
