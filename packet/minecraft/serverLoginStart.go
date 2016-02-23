package minecraft

type PacketServerLoginStart struct {
	IdMapPacket
	Name string
}

func NewPacketServerLoginStart(idMap *IdMap, name string) (this *PacketServerLoginStart) {
	this = new(PacketServerLoginStart)
	this.IdFrom(idMap)
	this.Name = name
	return
}

func (this *PacketServerLoginStart) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketServerLoginStart
}
