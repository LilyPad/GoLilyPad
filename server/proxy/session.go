package proxy

import "bytes"
import "crypto/aes"
import "crypto/cipher"
import "crypto/rand"
import "crypto/rsa"
import "crypto/x509"
import "encoding/base64"
import "encoding/json"
import "errors"
import "fmt"
import "io/ioutil"
import "net"
import "sync"
import "time"
import "github.com/LilyPad/GoLilyPad/packet"
import "github.com/LilyPad/GoLilyPad/packet/minecraft"
import "github.com/LilyPad/GoLilyPad/server/proxy/connect"
import "github.com/LilyPad/GoLilyPad/server/proxy/auth"

type Session struct {
	server *Server
	conn net.Conn
	connCodec *packet.PacketConnCodec
	codec *packet.PacketCodecVariable
	outBridge *SessionOutBridge
	outMutex *sync.Mutex
	active bool
	redirecting bool

	serverAddress string
	name string
	uuid string
	serverId string
	publicKey []byte
	verifyToken []byte

	clientSettings packet.Packet
	clientEntityId int32
	serverEntityId int32
	playerList map[string]bool
	scoreboards map[string]bool
	teams map[string]bool

	remoteHost string
	remotePort string
	state SessionState
}

func NewSession(server *Server, conn net.Conn) *Session {
	host, port, _ := net.SplitHostPort(conn.RemoteAddr().String())
	return &Session{
		server: server,
		conn: conn,
		outMutex: &sync.Mutex{},
		active: true,
		redirecting: false,
		playerList: make(map[string]bool),
		scoreboards: make(map[string]bool),
		teams: make(map[string]bool),
		remoteHost: host,
		remotePort: port,
		state: STATE_DISCONNECTED,
	}
}

func (this *Session) Serve() {
	this.codec = packet.NewPacketCodecVariable(minecraft.HandshakePacketClientCodec, minecraft.HandshakePacketServerCodec)
	this.connCodec = packet.NewPacketConnCodec(this.conn, this.codec, 10 * time.Second)
	go this.connCodec.ReadConn(this)
}

func (this *Session) Write(packet packet.Packet) (err error) {
	return this.connCodec.Write(packet)
}

func (this *Session) Redirect(server *connect.Server) {
	conn, err := net.Dial("tcp", server.Addr)
	if err != nil {
		fmt.Println("Proxy server, name:", this.name, "ip:", this.RemoteIp(), "failed to redirect:", server.Name, "err:", err)
		if this.Initializing() {
			this.Disconnect("Error: Outbound Connection Mismatch")
		}
		return
	}
	fmt.Println("Proxy server, name:", this.name, "ip:", this.RemoteIp(), "redirected:", server.Name)
	NewSessionOutBridge(this, server, conn).Serve()
}

func (this *Session) SetAuthenticated(result bool) {
	if !result {
		this.Disconnect("Error: Authentication to Minecraft.net Failed")
		return
	}
	if this.server.SessionRegistry().HasName(this.name) {
		this.Disconnect(minecraft.Colorize(this.server.Localizer().LocaleLoggedIn()))
		return
	}
	if this.server.MaxPlayers() > 1 && this.server.SessionRegistry().Len() >= int(this.server.MaxPlayers()) {
		this.Disconnect(minecraft.Colorize(this.server.Localizer().LocaleFull()))
		return
	}
	servers := this.server.Router().Route(this.serverAddress)
	activeServers := []string{}
	for _, serverName := range servers {
		if !this.server.Connect().HasServer(serverName) {
			continue
		}
		activeServers = append(activeServers, serverName)
	}
	if len(availableServers) == 0 {
		this.Disconnect(minecraft.Colorize(this.server.Localizer().LocaleOffline()))
		return
	}
	serverName := activeServers[RandomInt(len(activeServers))]
	server := this.server.Connect().Server(serverName)
	if server == nil {
		this.Disconnect("Error: Outbound Server Mismatch: " + serverName)
		return
	}
	addResult := this.server.Connect().AddLocalPlayer(this.name)
	if addResult == 0 {
		this.Disconnect(minecraft.Colorize(this.server.Localizer().LocaleLoggedIn()))
		return
	} else if addResult == -1 {
		this.Disconnect(minecraft.Colorize(this.server.Localizer().LocaleLoggedIn()))
		return
	}
	this.state = STATE_INIT
	this.Write(&minecraft.PacketClientLoginSuccess{this.uuid, this.name})
	this.codec.SetEncodeCodec(minecraft.PlayPacketClientCodec)
	this.codec.SetDecodeCodec(minecraft.PlayPacketServerCodec)
	this.server.SessionRegistry().Register(this)
	this.Redirect(server)
}

