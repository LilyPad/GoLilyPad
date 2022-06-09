package minecraft

import "github.com/LilyPad/GoLilyPad/auth"

type PacketClientLoginSuccess struct {
	IdMapPacket
	UUID       string
	Name       string
	Properties []auth.GameProfileProperty
}

func NewPacketClientLoginSuccess(idMap *IdMap, uuid string, name string, properties []auth.GameProfileProperty) (this *PacketClientLoginSuccess) {
	this = new(PacketClientLoginSuccess)
	this.IdFrom(idMap)
	this.UUID = uuid
	this.Name = name
	this.Properties = properties
	return
}

func (this *PacketClientLoginSuccess) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketClientLoginSuccess
}
