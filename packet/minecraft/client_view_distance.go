package minecraft

type PacketClientViewDistance struct {
	IdMapPacket
	ViewDistance int
}

func NewPacketClientViewDistance(idMap *IdMap, viewDistance int) (this *PacketClientViewDistance) {
	this = new(PacketClientViewDistance)
	this.IdFrom(idMap)
	this.ViewDistance = viewDistance
	return
}

func (this *PacketClientViewDistance) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketClientUpdateViewDistance
}
