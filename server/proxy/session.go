package proxy

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"time"
	"strings"
	"sync"
	"github.com/LilyPad/GoLilyPad/packet"
	"github.com/LilyPad/GoLilyPad/packet/minecraft"
	"github.com/LilyPad/GoLilyPad/server/proxy/connect"
	"github.com/LilyPad/GoLilyPad/server/proxy/auth"
)

type Session struct {
	server *Server
	conn net.Conn
	connCodec *packet.PacketConnCodec
	pipeline *packet.PacketPipeline
	outBridge *SessionOutBridge
	compressionThreshold int
	active bool

	redirectMutex sync.Mutex
	redirecting bool

	protocolVersion int
	protocol17 bool
	serverAddress string
	name string
	profile auth.GameProfile
	serverId string
	publicKey []byte
	verifyToken []byte

	clientSettings packet.Packet
	clientEntityId int32
	serverEntityId int32
	pluginChannels map[string]struct{}
	playerList map[string]struct{}
	scoreboards map[string]struct{}
	teams map[string]struct{}

	remoteIp string
	remotePort string
	state SessionState
}

func NewSession(server *Server, conn net.Conn) (this *Session) {
	this = new(Session)
	this.server = server
	this.conn = conn
	this.compressionThreshold = -1
	this.active = true
	this.redirecting = false
	this.pluginChannels = make(map[string]struct{})
	this.playerList = make(map[string]struct{})
	this.scoreboards = make(map[string]struct{})
	this.teams = make(map[string]struct{})
	this.remoteIp, this.remotePort, _ = net.SplitHostPort(conn.RemoteAddr().String())
	this.state = STATE_DISCONNECTED
	return
}

func (this *Session) Serve() {
	this.pipeline = packet.NewPacketPipeline()
	this.pipeline.AddLast("varIntLength", packet.NewPacketCodecVarIntLength())
	this.pipeline.AddLast("registry", minecraft.HandshakePacketServerCodec)
	this.connCodec = packet.NewPacketConnCodec(this.conn, this.pipeline, 30 * time.Second)
	this.connCodec.ReadConn(this)
}

func (this *Session) Write(packet packet.Packet) (err error) {
	err = this.connCodec.Write(packet)
	return
}

func (this *Session) Redirect(server *connect.Server) {
	conn, err := net.Dial("tcp", server.Addr)
	if err != nil {
		fmt.Println("Proxy server, name:", this.name, "ip:", this.remoteIp, "failed to redirect:", server.Name, "err:", err)
		if this.Initializing() {
			this.Disconnect("Error: Outbound Connection Mismatch")
		}
		return
	}
	fmt.Println("Proxy server, name:", this.name, "ip:", this.remoteIp, "redirected:", server.Name)
	NewSessionOutBridge(this, server, conn).Serve()
}

func (this *Session) SetAuthenticated(result bool) {
	if !result {
		this.Disconnect("Error: Authentication to Minecraft.net Failed")
		return
	}
	if this.server.SessionRegistry.HasName(this.name) {
		this.Disconnect(minecraft.Colorize(this.server.localizer.LocaleLoggedIn()))
		return
	}
	if this.server.MaxPlayers() > 1 && this.server.SessionRegistry.Len() >= int(this.server.MaxPlayers()) {
		this.Disconnect(minecraft.Colorize(this.server.localizer.LocaleFull()))
		return
	}
	servers := this.server.router.Route(this.serverAddress)
	activeServers := []string{}
	for _, serverName := range servers {
		if !this.server.connect.HasServer(serverName) {
			continue
		}
		activeServers = append(activeServers, serverName)
	}
	if len(activeServers) == 0 {
		this.Disconnect(minecraft.Colorize(this.server.localizer.LocaleOffline()))
		return
	}
	serverName := activeServers[RandomInt(len(activeServers))]
	server := this.server.connect.Server(serverName)
	if server == nil {
		this.Disconnect("Error: Outbound Server Mismatch: " + serverName)
		return
	}
	addResult := this.server.connect.AddLocalPlayer(this.name)
	if addResult == 0 {
		this.Disconnect(minecraft.Colorize(this.server.localizer.LocaleLoggedIn()))
		return
	} else if addResult == -1 {
		this.Disconnect(minecraft.Colorize(this.server.localizer.LocaleLoggedIn()))
		return
	}
	this.state = STATE_INIT
	if this.protocolVersion >= 5 {
		this.Write(minecraft.NewPacketClientLoginSuccess(FormatUUID(this.profile.Id), this.name))
	} else {
		this.Write(minecraft.NewPacketClientLoginSuccess(this.profile.Id, this.name))
	}
	if this.protocol17 {
		this.pipeline.Replace("registry", minecraft.PlayPacketServerCodec17)
	} else {
		this.pipeline.Replace("registry", minecraft.PlayPacketServerCodec)
	}
	this.server.SessionRegistry.Register(this)
	this.Redirect(server)
}

