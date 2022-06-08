package minecraft

import (
	"github.com/LilyPad/GoLilyPad/packet/minecraft/nbt"
)

type PacketClientJoinGame struct {
	IdMapPacket
	EntityId            int32
	Hardcore            bool // 1.16.2+
	Gamemode            int8
	PreviousGamemode    int8     // 1.16+
	WorldNames          []string // 1.16+
	DimensionCodec      nbt.Nbt  // 1.16+
	DimensionNBT        nbt.Nbt  // 1.16.2+ // removed in 1.19+
	Dimension           int8     // removed in 1.16+
	DimensionName       string   // 1.16+ // removed in 1.16.2+ // added back in 1.19+
	WorldName           string   // 1.16+
	HashedSeed          int64
	Difficulty          int8
	MaxPlayers          int
	LevelType           string // removed in 1.16+
	ViewDistance        int
	ReducedDebugInfo    bool
	EnableRespawnScreen bool
	IsDebug             bool      // 1.16+
	IsFlat              bool      // 1.16+
	SimulationDistance  int       // 1.18+
	DeathLocation       *Location // 1.19+
}

func (this *PacketClientJoinGame) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketClientJoinGame
}
