package minecraft

type PacketClientSystemChatMessage struct {
	IdMapPacket
	Json     string
	Position int8
}

func NewPacketClientSystemChatMessage(idMap *IdMap, json string, position int8) (this *PacketClientSystemChatMessage) {
	this = new(PacketClientSystemChatMessage)
	this.IdFrom(idMap)
	this.Json = json
	this.Position = position
	return
}

func (this *PacketClientSystemChatMessage) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketClientSystemChatMessage
}