func (this *Session) SetEncryption(sharedSecret []byte) (err error) {
	codec, err := packet.NewPacketCodecCfb8(sharedSecret)
	if err != nil {
		return
	}
	this.pipeline.AddBefore("cfb8", "varIntLength", codec)
	return
}

func (this *Session) SetCompression(threshold int) {
	if this.compressionThreshold == threshold {
		return
	}
	this.compressionThreshold = threshold
	registry := this.pipeline.Get("registry")
	if registry == minecraft.LoginPacketServerCodec {
		this.Write(minecraft.NewPacketClientLoginSetCompression(threshold))
	} else if registry == minecraft.PlayPacketServerCodec {
		this.Write(minecraft.NewPacketClientSetCompression(threshold))
	}
	if threshold == -1 {
		this.pipeline.Remove("zlib")
		return
	} else {
		codec := packet.NewPacketCodecZlib(threshold)
		if this.pipeline.HasName("zlib") {
			this.pipeline.Replace("zlib", codec)
		} else {
			this.pipeline.AddBefore("zlib", "registry", codec)
		}
	}
}

func (this *Session) Disconnect(reason string) {
	reasonJson, _ := json.Marshal(reason)
	this.DisconnectJson("{\"text\":" + string(reasonJson) + "}")
}

func (this *Session) DisconnectJson(json string) {
	registry := this.pipeline.Get("registry")
	if registry == minecraft.LoginPacketServerCodec || registry == minecraft.LoginPacketServerCodec17 {
		this.Write(minecraft.NewPacketClientLoginDisconnect(json))
	} else if registry == minecraft.PlayPacketServerCodec || registry == minecraft.PlayPacketServerCodec17 {
		this.Write(minecraft.NewPacketClientDisconnect(json))
	}
	this.conn.Close()
}

