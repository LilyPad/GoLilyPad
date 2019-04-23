package minecraft

type PacketClientJoinGame struct {
	IdMapPacket
	EntityId         int32
	Gamemode         int8
	Dimension        int8
	Difficulty       int8
	MaxPlayers       int8
	LevelType        string
	ViewDistance     int
	ReducedDebugInfo bool
}

func NewPacketClientJoinGame(idMap *IdMap, entityId int32, gamemode int8, dimension int8, difficulty int8, maxPlayers int8, levelType string, viewDistance int, reducedDebugInfo bool) (this *PacketClientJoinGame) {
	this = new(PacketClientJoinGame)
	this.IdFrom(idMap)
	this.EntityId = entityId
	this.Gamemode = gamemode
	this.Dimension = dimension
	this.Difficulty = difficulty
	this.MaxPlayers = maxPlayers
	this.LevelType = levelType
	this.ViewDistance = viewDistance
	this.ReducedDebugInfo = reducedDebugInfo
	return
}

func (this *PacketClientJoinGame) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketClientJoinGame
}
