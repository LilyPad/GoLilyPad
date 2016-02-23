package packet

import (
	"io"
	"net"
	"sync"
	"time"
)

type PacketConnCodec struct {
	reader      io.Reader
	writer      io.Writer
	conn        net.Conn
	packetCodec PacketCodec
	timeout     time.Duration
	writeMutex  sync.Mutex
}

func NewPacketConnCodec(conn net.Conn, packetCodec PacketCodec, timeout time.Duration) (this *PacketConnCodec) {
	this = new(PacketConnCodec)
	this.reader = NewFullReader(conn)
	this.writer = conn
	this.conn = conn
	this.packetCodec = packetCodec
	this.timeout = timeout
	return
}

func (this *PacketConnCodec) Write(packet Packet) (err error) {
	this.writeMutex.Lock()
	err = this.packetCodec.Encode(this.writer, packet)
	this.writeMutex.Unlock()
	return
}

func (this *PacketConnCodec) ReadConn(packetHandler PacketHandler) {
	for {
		if this.timeout != -1 {
			this.conn.SetReadDeadline(time.Now().Add(this.timeout))
		}
		packet, err := this.packetCodec.Decode(this.reader)
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
