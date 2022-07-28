package minecraft

import (
	uuid "github.com/satori/go.uuid"
)

type PacketServerLoginStart struct {
	IdMapPacket
	Name      string
	PublicKey *ProfilePublicKey
	UUID      *uuid.UUID // 1.19.1+
}

func NewPacketServerLoginStart(idMap *IdMap, name string, uuid *uuid.UUID) (this *PacketServerLoginStart) {
	this = new(PacketServerLoginStart)
	this.IdFrom(idMap)
	this.Name = name
	this.UUID = uuid
	return
}

func (this *PacketServerLoginStart) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketServerLoginStart
}
