package connect

import (
	"errors"
	"fmt"
	"github.com/LilyPad/GoLilyPad/packet"
	"github.com/LilyPad/GoLilyPad/packet/connect"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

type ConnectImpl struct {
	EventDispatcher
	conn      net.Conn
	connCodec *packet.PacketConnCodec
	pipeline  *packet.PacketPipeline

	records      map[int32]*RequestRecord
	recordsMutex sync.Mutex
	sequenceId   int32
}

func NewConnectImpl() (this *ConnectImpl) {
	this = new(ConnectImpl)
	return
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
	this.pipeline = packet.NewPacketPipeline()
	this.pipeline.AddLast("varIntLength", packet.NewPacketCodecVarIntLength())
	this.pipeline.AddLast("registry", NewCodecRegistry(this))
	this.connCodec = packet.NewPacketConnCodec(this.conn, this.pipeline, 10*time.Second)
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

func (this *ConnectImpl) Connected() (val bool) {
	val = this.conn != nil
	return
}

func (this *ConnectImpl) Write(packet packet.Packet) (err error) {
	err = this.connCodec.Write(packet)
	return
}

func (this *ConnectImpl) HandlePacket(packet packet.Packet) (err error) {
	switch packet.Id() {
	case connect.PACKET_KEEPALIVE:
		err = this.Write(packet)
	case connect.PACKET_RESULT:
		packetResult := packet.(*connect.PacketResult)
		err = this.DispatchResult(packetResult.SequenceId, packetResult.StatusCode, packetResult.Result)
	case connect.PACKET_MESSAGE_EVENT:
		this.DispatchEvent("message", WrapEventMessage(packet.(*connect.PacketMessageEvent)))
	case connect.PACKET_REDIRECT_EVENT:
		this.DispatchEvent("redirect", WrapEventRedirect(packet.(*connect.PacketRedirectEvent)))
	case connect.PACKET_SERVER_EVENT:
		this.DispatchEvent("server", WrapEventServer(packet.(*connect.PacketServerEvent)))
	case connect.PACKET_PLAYER_EVENT:
		this.DispatchEvent("player", WrapEventPlayer(packet.(*connect.PacketPlayerEvent)))
	}
	return
}

func (this *ConnectImpl) ErrorCaught(err error) {
	fmt.Println("Connect client, disconnected:", err)
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
	statusCode = <-statusCodeChannel
	result = <-resultChannel
	return
}

func (this *ConnectImpl) RequestLater(request connect.Request, callback RequestCallback) (err error) {
	if !this.Connected() {
		err = errors.New("Not connected")
		return
	}
	sequenceId := atomic.AddInt32(&this.sequenceId, 1)
	this.recordsMutex.Lock()
	if this.records != nil {
		this.records[sequenceId] = NewRequestRecord(request, callback)
	}
	this.recordsMutex.Unlock()
	err = this.Write(connect.NewPacketRequest(sequenceId, request))
	if err != nil {
		this.recordsMutex.Lock()
		delete(this.records, sequenceId)
		this.recordsMutex.Unlock()
		return
	}
	return
}

func (this *ConnectImpl) DispatchResult(sequenceId int32, statusCode uint8, result connect.Result) (err error) {
	this.recordsMutex.Lock()
	var record *RequestRecord
	var ok bool
	if record, ok = this.records[sequenceId]; !ok {
		err = errors.New("No matching request for result")
		return
	}
	this.recordsMutex.Unlock()
	if record.callback != nil {
		go record.callback(statusCode, result)
	}
	this.recordsMutex.Lock()
	delete(this.records, sequenceId)
	this.recordsMutex.Unlock()
	return
}

func (this *ConnectImpl) RequestIdBySequenceId(sequenceId int32) (requestId int) {
	this.recordsMutex.Lock()
	defer this.recordsMutex.Unlock()
	if record, ok := this.records[sequenceId]; !ok {
		requestId = -1
	} else {
		requestId = record.request.Id()
	}
	return
}
