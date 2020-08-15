package connect

import (
	"errors"
	"fmt"
	"github.com/LilyPad/GoLilyPad/packet"
	"github.com/LilyPad/GoLilyPad/packet/connect"
	uuid "github.com/satori/go.uuid"
	"net"
	"time"
)

type Session struct {
	server    *Server
	conn      net.Conn
	connCodec *packet.PacketConnCodec
	pipeline  *packet.PacketPipeline

	uuid      string
	salt      string
	username  string
	keepalive *int32
	role      SessionRole

	roleAddress       string
	rolePort          uint16
	serverSecurityKey string
	proxyMotd         string
	proxyVersion      string
	proxyPlayers      map[string]uuid.UUID
	proxyMaxPlayers   uint16

	remoteIp   string
	remotePort string
}

func NewSession(server *Server, conn net.Conn) (this *Session) {
	this = new(Session)
	this.server = server
	this.conn = conn
	this.role = ROLE_UNAUTHORIZED
	this.uuid, _ = GenUUID()
	this.salt, _ = GenSalt()
	this.remoteIp, this.remotePort, _ = net.SplitHostPort(conn.RemoteAddr().String())
	return
}

func (this *Session) Serve() {
	this.pipeline = packet.NewPacketPipeline()
	this.pipeline.AddLast("varIntLength", packet.NewPacketCodecVarIntLength())
	this.pipeline.AddLast("registry", connect.PacketCodec)
	this.connCodec = packet.NewPacketConnCodec(this.conn, this.pipeline, 10*time.Second)
	this.connCodec.ReadConn(this)
}

func (this *Session) Write(packet packet.Packet) (err error) {
	err = this.connCodec.Write(packet)
	return
}

func (this *Session) Keepalive() (err error) {
	if this.keepalive != nil {
		return
	}
	keepalive := RandomInt()
	err = this.Write(connect.NewPacketKeepalive(keepalive))
	if err != nil {
		return
	}
	this.keepalive = &keepalive
	return
}

func (this *Session) RegisterAuthorized(username string) {
	this.username = username
	this.role = ROLE_AUTHORIZED
	this.server.SessionRegistry(ROLE_AUTHORIZED).Register(this)
}

func (this *Session) RegisterProxy(address string, port uint16, motd string, version string, maxPlayers uint16) (ok bool) {
	sessionRegistry := this.server.SessionRegistry(ROLE_AUTHORIZED)
	if sessionRegistry.HasId(this.username) {
		ok = false
		return
	}
	if len(address) == 0 {
		address = this.remoteIp
	}
	this.Unregister()
	this.role = ROLE_PROXY
	this.roleAddress = address
	this.rolePort = port
	this.proxyMotd = motd
	this.proxyVersion = version
	this.proxyPlayers = make(map[string]uuid.UUID)
	this.proxyMaxPlayers = maxPlayers
	this.server.SessionRegistry(ROLE_AUTHORIZED).Register(this)
	this.server.SessionRegistry(ROLE_PROXY).Register(this)
	this.server.networkCache.RegisterProxy(this)
	for _, session := range this.server.SessionRegistry(ROLE_SERVER).GetAll() {
		this.Write(connect.NewPacketServerEventAdd(session.username, session.serverSecurityKey, session.roleAddress, session.rolePort))
	}
	ok = true
	return
}

func (this *Session) RegisterServer(address string, port uint16) (ok bool) {
	sessionRegistry := this.server.SessionRegistry(ROLE_AUTHORIZED)
	if sessionRegistry.HasId(this.username) {
		ok = false
		return
	}
	if len(address) == 0 {
		address = this.remoteIp
	}
	securityKey, err := GenSalt()
	if err != nil {
		ok = false
		return
	}
	this.Unregister()
	this.role = ROLE_SERVER
	this.roleAddress = address
	this.rolePort = port
	this.serverSecurityKey = securityKey
	this.server.SessionRegistry(ROLE_AUTHORIZED).Register(this)
	this.server.SessionRegistry(ROLE_SERVER).Register(this)
	for _, session := range this.server.SessionRegistry(ROLE_PROXY).GetAll() {
		session.Write(connect.NewPacketServerEventAdd(this.username, this.serverSecurityKey, this.roleAddress, this.rolePort))
	}
	ok = true
	return
}

