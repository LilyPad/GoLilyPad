package minecraft

const (
	PACKET_CLIENT_TEAMS_ACTION_ADD            = int8(0)
	PACKET_CLIENT_TEAMS_ACTION_REMOVE         = int8(1)
	PACKET_CLIENT_TEAMS_ACTION_INFO_UPDATE    = int8(2)
	PACKET_CLIENT_TEAMS_ACTION_PLAYERS_ADD    = int8(3)
	PACKET_CLIENT_TEAMS_ACTION_PLAYERS_REMOVE = int8(4)
)

type PacketClientTeams struct {
	IdMapPacket
	Name              string
	Action            int8
	DisplayName       string
	Prefix            string
	Suffix            string
	FriendlyFire      int8
	NameTagVisibility string
	CollisionRule     string
	Color             int
	Players           []string
}

func NewPacketClientTeamsAdd(idMap *IdMap, name string, displayName string, prefix string, suffix string, friendlyFire int8, nameTagVisibility string, color int, players []string) (this *PacketClientTeams) {
	this = new(PacketClientTeams)
	this.IdFrom(idMap)
	this.Name = name
	this.Action = PACKET_CLIENT_TEAMS_ACTION_ADD
	this.DisplayName = displayName
	this.Prefix = prefix
	this.Suffix = suffix
	this.FriendlyFire = friendlyFire
	this.NameTagVisibility = nameTagVisibility
	this.Color = color
	this.Players = players
	return
}

func NewPacketClientTeamsRemove(idMap *IdMap, name string) (this *PacketClientTeams) {
	this = new(PacketClientTeams)
	this.IdFrom(idMap)
	this.Name = name
	this.Action = PACKET_CLIENT_TEAMS_ACTION_REMOVE
	return
}

func NewPacketClientTeamsInfoUpdate(idMap *IdMap, name string, displayName string, prefix string, suffix string, friendlyFire int8, nameTagVisibility string, collisionRule string, color int) (this *PacketClientTeams) {
	this = new(PacketClientTeams)
	this.IdFrom(idMap)
	this.Name = name
	this.Action = PACKET_CLIENT_TEAMS_ACTION_INFO_UPDATE
	this.DisplayName = displayName
	this.Prefix = prefix
	this.Suffix = suffix
	this.FriendlyFire = friendlyFire
	this.NameTagVisibility = nameTagVisibility
	this.CollisionRule = collisionRule
	this.Color = color
	return
}

func NewPacketClientTeamsPlayersAdd(idMap *IdMap, name string, players []string) (this *PacketClientTeams) {
	this = new(PacketClientTeams)
	this.IdFrom(idMap)
	this.Name = name
	this.Action = PACKET_CLIENT_TEAMS_ACTION_PLAYERS_ADD
	this.Players = players
	return
}

func NewPacketClientTeamsPlayersRemove(idMap *IdMap, name string, players []string) (this *PacketClientTeams) {
	this = new(PacketClientTeams)
	this.IdFrom(idMap)
	this.Name = name
	this.Action = PACKET_CLIENT_TEAMS_ACTION_PLAYERS_REMOVE
	this.Players = players
	return
}

func (this *PacketClientTeams) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketClientTeams
}
