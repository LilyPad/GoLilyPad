package minecraft

type PacketServerClientSettings struct {
	IdMapPacket
	Locale       string
	ViewDistance byte
	ChatFlags    byte
	ChatColours  bool
	Difficulty   byte
	SkinParts    byte
	MainHand     int
}

func NewPacketServerClientSettings(idMap *IdMap, locale string, viewDistance byte, chatFlags byte, chatColours bool, skinParts byte, mainHand int) (this *PacketServerClientSettings) {
	this = new(PacketServerClientSettings)
	this.IdFrom(idMap)
	this.Locale = locale
	this.ViewDistance = viewDistance
	this.ChatFlags = chatFlags
	this.ChatColours = chatColours
	this.SkinParts = skinParts
	this.MainHand = mainHand
	return
}

func (this *PacketServerClientSettings) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketServerClientSettings
}
