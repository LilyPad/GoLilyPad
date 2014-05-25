package proxy

import "bytes"
import "net"
import "time"
import "strconv"
import "github.com/LilyPad/GoLilyPad/packet"
import "github.com/LilyPad/GoLilyPad/packet/minecraft"
import "github.com/LilyPad/GoLilyPad/server/proxy/connect"

type SessionOutBridge struct {
	session *Session
	server *connect.Server
	conn net.Conn
	connCodec *packet.PacketConnCodec
	codec *packet.PacketCodecVariable

	remoteHost string
	remotePort string
	state SessionState
}

func NewSessionOutBridge(session *Session, server *connect.Server, conn net.Conn) *SessionOutBridge {
	host, port, _ := net.SplitHostPort(server.Addr)
	return &SessionOutBridge{
		session: session,
		server: server,
		conn: conn,
		remoteHost: host,
		remotePort: port,
		state: STATE_DISCONNECTED,
	}
}

func (this *SessionOutBridge) Serve() {
	this.codec = packet.NewPacketCodecVariable(minecraft.HandshakePacketServerCodec, minecraft.HandshakePacketClientCodec)
	this.connCodec = packet.NewPacketConnCodec(this.conn, this.codec, 30 * time.Second)
	remotePort, _ := strconv.ParseUint(this.remotePort, 10, 8)
	this.Write(&minecraft.PacketServerHandshake{this.session.protocolVersion, this.server.SecurityKey + ";" + this.session.remoteHost + ";" + this.session.remotePort + ";" + this.session.profile.Id, uint16(remotePort), 2})
	this.codec.SetEncodeCodec(minecraft.LoginPacketServerCodec)
	this.codec.SetDecodeCodec(minecraft.LoginPacketClientCodec)
	this.Write(&minecraft.PacketServerLoginStart{this.session.name})
	this.state = STATE_LOGIN
	go this.connCodec.ReadConn(this)
}

func (this *SessionOutBridge) Write(packet packet.Packet) (err error) {
	return this.connCodec.Write(packet)
}

