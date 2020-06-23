package minecraft

import (
	"github.com/LilyPad/GoLilyPad/packet/minecraft/nbt"
)

type PacketClientJoinGame struct {
	IdMapPacket
	EntityId            int32
	Gamemode            int8
	PreviousGamemode    int8     // 1.16+
	WorldNames          []string // 1.16+
	DimensionCodec      nbt.Nbt  // 1.16+
	Dimension           int8     // removed in 1.16+
	DimensionName       string   // 1.16+
	WorldName           string   // 1.16+
	HashedSeed          int64
	Difficulty          int8
	MaxPlayers          int8
	LevelType           string // removed in 1.16+
	ViewDistance        int
	ReducedDebugInfo    bool
	EnableRespawnScreen bool
	IsDebug             bool // 1.16+
	IsFlat              bool // 1.16+
}

func (this *PacketClientJoinGame) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketClientJoinGame
}
