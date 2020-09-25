package packet

type PacketPipelineChild interface {
	PacketCodec
	SetCodec(codec PacketCodec)
}
