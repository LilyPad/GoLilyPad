package packet

import (
	"bytes"
)

type PacketIntercept func(Packet, *bytes.Buffer) error
