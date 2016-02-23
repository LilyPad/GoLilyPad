package minecraft

type PacketClientLoginSetCompression struct {
	IdMapPacket
	Threshold int
}

func NewPacketClientLoginSetCompression(idMap *IdMap, threshold int) (this *PacketClientLoginSetCompression) {
	this = new(PacketClientLoginSetCompression)
	this.id = idMap.PacketClientLoginSetCompression
	this.Threshold = threshold
	return
}

func (this *PacketClientLoginSetCompression) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketClientLoginSetCompression
}
