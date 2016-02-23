package minecraft

type PacketServerPluginMessage struct {
	IdMapPacket
	Channel string
	Data    []byte
}

func NewPacketServerPluginMessage(idMap *IdMap, channel string, data []byte) (this *PacketServerPluginMessage) {
	this = new(PacketServerPluginMessage)
	this.IdFrom(idMap)
	this.Channel = channel
	this.Data = data
	return
}

func (this *PacketServerPluginMessage) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketServerPluginMessage
}
