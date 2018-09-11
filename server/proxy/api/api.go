package api

import (
	"github.com/LilyPad/GoLilyPad/packet"
	"github.com/LilyPad/GoLilyPad/packet/minecraft"
	uuid "github.com/satori/go.uuid"
	"net"
)

type Plugin interface {
	Init(context Context)
}

type Context interface {
	Config() Config
	EventBus() EventBus
	SessionRegistry() SessionRegistry
}

type Config interface {
	ListenAddr() string
	Motd() string
	MaxPlayers() uint16
	SyncMaxPlayers() bool
	Authenticate() bool
}

type Session interface {
	Conn() net.Conn
	Write(packet.Packet, PacketSubject)
	Profile() (name string, uuid uuid.UUID)
	Disconnect(reason string)
	DisconnectJson(json string)
	Remote() (ip string, port string)
	RemoteOverride(ip string, port string)
	State() SessionState
	Version() *minecraft.Version
}

type SessionRegistry interface {
	HasName(name string) bool
	HasUuid(uuid uuid.UUID) bool
	GetByName(name string) (session Session)
	GetByUuid(uuid uuid.UUID) (session Session)
	GetAll() (sessions []Session)
	Len() int
}

type SessionState int

var SessionStateAll []SessionState = []SessionState{
	SessionStateDisconnected,
	SessionStateStatus,
	SessionStateStatusPing,
	SessionStateLogin,
	SessionStateLoginEncrypt,
	SessionStateInit,
	SessionStateConnected,
}

const (
	SessionStateDisconnected SessionState = iota
	SessionStateStatus
	SessionStateStatusPing
	SessionStateLogin
	SessionStateLoginEncrypt
	SessionStateInit
	SessionStateConnected
	SessionStateMax
)

type PacketStage int

const (
	PacketStagePre PacketStage = iota
	PacketStageMonitor
	PacketStageMax
)

type PacketSubject int

const (
	PacketSubjectClient PacketSubject = iota
	PacketSubjectOutBridge
	PacketSubjectMax
)

type PacketDirection int

const (
	PacketDirectionRead PacketDirection = iota
	PacketDirectionWrite
	PacketDirectionMax
)

type EventBus interface {
	HandleSessionOpen(EventSessionHandler)
	HandleSessionClose(EventSessionHandler)
	HandleSessionState(EventSessionHandler)
	HandleSessionRedirect(EventSessionHandler)
	HandleSessionPacket(EventSessionHandler, PacketStage, PacketSubject, PacketDirection, ...SessionState)
}

type EventCancellable interface {
	SetCancelled(cancelled bool)
	IsCancelled() bool
}

type EventSession interface {
	Session() Session
}

type EventSessionHandler func(EventSession)

type EventSessionOpen interface {
	EventSession
	EventCancellable
}

type EventSessionClose interface {
	EventSession
}

type EventSessionState interface {
	EventSession
	State() SessionState
}

type EventSessionRedirect interface {
	EventSession
	EventCancellable
	Init() bool
	ServerName() string
	ServerAddr() string
}

type EventSessionPacket interface {
	EventSession
	EventCancellable
	Packet() packet.Packet
	PacketSubject() PacketSubject
	PacketDirection() PacketDirection
	SetPacket(packet.Packet)
}
