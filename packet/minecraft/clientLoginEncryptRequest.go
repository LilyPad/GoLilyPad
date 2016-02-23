package minecraft

type PacketClientLoginEncryptRequest struct {
	IdMapPacket
	ServerId    string
	PublicKey   []byte
	VerifyToken []byte
}

func NewPacketClientLoginEncryptRequest(idMap *IdMap, serverId string, publicKey []byte, verifyToken []byte) (this *PacketClientLoginEncryptRequest) {
	this = new(PacketClientLoginEncryptRequest)
	this.IdFrom(idMap)
	this.ServerId = serverId
	this.PublicKey = publicKey
	this.VerifyToken = verifyToken
	return
}

func (this *PacketClientLoginEncryptRequest) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketClientLoginEncryptRequest
}
