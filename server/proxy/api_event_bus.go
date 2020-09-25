package proxy

import (
	"github.com/LilyPad/GoLilyPad/packet"
	"github.com/LilyPad/GoLilyPad/server/proxy/api"
)

type eventBus struct {
	sessionOpen        []api.EventSessionHandler
	sessionLogin       []api.EventSessionHandler
	sessionClose       []api.EventSessionHandler
	sessionState       []api.EventSessionHandler
	sessionRedirect    []api.EventSessionHandler
	sessionPacket      [][][][][]api.EventSessionHandler
	emptySessionPacket *eventSessionPacket
}

func NewEventBus() *eventBus {
	this := new(eventBus)
	this.sessionOpen = make([]api.EventSessionHandler, 0)
	this.sessionLogin = make([]api.EventSessionHandler, 0)
	this.sessionClose = make([]api.EventSessionHandler, 0)
	this.sessionState = make([]api.EventSessionHandler, 0)
	this.sessionRedirect = make([]api.EventSessionHandler, 0)
	this.sessionPacket = make([][][][][]api.EventSessionHandler, api.PacketStageMax)
	for i := range this.sessionPacket {
		this.sessionPacket[i] = make([][][][]api.EventSessionHandler, api.PacketSubjectMax)
		for j := range this.sessionPacket[i] {
			this.sessionPacket[i][j] = make([][][]api.EventSessionHandler, api.PacketDirectionMax)
			for k := range this.sessionPacket[i][j] {
				this.sessionPacket[i][j][k] = make([][]api.EventSessionHandler, api.SessionStateMax)
				for l := range this.sessionPacket[i][j][k] {
					this.sessionPacket[i][j][k][l] = make([]api.EventSessionHandler, 0)
				}
			}
		}
	}
	this.emptySessionPacket = new(eventSessionPacket)
	return this
}

func (this *eventBus) fireEventSession(handlers []api.EventSessionHandler, event api.EventSession) {
	for _, handler := range handlers {
		handler(event)
	}
}

func (this *eventBus) fireEventSessionPacket(session *Session, packet *packet.Packet, stage api.PacketStage, subject api.PacketSubject, direction api.PacketDirection) (event *eventSessionPacket) {
	handlers := this.sessionPacket[stage][subject][direction][session.state]
	if len(handlers) == 0 {
		event = this.emptySessionPacket
		return
	}
	event = &eventSessionPacket{
		eventSessionCancellable: eventSessionCancellable{eventSession: eventSession{session}},
		packet:                  *packet,
		packetStage:             stage,
		packetSubject:           subject,
		packetDirection:         direction,
	}
	this.fireEventSession(handlers, event)
	*packet = event.packet
	return
}

func (this *eventBus) HandleSessionOpen(handler api.EventSessionHandler) {
	this.sessionOpen = append(this.sessionOpen, handler)
}

func (this *eventBus) HandleSessionLogin(handler api.EventSessionHandler) {
	this.sessionLogin = append(this.sessionLogin, handler)
}

func (this *eventBus) HandleSessionClose(handler api.EventSessionHandler) {
	this.sessionClose = append(this.sessionClose, handler)
}

func (this *eventBus) HandleSessionState(handler api.EventSessionHandler) {
	this.sessionState = append(this.sessionState, handler)
}

func (this *eventBus) HandleSessionRedirect(handler api.EventSessionHandler) {
	this.sessionRedirect = append(this.sessionRedirect, handler)
}

func (this *eventBus) HandleSessionPacket(handler api.EventSessionHandler, stage api.PacketStage, subject api.PacketSubject, direction api.PacketDirection, states ...api.SessionState) {
	if len(states) == 0 {
		states = api.SessionStateAll
	}
	for _, state := range states {
		this.sessionPacket[stage][subject][direction][state] = append(this.sessionPacket[stage][subject][direction][state], handler)
	}
}

type eventSession struct {
	session *Session
}

func (this *eventSession) Session() api.Session {
	return this.session.apiSession
}

type eventSessionCancellable struct {
	eventSession
	cancelled bool
}

func (this *eventSessionCancellable) SetCancelled(cancelled bool) {
	this.cancelled = cancelled
}

func (this *eventSessionCancellable) IsCancelled() bool {
	return this.cancelled
}

type eventSessionOpen struct {
	eventSessionCancellable
}

type eventSessionLogin struct {
	eventSessionCancellable
	reason string
}

func (this *eventSessionLogin) SetReason(reason string) {
	this.reason = reason
}

func (this *eventSessionLogin) GetReason() string {
	return this.reason
}

type eventSessionClose struct {
	eventSession
}

type eventSessionState struct {
	eventSession
	state api.SessionState
}

func (this *eventSessionState) State() api.SessionState {
	return this.state
}

type eventSessionRedirect struct {
	eventSessionCancellable
	init       bool
	serverName string
	serverAddr string
}

func (this *eventSessionRedirect) Init() bool {
	return this.init
}

func (this *eventSessionRedirect) ServerName() string {
	return this.serverName
}

func (this *eventSessionRedirect) ServerAddr() string {
	return this.serverAddr
}

type eventSessionPacket struct {
	eventSessionCancellable
	packet          packet.Packet
	packetStage     api.PacketStage
	packetSubject   api.PacketSubject
	packetDirection api.PacketDirection
}

func (this *eventSessionPacket) Packet() packet.Packet {
	return this.packet
}

func (this *eventSessionPacket) PacketSubject() api.PacketSubject {
	return this.packetSubject
}

func (this *eventSessionPacket) PacketDirection() api.PacketDirection {
	return this.packetDirection
}

func (this *eventSessionPacket) SetPacket(packet packet.Packet) {
	this.packet = packet
}
