package minecraft

type PacketServerLoginStart struct {
	IdMapPacket
	Name         string
	HasPlayerKey bool
	PlayerKey    *GameKey
}

func NewPacketServerLoginStart(idMap *IdMap, name string, playerKey *GameKey) (this *PacketServerLoginStart) {
	this = new(PacketServerLoginStart)
	this.IdFrom(idMap)
	this.Name = name
	this.PlayerKey = playerKey
	return
}

func (this *PacketServerLoginStart) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketServerLoginStart
}
