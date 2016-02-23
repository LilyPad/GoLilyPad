package minecraft

type PacketClientDisconnect struct {
	IdMapPacket
	Json string
}

func NewPacketClientDisconnect(idMap *IdMap, json string) (this *PacketClientDisconnect) {
	this = new(PacketClientDisconnect)
	this.IdFrom(idMap)
	this.Json = json
	return
}

func (this *PacketClientDisconnect) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketClientDisconnect
}
