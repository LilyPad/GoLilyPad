package connect

import "errors"
import "net"
import "time"
import "github.com/LilyPad/GoLilyPad/packet"
import "github.com/LilyPad/GoLilyPad/packet/connect"

type Session struct {
	server *Server
	conn net.Conn
	connCodec *packet.PacketConnCodec

	uuid string
	salt string
	username string
	keepalive *int32
	role SessionRole

	roleAddress string
	rolePort uint16
	serverSecurityKey string
	proxyMotd string
	proxyVersion string
	proxyPlayers map[string]bool
	proxyMaxPlayers uint16
}

func NewSession(server *Server, conn net.Conn) (session *Session, err error) {
	uuid, err := GenUUID()
	if err != nil {
		return
	}
	salt, err := GenSalt()
	if err != nil {
		return
	}
	session = &Session{
		server: server,
		conn: conn,
		uuid: uuid,
		salt: salt,
		role: UNAUTHORIZED,
	}
	return
}

func (this *Session) Serve() {
	this.connCodec = packet.NewPacketConnCodec(this.conn, connect.PacketCodec, 10 * time.Second)
	go this.connCodec.ReadConn(this)
}

func (this *Session) Write(packet packet.Packet) (err error) {
	return this.connCodec.Write(packet)
}

func (this *Session) Close() {
	this.conn.Close()
}

func (this *Session) Keepalive() (err error) {
	if this.keepalive != nil {
		return
	}
	keepalive := RandomInt()
	err = this.Write(&connect.PacketKeepalive{keepalive})
	if err != nil {
		return
	}
	this.keepalive = &keepalive
	return
}

func (this *Session) RegisterAuthorized(username string) {
	this.username = username
	this.role = AUTHORIZED
	this.server.SessionRegistry(AUTHORIZED).Register(this)
}

func (this *Session) RegisterProxy(address string, port uint16, motd string, version string, maxPlayers uint16) bool {
	sessionRegistry := this.server.SessionRegistry(AUTHORIZED)
	if sessionRegistry.HasId(this.username) {
		return false
	}
	var err error
	if len(address) == 0 {
		address, err = this.RemoteIp()
		if err != nil {
			return false
		}
	}
	this.Unregister()
	this.role = ROLE_PROXY
	this.roleAddress = address
	this.rolePort = port
	this.proxyMotd = motd
	this.proxyVersion = version
	this.proxyPlayers = make(map[string]bool)
	this.proxyMaxPlayers = maxPlayers
	this.server.SessionRegistry(AUTHORIZED).Register(this)
	this.server.SessionRegistry(ROLE_PROXY).Register(this)
	this.server.NetworkCache().RegisterProxy(this)
	for _, session := range this.server.SessionRegistry(ROLE_SERVER).GetAll() {
		this.Write(&connect.PacketServerEvent{true, session.username, session.serverSecurityKey, session.roleAddress, session.rolePort})
	}
	return true
}

func (this *Session) RegisterServer(address string, port uint16) bool {
	sessionRegistry := this.server.SessionRegistry(AUTHORIZED)
	if sessionRegistry.HasId(this.username) {
		return false
	}
	var err error
	if len(address) == 0 {
		address, err = this.RemoteIp()
		if err != nil {
			return false
		}
	}
	securityKey, err := GenSalt()
	if err != nil {
		return false
	}
	this.Unregister()
	this.role = ROLE_SERVER
	this.roleAddress = address
	this.rolePort = port
	this.serverSecurityKey = securityKey
	this.server.SessionRegistry(AUTHORIZED).Register(this)
	this.server.SessionRegistry(ROLE_SERVER).Register(this)
	for _, session := range this.server.SessionRegistry(ROLE_PROXY).GetAll() {
		session.Write(&connect.PacketServerEvent{true, this.username, this.serverSecurityKey, this.roleAddress, this.rolePort})
	}
	return true
}

func (this *Session) Unregister() {
	if this.role == AUTHORIZED {
		this.server.SessionRegistry(AUTHORIZED).Unregister(this)
	} else if this.role == ROLE_PROXY {
		this.server.SessionRegistry(AUTHORIZED).Unregister(this)
		this.server.SessionRegistry(ROLE_PROXY).Unregister(this)
		this.server.NetworkCache().UnregisterProxy(this)
	} else if this.role == ROLE_SERVER {
		for _, session := range this.server.SessionRegistry(ROLE_PROXY).GetAll() {
			session.Write(&connect.PacketServerEvent{
				Add: false, 
				Server: this.username,
			})
		}
		this.server.SessionRegistry(AUTHORIZED).Unregister(this)
		this.server.SessionRegistry(ROLE_SERVER).Unregister(this)
	}
	this.role = UNAUTHORIZED
}

