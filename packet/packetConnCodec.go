package packet

import "io"
import "net"
import "time"
import "sync"

type PacketConnCodec struct {
	conn net.Conn
	reader io.Reader
	writer io.Writer
	packetCodec PacketCodec
	timeout time.Duration
	writeMutex *sync.Mutex
	writeUtil []byte
	readUtils []byte

}

func NewPacketConnCodec(conn net.Conn, packetCodec PacketCodec, timeout time.Duration) *PacketConnCodec {
	return &PacketConnCodec{
		conn: conn,
		reader: &FullReader{conn},
		writer: conn,
		packetCodec: packetCodec,
		timeout: timeout,
		writeMutex: &sync.Mutex{},
		writeUtil: make([]byte, UTIL_BUFFER_LENGTH),
		readUtil: make([]byte, UTIL_BUFFER_LENGTH),
	}
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

func (this *PacketConnCodec) SetReader(reader io.Reader) {
	this.reader = reader
}

func (this *PacketConnCodec) SetWriter(writer io.Writer) {
	this.writer = writer
}