func (this *Session) Unregister() {
	if this.role == ROLE_AUTHORIZED {
		this.server.SessionRegistry(ROLE_AUTHORIZED).Unregister(this)
	} else if this.role == ROLE_PROXY {
		this.server.SessionRegistry(ROLE_AUTHORIZED).Unregister(this)
		this.server.SessionRegistry(ROLE_PROXY).Unregister(this)
		this.server.networkCache.UnregisterProxy(this)
	} else if this.role == ROLE_SERVER {
		for _, session := range this.server.SessionRegistry(ROLE_PROXY).GetAll() {
			session.Write(connect.NewPacketServerEventRemove(this.username))
		}
		this.server.SessionRegistry(ROLE_AUTHORIZED).Unregister(this)
		this.server.SessionRegistry(ROLE_SERVER).Unregister(this)
	}
	this.role = ROLE_UNAUTHORIZED
}

func (this *Session) HandlePacket(packet packet.Packet) (err error) {
	switch packet.Id() {
	case connect.PACKET_KEEPALIVE:
		if this.keepalive == nil {
			this.conn.Close()
			return
		}
		if *this.keepalive != packet.(*connect.PacketKeepalive).Random {
			this.conn.Close()
			return
		}
		this.keepalive = nil
	case connect.PACKET_REQUEST:
		request := packet.(*connect.PacketRequest).Request
		var result connect.Result
		var statusCode uint8
		switch request.Id() {
		case connect.REQUEST_AUTHENTICATE:
			if !this.Authorized() {
				username := request.(*connect.RequestAuthenticate).Username
				var ok bool
				ok, err = this.server.authenticator.Authenticate(username, request.(*connect.RequestAuthenticate).Password, this.salt)
				if err != nil {
					return
				}
				if ok {
					this.RegisterAuthorized(username)
					fmt.Println("Connect server, authorized:", this.Id(), "ip:", this.remoteIp)
					result = connect.NewResultAuthenticate()
					statusCode = connect.STATUS_SUCCESS
				} else {
					fmt.Println("Connect server, failure to authorize:", username, "ip:", this.remoteIp)
					statusCode = connect.STATUS_ERROR_GENERIC
				}
			} else {
				statusCode = connect.STATUS_ERROR_ROLE
			}
		case connect.REQUEST_AS_PROXY:
			if this.Authorized() {
				requestAsProxy := request.(*connect.RequestAsProxy)
				if this.RegisterProxy(requestAsProxy.Address, requestAsProxy.Port, requestAsProxy.Motd, requestAsProxy.Version, requestAsProxy.MaxPlayers) {
					fmt.Println("Connect server, roled as proxy:", this.Id(), "ip:", this.remoteIp)
					result = connect.NewResultAsProxy()
					statusCode = connect.STATUS_SUCCESS
				} else {
					statusCode = connect.STATUS_ERROR_GENERIC
				}
			} else {
				statusCode = connect.STATUS_ERROR_ROLE
			}
		case connect.REQUEST_AS_SERVER:
			if this.Authorized() {
				requestAsServer := request.(*connect.RequestAsServer)
				if this.RegisterServer(requestAsServer.Address, requestAsServer.Port) {
					fmt.Println("Connect server, roled as server:", this.Id(), "ip:", this.remoteIp)
					result = connect.NewResultAsServer(this.serverSecurityKey)
					statusCode = connect.STATUS_SUCCESS
				} else {
					statusCode = connect.STATUS_ERROR_GENERIC
				}
			} else {
				statusCode = connect.STATUS_ERROR_ROLE
			}
		case connect.REQUEST_GET_DETAILS:
			if this.Authorized() {
				result = connect.NewResultGetDetails(this.server.networkCache.Address(), this.server.networkCache.Port(), this.server.networkCache.Motd(), this.server.networkCache.Version())
				statusCode = connect.STATUS_SUCCESS
			} else {
				statusCode = connect.STATUS_ERROR_ROLE
			}
		case connect.REQUEST_GET_PLAYERS:
			if this.Authorized() {
				players := this.server.networkCache.Players()
				if request.(*connect.RequestGetPlayers).List {
					result = connect.NewResultGetPlayersList(uint16(len(players)), this.server.networkCache.MaxPlayers(), players)
				} else {
					result = connect.NewResultGetPlayers(uint16(len(players)), this.server.networkCache.MaxPlayers())
				}
				statusCode = connect.STATUS_SUCCESS
			} else {
				statusCode = connect.STATUS_ERROR_ROLE
			}
		case connect.REQUEST_GET_SALT:
			result = connect.NewResultGetSalt(this.salt)
			statusCode = connect.STATUS_SUCCESS
		case connect.REQUEST_GET_WHOAMI:
			result = connect.NewResultGetWhoami(this.Id())
			statusCode = connect.STATUS_SUCCESS
		case connect.REQUEST_MESSAGE:
			if this.Authorized() {
				requestMessage := request.(*connect.RequestMessage)
				sessionRegistry := this.server.SessionRegistry(ROLE_AUTHORIZED)
				recipients := request.(*connect.RequestMessage).Recipients
				messagePacket := connect.NewPacketMessageEvent(this.Id(), requestMessage.Channel, requestMessage.Message)
				messageSent := false
				if len(recipients) == 0 {
					for _, recipient := range sessionRegistry.GetAll() {
						recipient.Write(messagePacket)
						messageSent = true
					}
				} else {
					for _, recipientId := range recipients {
						recipient := sessionRegistry.GetById(recipientId)
						if recipient == nil {
							continue
						}
						recipient.Write(messagePacket)
						messageSent = true
					}
				}
				if messageSent {
					result = connect.NewResultMessage()
					statusCode = connect.STATUS_SUCCESS
				} else {
					statusCode = connect.STATUS_ERROR_GENERIC
				}
			} else {
				statusCode = connect.STATUS_ERROR_ROLE
			}
		case connect.REQUEST_NOTIFY_PLAYER:
			if this.Authorized() && this.role == ROLE_PROXY {
				var playerPacket *connect.PacketPlayerEvent

				add := request.(*connect.RequestNotifyPlayer).Add
				player := request.(*connect.RequestNotifyPlayer).Player
				uuid := request.(*connect.RequestNotifyPlayer).Uuid
				if add {
					if !this.server.networkCache.AddPlayer(player, uuid, this) {
						statusCode = connect.STATUS_ERROR_GENERIC
						break
					}
					this.proxyPlayers[player] = uuid
					playerPacket = connect.NewPacketPlayerEventJoin(player, uuid)
				} else {
					if uuid, ok := this.proxyPlayers[player]; ok {
						this.server.networkCache.RemovePlayer(player, uuid)
						delete(this.proxyPlayers, player)
						playerPacket = connect.NewPacketPlayerEventLeave(player, uuid)
					}
				}
				for _, session := range this.server.SessionRegistry(ROLE_AUTHORIZED).GetAll() {
					session.Write(playerPacket)
				}
				result = connect.NewResultNotifyPlayer()
			} else {
				statusCode = connect.STATUS_ERROR_ROLE
			}
		case connect.REQUEST_REDIRECT:
			if this.Authorized() {
				requestRedirect := request.(*connect.RequestRedirect)
				session := this.server.networkCache.ProxyByPlayer(requestRedirect.Player)
				if session != nil {
					session.Write(connect.NewPacketRedirectEvent(requestRedirect.Server, requestRedirect.Player))
					result = connect.NewResultRedirect()
					statusCode = connect.STATUS_SUCCESS
				} else {
					statusCode = connect.STATUS_ERROR_GENERIC
				}
			} else {
				statusCode = connect.STATUS_ERROR_ROLE
			}
		default:
			err = errors.New(fmt.Sprintf("Request Id is not handled by server: %d", request.Id()))
			return
		}
		err = this.Write(connect.NewPacketResult(packet.(*connect.PacketRequest).SequenceId, statusCode, result))
	default:
		err = errors.New(fmt.Sprintf("Packet Id is not handled by server: %d", packet.Id()))
	}
	return
}

func (this *Session) ErrorCaught(err error) {
	id := this.Id()
	if len(id) > 0 {
		fmt.Println("Connect server, id:", this.Id(), "ip:", this.remoteIp, "disconnected:", err)
	}
	this.conn.Close()
	this.Unregister()
	return
}

func (this *Session) Authorized() (val bool) {
	val = this.role != ROLE_UNAUTHORIZED
	return
}

func (this *Session) Id() (id string) {
	if this.role == ROLE_UNAUTHORIZED {
		id = ""
	} else if this.role == ROLE_AUTHORIZED {
		id = this.username + "." + this.uuid
	} else {
		id = this.username
	}
	return
}
