package packet

type PacketHandler interface {
	HandlePacket(packet Packet) (err error)
	ErrorCaught(err error)
}