func (this *Session) HandlePacket(packet packet.Packet) (err error) {
	switch this.state {
	case STATE_DISCONNECTED:
		if packet.Id() == minecraft.PACKET_SERVER_HANDSHAKE {
			handshakePacket := packet.(*minecraft.PacketServerHandshake)
			this.protocolVersion = handshakePacket.ProtocolVersion
			this.serverAddress = handshakePacket.ServerAddress
			supportedVersion := false
			for _, version := range minecraft.Versions {
				if version != this.protocolVersion {
					continue
				}
				supportedVersion = true
				break
			}
			if handshakePacket.State == 1 {
				if !supportedVersion {
					this.protocolVersion = minecraft.Versions[0]
				}
				this.pipeline.Replace("registry", minecraft.StatusPacketServerCodec)
				this.state = STATE_STATUS
			} else if handshakePacket.State == 2 {
				if !supportedVersion {
					err = errors.New("Protocol version does not match")
					return
				}
				this.protocol17 = this.protocolVersion < minecraft.Versions[0]
				if this.protocol17 {
					this.pipeline.Replace("registry", minecraft.LoginPacketServerCodec17)
				} else {
					this.pipeline.Replace("registry", minecraft.LoginPacketServerCodec)
				}
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
			samplePath := this.server.router.RouteSample(this.serverAddress)
			sampleTxt, sampleErr := ioutil.ReadFile(samplePath)
			sample := make([]map[string]interface{}, 0)
			if sampleErr == nil {
				lines := strings.Split(string(sampleTxt), "\n")
				for _, line := range lines {
					line = strings.Replace(line, "\r", "", -1)
					if(len(strings.TrimSpace(line)) == 0) {
						continue
					}
					entry := make(map[string]interface{})
					entry["name"] = minecraft.Colorize(line)
					entry["id"] = "00000000-0000-0000-0000-000000000000"
					sample = append(sample, entry)
				}
			}
			icons := this.server.router.RouteIcons(this.serverAddress)
			iconPath := icons[RandomInt(len(icons))]
			icon, iconErr := ioutil.ReadFile(iconPath)
			var iconString string
			if iconErr == nil {
				iconString = "data:image/png;base64," + base64.StdEncoding.EncodeToString(icon)
			}
			version := make(map[string]interface{})
			version["name"] = minecraft.STRING_VERSION
			version["protocol"] = this.protocolVersion
			players := make(map[string]interface{})
			if this.server.SyncMaxPlayers() {
				players["max"] = this.server.connect.MaxPlayers()
			} else {
				players["max"] = this.server.MaxPlayers()
			}
			players["online"] = this.server.connect.Players()
			players["sample"] = sample
			description := make(map[string]interface{})
			motds := this.server.router.RouteMotds(this.serverAddress)
			motd := motds[RandomInt(len(motds))]
			description["text"] = minecraft.Colorize(motd)
			response := make(map[string]interface{})
			response["version"] = version
			response["players"] = players
			response["description"] = description
			if iconString != "" {
				response["favicon"] = iconString
			}
			var marshalled []byte
			marshalled, err = json.Marshal(response)
			if err != nil {
				return
			}
			err = this.Write(minecraft.NewPacketClientStatusResponse(string(marshalled)))
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
			err = this.Write(minecraft.NewPacketClientStatusPing(packet.(*minecraft.PacketServerStatusPing).Time))
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
				this.publicKey, err = x509.MarshalPKIXPublicKey(&this.server.privateKey.PublicKey)
				if err != nil {
					return
				}
				this.verifyToken, err = RandomBytes(4)
				if err != nil {
					return
				}
				err = this.Write(minecraft.NewPacketClientLoginEncryptRequest(this.serverId, this.publicKey, this.verifyToken))
				if err != nil {
					return
				}
				this.state = STATE_LOGIN_ENCRYPT
			} else {
				this.profile = auth.GameProfile{
					Id: GenNameUUID("OfflinePlayer:" + this.name),
					Properties: make([]auth.GameProfileProperty, 0),
				}
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
			sharedSecret, err = rsa.DecryptPKCS1v15(rand.Reader, this.server.privateKey, loginEncryptResponsePacket.SharedSecret)
			if err != nil {
				return
			}
			var verifyToken []byte
			verifyToken, err = rsa.DecryptPKCS1v15(rand.Reader, this.server.privateKey, loginEncryptResponsePacket.VerifyToken)
			if err != nil {
				return
			}
			if bytes.Compare(this.verifyToken, verifyToken) != 0 {
				err = errors.New("Verify token does not match")
				return
			}
			err = this.SetEncryption(sharedSecret)
			if err != nil {
				return
			}
			var authErr error
			this.profile, authErr = auth.Authenticate(this.name, this.serverId, sharedSecret, this.publicKey)
			if authErr != nil {
				this.SetAuthenticated(false)
				fmt.Println("Proxy server, failed to authorize:", this.name, "ip:", this.remoteIp, "err:", authErr)
			} else {
				this.SetAuthenticated(true)
				fmt.Println("Proxy server, authorized:", this.name, "ip:", this.remoteIp)
			}
		} else {
			err = errors.New("Unexpected packet")
			return
		}
	case STATE_CONNECTED:
		if packet.Id() == minecraft.PACKET_SERVER_CLIENT_SETTINGS {
			this.clientSettings = packet
		}
		if packet.Id() == minecraft.PACKET_SERVER_PLUGIN_MESSAGE {
			pluginMessagePacket := packet.(*minecraft.PacketServerPluginMessage)
			if pluginMessagePacket.Channel == "REGISTER" {
				for _, channelBytes := range bytes.Split(pluginMessagePacket.Data[:], []byte{0}) {
					channel := string(channelBytes)
					if _, ok := this.pluginChannels[channel]; ok {
						continue
					}
					if len(this.pluginChannels) >= 128 {
						break
					}
					this.pluginChannels[channel] = struct{}{}
				}
			} else if pluginMessagePacket.Channel == "UNREGISTER" {
				for _, channelBytes := range bytes.Split(pluginMessagePacket.Data[:], []byte{0}) {
					channel := string(channelBytes)
					delete(this.pluginChannels, channel)
				}
			}
		}
		if this.redirecting {
			break
		}
		if genericPacket, ok := packet.(*minecraft.PacketGeneric); ok {
			genericPacket.SwapEntities(this.clientEntityId, this.serverEntityId, false, this.protocol17)
		}
		this.outBridge.Write(packet)
	}
	return
}

func (this *Session) ErrorCaught(err error) {
	if this.Authenticated() {
		this.server.connect.RemoveLocalPlayer(this.name)
		this.server.SessionRegistry.Unregister(this)
		fmt.Println("Proxy server, name:", this.name, "ip:", this.remoteIp, "disconnected:", err)
	}
	this.state = STATE_DISCONNECTED
	this.conn.Close()
	return
}

func (this *Session) Authenticated() (val bool) {
	val = this.state == STATE_INIT || this.state == STATE_CONNECTED
	return
}

func (this *Session) Initializing() (val bool) {
	val = this.state == STATE_INIT
	return
}
