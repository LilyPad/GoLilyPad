package connect

import "errors"
import "net"
import "sync"
import "sync/atomic"
import "time"
import "github.com/LilyPad/GoLilyPad/packet"
import "github.com/LilyPad/GoLilyPad/packet/connect"

type ConnectImpl struct {
	EventDispatcher
	conn net.Conn
	connCodec *packet.PacketConnCodec

	records map[int32]*RequestRecord
	recordsMutex *sync.Mutex
	sequenceId int32 
}

func NewConnect() Connect {
	return &ConnectImpl{
		recordsMutex: &sync.Mutex{},
	}
}

func (this *ConnectImpl) Connect(addr string) (err error) {
	this.Disconnect()
	this.conn, err = net.Dial("tcp", addr)
	if err != nil {
		return
	}
	this.recordsMutex.Lock()
	defer this.recordsMutex.Unlock()
	this.records = make(map[int32]*RequestRecord)
	this.connCodec = packet.NewPacketConnCodec(this.conn, NewCodec(this), 10 * time.Second)
	go this.connCodec.ReadConn(this)
	return
}

func (this *ConnectImpl) Disconnect() {
	this.recordsMutex.Lock()
	defer this.recordsMutex.Unlock()
	if this.records != nil {
		for _, record := range this.records {
			if record.callback == nil {
				continue
			}
			go record.callback(255, nil)
		}
	}
	if this.conn != nil {
		this.conn.Close()
	}
	this.records = nil
	this.conn = nil
}

func (this *ConnectImpl) Connected() bool {
	return this.conn != nil
}

func (this *ConnectImpl) Write(packet packet.Packet) (err error) {
	return this.connCodec.Write(packet)
}

func (this *ConnectImpl) HandlePacket(packet packet.Packet) (err error) {
	switch packet.Id() {
	case connect.PACKET_KEEPALIVE:
		this.Write(packet)
	case connect.PACKET_RESULT:
		packetResult := packet.(*connect.PacketResult)
		this.DispatchResult(packetResult.SequenceId, packetResult.StatusCode, packetResult.Result)
	case connect.PACKET_MESSAGE_EVENT:
		this.DispatchEvent("message", WrapEventMessage(packet.(*connect.PacketMessageEvent)))
	case connect.PACKET_REDIRECT_EVENT:
		this.DispatchEvent("redirect", WrapEventRedirect(packet.(*connect.PacketRedirectEvent)))
	case connect.PACKET_SERVER_EVENT:
		this.DispatchEvent("server", WrapEventServer(packet.(*connect.PacketServerEvent)))
	}
	return
}

func (this *ConnectImpl) ErrorCaught(err error) {
	this.Disconnect()
}

func (this *ConnectImpl) Request(request connect.Request) (statusCode uint8, result connect.Result, err error) {
	statusCodeChannel := make(chan uint8)
	resultChannel := make(chan connect.Result)
	err = this.RequestLater(request, func(statusCode uint8, result connect.Result) {
		statusCodeChannel <- statusCode
		resultChannel <- request
	})
	if err != nil {
		return
	}
	return <-statusCodeChannel, <-resultChannel, nil
}

func (this *ConnectImpl) RequestLater(request connect.Request, callback RequestCallback) (err error) {
	if !this.Connected() {
		return errors.New("Not connected")
	}
	sequenceId := atomic.AddInt32(&this.sequenceId, 1)
	err = this.Write(&connect.PacketRequest{sequenceId, request})
	if err != nil {
		return
	}
	this.recordsMutex.Lock()
	defer this.recordsMutex.Unlock()
	if this.records != nil {
		this.records[sequenceId] = &RequestRecord{request, callback}
	}
	return
}

func (this *ConnectImpl) DispatchResult(sequenceId int32, statusCode uint8, result connect.Result) {
	this.recordsMutex.Lock()
	defer this.recordsMutex.Unlock()
	if _, ok := this.records[sequenceId]; !ok {
		return // should there be an error here?
	}
	if this.records[sequenceId].callback != nil {
		go this.records[sequenceId].callback(statusCode, result)
	}
	delete(this.records, sequenceId)
}

func (this *ConnectImpl) RequestIdBySequenceId(sequenceId int32) int {
	this.recordsMutex.Lock()
	defer this.recordsMutex.Unlock()
	if _, ok := this.records[sequenceId]; !ok {
		return -1
	}
	return this.records[sequenceId].request.Id()
}