func (this *Session) HandlePacket(packet packet.Packet) (err error) {
	switch packet.Id() {
	case connect.PACKET_KEEPALIVE:
		if this.keepalive == nil {
			this.Close()
			return
		}
		if *this.keepalive != packet.(*connect.PacketKeepalive).Random {
			this.Close()
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
				ok, err = this.server.Authenticator().Authenticate(username, request.(*connect.RequestAuthenticate).Password, this.salt)
				if err != nil {
					return
				}
				if ok {
					this.RegisterAuthorized(username)
					result = &connect.ResultAuthenticate{}
					statusCode = connect.STATUS_SUCCESS
				} else {
					statusCode = connect.STATUS_ERROR_GENERIC
				}
			} else {
				statusCode = connect.STATUS_ERROR_ROLE
			}
		case connect.REQUEST_AS_PROXY:
			if this.Authorized() {
				if this.RegisterProxy(request.(*connect.RequestAsProxy).Address, request.(*connect.RequestAsProxy).Port, request.(*connect.RequestAsProxy).Motd, request.(*connect.RequestAsProxy).Version, request.(*connect.RequestAsProxy).Maxplayers) {
					result = &connect.ResultAsProxy{}
					statusCode = connect.STATUS_SUCCESS
				} else {
					statusCode = connect.STATUS_ERROR_GENERIC
				}
			} else {
				statusCode = connect.STATUS_ERROR_ROLE
			}
		case connect.REQUEST_AS_SERVER:
			if this.Authorized() {
				if this.RegisterServer(request.(*connect.RequestAsServer).Address, request.(*connect.RequestAsServer).Port) {
					result = &connect.ResultAsServer{this.serverSecurityKey}
					statusCode = connect.STATUS_SUCCESS
				} else {
					statusCode = connect.STATUS_ERROR_GENERIC
				}
			} else {
				statusCode = connect.STATUS_ERROR_ROLE
			}
		case connect.REQUEST_GET_DETAILS:
			if this.Authorized() {
				networkCache := this.server.NetworkCache()
				result = &connect.ResultGetDetails{networkCache.Address(), networkCache.Port(), networkCache.Motd(), networkCache.Version()}
				statusCode = connect.STATUS_SUCCESS
			} else {
				statusCode = connect.STATUS_ERROR_ROLE
			}
		case connect.REQUEST_GET_PLAYERS:
			if this.Authorized() {
				networkCache := this.server.NetworkCache()
				players := networkCache.Players()
				if request.(*connect.RequestGetPlayers).List {
					result = &connect.ResultGetPlayers{true, uint16(len(players)), networkCache.MaxPlayers(), players}
				} else {
					result = &connect.ResultGetPlayers{false, uint16(len(players)), networkCache.MaxPlayers(), nil}
				}
				statusCode = connect.STATUS_SUCCESS
			} else {
				statusCode = connect.STATUS_ERROR_ROLE
			}
		case connect.REQUEST_GET_SALT:
			result = &connect.ResultGetSalt{this.salt}
			statusCode = connect.STATUS_SUCCESS
		case connect.REQUEST_GET_WHOAMI:
			result = &connect.ResultGetWhoami{this.Id()}
			statusCode = connect.STATUS_SUCCESS
		case connect.REQUEST_MESSAGE:
			if this.Authorized() {
				sessionRegistry := this.server.SessionRegistry(AUTHORIZED)
				recipients := request.(*connect.RequestMessage).Recipients
				messagePacket := &connect.PacketMessageEvent{this.Id(), request.(*connect.RequestMessage).Channel, request.(*connect.RequestMessage).Message}
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
					result = &connect.ResultMessage{}
					statusCode = connect.STATUS_SUCCESS
				} else {
					statusCode = connect.STATUS_ERROR_GENERIC
				}
			} else {
				statusCode = connect.STATUS_ERROR_ROLE
			}
		case connect.REQUEST_NOTIFY_PLAYER:
			if this.Authorized() && this.role == ROLE_PROXY {
				add := request.(*connect.RequestNotifyPlayer).Add
				player := request.(*connect.RequestNotifyPlayer).Player
				if add {
					if !this.server.NetworkCache().AddPlayer(player, this) {
						statusCode = connect.STATUS_ERROR_GENERIC
						break
					}
					this.proxyPlayers[player] = true
				} else {
					if _, ok := this.proxyPlayers[player]; ok {
						this.server.NetworkCache().RemovePlayer(player)
						delete(this.proxyPlayers, player)
					}
				}
				result = &connect.ResultNotifyPlayer{}
				statusCode = connect.STATUS_SUCCESS
			} else {
				statusCode = connect.STATUS_ERROR_ROLE
			}
		case connect.REQUEST_REDIRECT:
			if this.Authorized() {
				server := request.(*connect.RequestRedirect).Server
				player := request.(*connect.RequestRedirect).Player
				session := this.server.NetworkCache().ProxyByPlayer(player)
				if session == nil {
					statusCode = connect.STATUS_ERROR_GENERIC
					break
				}
				session.Write(&connect.PacketRedirectEvent{server, player})
				result = &connect.ResultRedirect{}
				statusCode = connect.STATUS_SUCCESS
			} else {
				statusCode = connect.STATUS_ERROR_ROLE
			}
		default:
			err = errors.New("Request Id is not handled by server")
			return
		}
		err = this.Write(&connect.PacketResult{packet.(*connect.PacketRequest).SequenceId, statusCode, result})
	default:
		err = errors.New("Packet Id is not handled by server")
	}
	return
}

func (this *Session) ErrorCaught(err error) {
	this.conn.Close()
	this.Unregister()
	return
}

func (this *Session) Authorized() bool {
	return this.role != UNAUTHORIZED
}

func (this *Session) RemoteAddr() (addr net.Addr) {
	return this.conn.RemoteAddr()
}

func (this *Session) RemoteIp() (ip string, err error) {
	ip, _, err = net.SplitHostPort(this.RemoteAddr().String())
	return
}

func (this *Session) Id() string {
	if this.role == UNAUTHORIZED {
		return ""
	}
	if this.role == AUTHORIZED {
		return this.username + "." + this.uuid
	}
	return this.username
}