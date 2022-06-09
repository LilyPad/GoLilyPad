package minecraft

type PacketServerLoginEncryptResponse struct {
	IdMapPacket
	SharedSecret    []byte
	VerifyToken     []byte
	Salt            uint64
	DisableSaltAuth bool
}

func (this *PacketServerLoginEncryptResponse) IdFrom(idMap *IdMap) {
	this.IdMapPacket.id = idMap.PacketServerLoginEncryptResponse
}
