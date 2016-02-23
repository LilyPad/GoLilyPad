package minecraft

type PacketClientLoginSuccess struct {
	IdMapPacket
	UUID string
	Name string
}

func NewPacketClientLoginSuccess(idMap *IdMap, uuid string, name string) (this *PacketClientLoginSuccess) {
	this = new(PacketClientLoginSuccess)
	this.IdFrom(idMap)
	this.UUID = uuid
	this.Name = name
	return
}

func (this *PacketClientLoginSuccess) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketClientLoginSuccess
}
