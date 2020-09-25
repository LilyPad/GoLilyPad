package minecraft

type PacketClientSetCompression struct {
	IdMapPacket
	Threshold int
}

func NewPacketClientSetCompression(idMap *IdMap, threshold int) (this *PacketClientSetCompression) {
	this = new(PacketClientSetCompression)
	this.IdFrom(idMap)
	this.Threshold = threshold
	return
}

func (this *PacketClientSetCompression) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketClientSetCompression
}
