package packet

import (
	"io"
	"net"
	"time"
	"sync"
)

type PacketConnCodec struct {
	Reader io.Reader
	Writer io.Writer
	conn net.Conn
	packetCodec PacketCodec
	timeout time.Duration
	writeMutex sync.Mutex
	writeUtil []byte
	readUtil []byte
}

func NewPacketConnCodec(conn net.Conn, packetCodec PacketCodec, timeout time.Duration) (this *PacketConnCodec) {
	this = new(PacketConnCodec)
	this.Reader = NewFullReader(conn)
	this.Writer = conn
	this.conn = conn
	this.packetCodec = packetCodec
	this.timeout = timeout
	this.writeUtil = make([]byte, UTIL_BUFFER_LENGTH)
	this.readUtil = make([]byte, UTIL_BUFFER_LENGTH)
	return
}

func (this *PacketConnCodec) Write(packet Packet) (err error) {
	this.writeMutex.Lock()
	defer this.writeMutex.Unlock()
	err = this.packetCodec.Encode(this.Writer, this.writeUtil, packet)
	return
}

func (this *PacketConnCodec) ReadConn(packetHandler PacketHandler) {
	for {
		if this.timeout != -1 {
			this.conn.SetReadDeadline(time.Now().Add(this.timeout))
		}
		packet, err := this.packetCodec.Decode(this.Reader, this.readUtil)
		if err != nil {
			packetHandler.ErrorCaught(err)
			return
		}
		err = packetHandler.HandlePacket(packet)
		if err != nil {
			packetHandler.ErrorCaught(err)
			return
		}
	}
}
