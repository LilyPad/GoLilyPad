package packet

import (
	"io"
	"net"
	"time"
	"sync"
)

type PacketConnCodec struct {
	reader io.Reader
	writer io.Writer
	conn net.Conn
	packetCodec PacketCodec
	timeout time.Duration
	writeMutex sync.Mutex
	writeUtil []byte
	readUtil []byte
}

func NewPacketConnCodec(conn net.Conn, packetCodec PacketCodec, timeout time.Duration) (this *PacketConnCodec) {
	this = new(PacketConnCodec)
	this.reader = NewFullReader(conn)
	this.writer = conn
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
	err = this.packetCodec.Encode(this.writer, this.writeUtil, packet)
	return
}

func (this *PacketConnCodec) ReadConn(packetHandler PacketHandler) {
	for {
		if this.timeout != -1 {
			this.conn.SetReadDeadline(time.Now().Add(this.timeout))
		}
		packet, err := this.packetCodec.Decode(this.reader, this.readUtil)
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

func (this *PacketConnCodec) SetTimeout(timeout time.Duration) {
	this.timeout = timeout
}
