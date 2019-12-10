package minecraft

type PacketClientRespawn struct {
	IdMapPacket
	Dimension  int32
	HashedSeed int64
	Difficulty int8
	Gamemode   int8
	LevelType  string
}

func NewPacketClientRespawn(idMap *IdMap, dimension int32, difficulty int8, gamemode int8, levelType string) (this *PacketClientRespawn) {
	this = new(PacketClientRespawn)
	this.IdFrom(idMap)
	this.Dimension = dimension
	this.Difficulty = difficulty
	this.Gamemode = gamemode
	this.LevelType = levelType
	return
}

func (this *PacketClientRespawn) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketClientRespawn
}
