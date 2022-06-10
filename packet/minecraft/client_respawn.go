package minecraft

import (
	"github.com/LilyPad/GoLilyPad/packet/minecraft/nbt"
)

type PacketClientRespawn struct {
	IdMapPacket
	Dimension        int32   // removed in 1.16+
	DimensionName    string  // 1.16+ // removed in 1.16.2+ // added back in 1.19+
	DimensionNBT     nbt.Nbt // 1.16.2+ // removed in 1.19+
	WorldName        string  // 1.16+
	HashedSeed       int64
	Difficulty       int8
	Gamemode         int8
	PreviousGamemode int8      // 1.16+
	LevelType        string    // removed in 1.16+
	IsDebug          bool      // 1.16+
	IsFlat           bool      // 1.16+
	DeathLocation    *Location // 1.19+
	CopyMetadata     bool      // 1.16+
}

func NewPacketClientRespawnFrom(idMap *IdMap, joinGame *PacketClientJoinGame) (this *PacketClientRespawn) {
	this = new(PacketClientRespawn)
	this.IdFrom(idMap)
	this.Dimension = int32(joinGame.Dimension)
	this.DimensionName = joinGame.DimensionName
	this.DimensionNBT = joinGame.DimensionNBT
	this.WorldName = joinGame.WorldName
	this.HashedSeed = joinGame.HashedSeed
	this.Difficulty = joinGame.Difficulty
	this.Gamemode = joinGame.Gamemode
	this.PreviousGamemode = joinGame.PreviousGamemode
	this.LevelType = joinGame.LevelType
	this.IsDebug = joinGame.IsDebug
	this.IsFlat = joinGame.IsFlat
	this.DeathLocation = joinGame.DeathLocation
	this.CopyMetadata = false
	return
}

func (this *PacketClientRespawn) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketClientRespawn
}
