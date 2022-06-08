package minecraft

type PacketServerLoginEncryptResponse struct {
	IdMapPacket
	SharedSecret []byte
	VerifyToken  []byte
	Salt         int64  // 1.19+
	Signature    []byte // 1.19+
}

func (this *PacketServerLoginEncryptResponse) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketServerLoginEncryptResponse
}
