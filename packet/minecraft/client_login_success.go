package minecraft

type PacketClientLoginSuccess struct {
	IdMapPacket
	UUID       string
	Name       string
	Properties []GameProfileProperty // 1.19+
}

func NewPacketClientLoginSuccess(idMap *IdMap, uuid string, name string, properties []GameProfileProperty) (this *PacketClientLoginSuccess) {
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
