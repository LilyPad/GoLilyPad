package minecraft

type PacketServerLoginEncryptResponse struct {
	IdMapPacket
	SharedSecret []byte
	VerifyToken  []byte
}

func NewPacketServerLoginEncryptResponse(idMap *IdMap, sharedSecret []byte, verifyToken []byte) (this *PacketServerLoginEncryptResponse) {
	this = new(PacketServerLoginEncryptResponse)
	this.IdFrom(idMap)
	this.SharedSecret = sharedSecret
	this.VerifyToken = verifyToken
	return
}

func (this *PacketServerLoginEncryptResponse) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketServerLoginEncryptResponse
}