func (this *SessionOutBridge) HandlePacket(packet packet.Packet) (err error) {
	if !this.session.Authenticated() {
		this.conn.Close()
		return
	}
	switch this.state {
	case STATE_LOGIN:
		if packet.Id() == minecraft.PACKET_CLIENT_LOGIN_SUCCESS {
			this.session.redirectMutex.Lock()
			this.state = STATE_INIT
			this.session.redirecting = true
			this.codec.SetEncodeCodec(minecraft.PlayPacketServerCodec)
			this.codec.SetDecodeCodec(minecraft.PlayPacketClientCodec)
		} else if packet.Id() == minecraft.PACKET_CLIENT_LOGIN_DISCONNECT {
			this.session.DisconnectRaw(packet.(*minecraft.PacketClientLoginDisconnect).Json)
			this.conn.Close()
		} else {
			if this.session.Initializing() {
				this.session.Disconnect("Error: Outbound Protocol Mismatch")
			}
			this.conn.Close()
		}
	case STATE_INIT:
		if packet.Id() == minecraft.PACKET_CLIENT_JOIN_GAME {
			this.Write(this.buildPluginMessage())
			this.session.outBridge = this
			this.session.redirecting = false
			this.session.state = STATE_CONNECTED
			this.state = STATE_CONNECTED
			this.session.redirectMutex.Unlock()
		}
		fallthrough
	case STATE_CONNECTED:
		if packet.Id() == minecraft.PACKET_CLIENT_DISCONNECT {
			this.state = STATE_DISCONNECTED
		}
		if this.state == STATE_CONNECTED {
			this.session.redirectMutex.Lock()
			if this.session.outBridge != this {
				this.conn.Close()
				this.session.redirectMutex.Unlock()
				break
			}
			this.session.redirectMutex.Unlock()
		}
		switch packet.Id() {
		case minecraft.PACKET_CLIENT_JOIN_GAME:
			joinGamePacket := packet.(*minecraft.PacketClientJoinGame)
			if this.session.clientSettings != nil {
				this.Write(this.session.clientSettings)
			}
			this.session.serverEntityId = joinGamePacket.EntityId
			if this.session.state == STATE_INIT {
				this.session.clientEntityId = joinGamePacket.EntityId
			} else {
				var swapDimension int32
				if joinGamePacket.Dimension == 0 {
					swapDimension = 1
				} else {
					swapDimension = 0
				}
				this.session.Write(&minecraft.PacketClientRespawn{swapDimension, 2, 0, "DEFAULT"})
				this.session.Write(&minecraft.PacketClientRespawn{int32(joinGamePacket.Dimension), joinGamePacket.Difficulty, joinGamePacket.Gamemode, joinGamePacket.LevelType})
				for player, _ := range this.session.playerList {
					this.session.Write(&minecraft.PacketClientPlayerListItem{player, false, 0})
				}
				this.session.playerList = make(map[string]bool)
				for scoreboard, _ := range this.session.scoreboards {
					this.session.Write(&minecraft.PacketClientScoreboardObjective{scoreboard, "", 1})
				}
				this.session.scoreboards = make(map[string]bool)
				for team, _ := range this.session.teams {
					this.session.Write(&minecraft.PacketClientTeams{team, 1, "", "", "", 0, nil})
				}
				this.session.teams = make(map[string]bool)
				return
			}
		case minecraft.PACKET_CLIENT_PLAYER_LIST_ITEM:
			playerListPacket := packet.(*minecraft.PacketClientPlayerListItem)
			if playerListPacket.Online {
				this.session.playerList[playerListPacket.Name] = true
			} else {
				delete(this.session.playerList, playerListPacket.Name)
			}
		case minecraft.PACKET_CLIENT_SCOREBOARD_OBJECTIVE:
			scoreboardPacket := packet.(*minecraft.PacketClientScoreboardObjective)
			if scoreboardPacket.Action == 0 {
				this.session.scoreboards[scoreboardPacket.Name] = true
			} else if scoreboardPacket.Action == 1 {
				delete(this.session.scoreboards, scoreboardPacket.Name)
			}
		case minecraft.PACKET_CLIENT_TEAMS:
			teamPacket := packet.(*minecraft.PacketClientTeams)
			if teamPacket.Action == 0 {
				this.session.teams[teamPacket.Name] = true
			} else if teamPacket.Action == 1 {
				delete(this.session.teams, teamPacket.Name)
			}
		case minecraft.PACKET_CLIENT_DISCONNECT:
			this.session.DisconnectRaw(packet.(*minecraft.PacketClientDisconnect).Json)
			return
		default:
			if genericPacket, ok := packet.(*minecraft.PacketGeneric); ok {
				genericPacket.SwapEntities(this.session.clientEntityId, this.session.serverEntityId, true)
			}
		}
		this.session.Write(packet)
	}
	return
}

func (this *SessionOutBridge) ErrorCaught(err error) {
	if this.state == STATE_INIT {
		this.session.redirecting = false
		this.session.redirectMutex.Unlock()
	}
	if this.state != STATE_DISCONNECTED && this.session.outBridge == this {
		this.session.Disconnect(minecraft.Colorize(this.session.server.Localizer().LocaleLostConn()))
	}
	this.session = nil
	this.server = nil
	this.state = STATE_DISCONNECTED
	this.conn.Close()
}

func (this *SessionOutBridge) buildPluginMessage() *minecraft.PacketServerPluginMessage {
	buffer := &bytes.Buffer{}
	util := make([]byte, packet.UTIL_BUFFER_LENGTH)

	packet.WriteVarInt(buffer, util, len(this.session.profile.Properties))
	for _, property := range this.session.profile.Properties {
		packet.WriteString(buffer, util, property.Name)
		packet.WriteString(buffer, util, property.Value)
		packet.WriteString(buffer, util, property.Signature)
	}
	return &minecraft.PacketServerPluginMessage{"LilyPad", buffer.Bytes()}
}
