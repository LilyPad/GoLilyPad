package minecraft

const (
	PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_ADD    = int8(0)
	PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_REMOVE = int8(1)
	PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_UPDATE = int8(2)
)

type PacketClientScoreboardObjective struct {
	IdMapPacket
	Name   string
	Action int8
	Value  string
	Type   string
}

func NewPacketClientScoreboardObjectiveAdd(idMap *IdMap, name string, value string, stype string) (this *PacketClientScoreboardObjective) {
	this = new(PacketClientScoreboardObjective)
	this.IdFrom(idMap)
	this.Name = name
	this.Action = PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_ADD
	this.Value = value
	this.Type = stype
	return
}

func NewPacketClientScoreboardObjectiveRemove(idMap *IdMap, name string) (this *PacketClientScoreboardObjective) {
	this = new(PacketClientScoreboardObjective)
	this.IdFrom(idMap)
	this.Name = name
	this.Action = PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_REMOVE
	return
}

func NewPacketClientScoreboardObjectiveUpdate(idMap *IdMap, name string, value string, stype string) (this *PacketClientScoreboardObjective) {
	this = new(PacketClientScoreboardObjective)
	this.IdFrom(idMap)
	this.Name = name
	this.Value = value
	this.Action = PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_UPDATE
	this.Type = stype
	return
}

func (this *PacketClientScoreboardObjective) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketClientScoreboardObjective
}
