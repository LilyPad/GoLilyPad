package minecraft

type PacketClientLoginDisconnect struct {
	IdMapPacket
	Json string
}

func NewPacketClientLoginDisconnect(idMap *IdMap, json string) (this *PacketClientLoginDisconnect) {
	this = new(PacketClientLoginDisconnect)
	this.IdFrom(idMap)
	this.Json = json
	return
}

func (this *PacketClientLoginDisconnect) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketClientLoginDisconnect
}