func (this *Session) Disconnect(reason string) {
	reasonJson, _ := json.Marshal(reason);
	this.DisconnectRaw("{\"text\":" + string(reasonJson) + "}")
}

func (this *Session) DisconnectRaw(raw string) {
	if this.codec.EncodeCodec() == minecraft.LoginPacketClientCodec {
		this.Write(&minecraft.PacketClientLoginDisconnect{raw})
	} else if this.codec.EncodeCodec() == minecraft.PlayPacketClientCodec {
		this.Write(&minecraft.PacketClientDisconnect{raw})
	}
	this.conn.Close()
}

func (this *Session) HandlePacket(packet packet.Packet) (err error) {
	switch this.state {
	case STATE_DISCONNECTED:
		if packet.Id() == minecraft.PACKET_SERVER_HANDSHAKE {
			handshakePacket := packet.(*minecraft.PacketServerHandshake)
			this.serverAddress = handshakePacket.ServerAddress
			if handshakePacket.State == 1 {
				this.codec.SetEncodeCodec(minecraft.StatusPacketClientCodec)
				this.codec.SetDecodeCodec(minecraft.StatusPacketServerCodec)
				this.state = STATE_STATUS
			} else if handshakePacket.State ==  2 {
				if handshakePacket.ProtocolVersion != minecraft.VERSION {
					err = errors.New("Protocol version does not match")
					return
				}
				this.codec.SetEncodeCodec(minecraft.LoginPacketClientCodec)
				this.codec.SetDecodeCodec(minecraft.LoginPacketServerCodec)
				this.state = STATE_LOGIN
			} else {
				err = errors.New("Unexpected state")
				return
			}
		} else {
			err = errors.New("Unexpected packet")
			return
		}
	case STATE_STATUS:
		if packet.Id() == minecraft.PACKET_SERVER_STATUS_REQUEST {
			favicon, faviconErr := ioutil.ReadFile("server-icon.png")
			var faviconString string
			if faviconErr == nil {
				faviconString = "data:image/png;base64," + base64.StdEncoding.EncodeToString(favicon)
			}
			version := make(map[string]interface{})
			version["name"] = minecraft.STRING_VERSION
			version["protocol"] = minecraft.VERSION
			players := make(map[string]interface{})
			players["max"] = this.server.Connect().MaxPlayers()
			players["online"] = this.server.Connect().Players()
			description := make(map[string]interface{})
			description["text"] = minecraft.Colorize(this.server.Router().RouteMotd(this.serverAddress))
			response := make(map[string]interface{})
			response["version"] = version
			response["players"] = players
			response["description"] = description
			if faviconString != "" {
				response["favicon"] = faviconString
			}
			var marshalled []byte
			marshalled, err = json.Marshal(response)
			if err != nil {
				return
			}
			err = this.Write(&minecraft.PacketClientStatusResponse{string(marshalled)})
			if err != nil {
				return
			}
			this.state = STATE_STATUS_PING
		} else {
			err = errors.New("Unexpected packet")
			return
		}
	case STATE_STATUS_PING:
		if packet.Id() == minecraft.PACKET_SERVER_STATUS_PING {
			err = this.Write(&minecraft.PacketClientStatusPing{packet.(*minecraft.PacketServerStatusPing).Time})
			if err != nil {
				return
			}
			this.conn.Close()
		} else {
			err = errors.New("Unexpected packet")
			return
		}
	case STATE_LOGIN:
		if packet.Id() == minecraft.PACKET_SERVER_LOGIN_START {
			this.name = packet.(*minecraft.PacketServerLoginStart).Name
			if this.server.Authenticate() {
				this.serverId, err = GenSalt()
				if err != nil {
					return
				}
				this.publicKey, err = x509.MarshalPKIXPublicKey(&this.server.PrivateKey().PublicKey)
				if err != nil {
					return
				}
				this.verifyToken, err = RandomBytes(4)
				if err != nil {
					return
				}
				err = this.Write(&minecraft.PacketClientLoginEncryptRequest{this.serverId, this.publicKey, this.verifyToken})
				if err != nil {
					return
				}
				this.state = STATE_LOGIN_ENCRYPT
			} else {
				this.uuid = GenNameUUID("OfflinePlayer:" + this.name)
				this.SetAuthenticated(true)
			}
		} else {
			err = errors.New("Unexpected packet")
			return
		}
	case STATE_LOGIN_ENCRYPT:
		if packet.Id() == minecraft.PACKET_SERVER_LOGIN_ENCRYPT_RESPONSE {
			loginEncryptResponsePacket := packet.(*minecraft.PacketServerLoginEncryptResponse)
			var sharedSecret []byte
			sharedSecret, err = rsa.DecryptPKCS1v15(rand.Reader, this.server.PrivateKey(), loginEncryptResponsePacket.SharedSecret)
			if err != nil {
				return
			}
			var verifyToken []byte
			verifyToken, err = rsa.DecryptPKCS1v15(rand.Reader, this.server.PrivateKey(), loginEncryptResponsePacket.VerifyToken)
			if err != nil {
				return
			}
			if bytes.Compare(this.verifyToken, verifyToken) != 0 {
				err = errors.New("Verify token does not match")
				return
			}
			var block cipher.Block
			block, err = aes.NewCipher(sharedSecret)
			if err != nil {
				return
			}
			this.connCodec.SetReader(&cipher.StreamReader{
				R: this.connCodec.Reader(),
				S: minecraft.NewCFB8Decrypter(block, sharedSecret),
			})
			this.connCodec.SetWriter(&cipher.StreamWriter{
				W: this.connCodec.Writer(),
				S: minecraft.NewCFB8Encrypter(block, sharedSecret),
			})
			var authErr error
			this.uuid, authErr = auth.Authenticate(this.name, this.serverId, sharedSecret, this.publicKey)
			if authErr != nil {
				this.SetAuthenticated(false)
				fmt.Println("Proxy server, failed to authorize:", this.name, "ip:", this.RemoteIp(), "err:", authErr)
			} else {
				this.SetAuthenticated(true)
				fmt.Println("Proxy server, authorized:", this.name, "ip:", this.RemoteIp())
			}
		} else {
			err = errors.New("Unexpected packet")
			return
		}
	case STATE_CONNECTED:
		if packet.Id() == minecraft.PACKET_SERVER_CLIENT_SETTINGS {
			this.clientSettings = packet
		}
		if this.redirecting {
			break
		}
		if genericPacket, ok := packet.(*minecraft.PacketGeneric); ok {
			genericPacket.SwapEntities(this.clientEntityId, this.serverEntityId, false)
		}
		this.outBridge.Write(packet)
	}
	return
}

func (this *Session) ErrorCaught(err error) {
	if this.Authenticated() {
		this.server.Connect().RemoveLocalPlayer(this.name)
		this.server.SessionRegistry().Unregister(this)
		fmt.Println("Proxy server, name:", this.name, "ip:", this.RemoteIp(), "disconnected:", err)
	}
	this.state = STATE_DISCONNECTED
	this.conn.Close()
	return
}

func (this *Session) Name() string {
	return this.name
}

func (this *Session) RemoteAddr() (addr net.Addr) {
	return this.conn.RemoteAddr()
}

func (this *Session) RemoteIp() (ip string) {
	ip, _, _ = net.SplitHostPort(this.RemoteAddr().String())
	return
}

func (this *Session) Authenticated() bool {
	return this.state == STATE_INIT || this.state == STATE_CONNECTED
}

func (this *Session) Initializing() bool {
	return this.state == STATE_INIT
}
