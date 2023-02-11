package minecraft

import (
	uuid "github.com/satori/go.uuid"
)

type PacketClientPlayerInfoRemove struct {
	IdMapPacket
	UUIDs []uuid.UUID
}

func (this *PacketClientPlayerInfoRemove) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketClientPlayerInfoRemove
}

func NewPacketClientPlayerInfoRemove(idMap *IdMap, uuids []uuid.UUID) (this *PacketClientPlayerInfoRemove) {
	this = new(PacketClientPlayerInfoRemove)
	this.IdFrom(idMap)
	this.UUIDs = uuids
	return
}
