package packet

import (
	"io"
)

type PacketCodec interface {
	Decode(reader io.Reader) (packet Packet, err error)
	Encode(writer io.Writer, packet Packet) (err error)
}
