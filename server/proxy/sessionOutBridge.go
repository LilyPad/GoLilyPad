package proxy

import (
	"bytes"
	"fmt"
	"github.com/LilyPad/GoLilyPad/packet"
	"github.com/LilyPad/GoLilyPad/packet/minecraft"
	mc17 "github.com/LilyPad/GoLilyPad/packet/minecraft/v17"
	mc19 "github.com/LilyPad/GoLilyPad/packet/minecraft/v19"
	"github.com/LilyPad/GoLilyPad/server/proxy/api"
	"github.com/LilyPad/GoLilyPad/server/proxy/connect"
	uuid "github.com/satori/go.uuid"
	"net"
	"strconv"
	"time"
)

type SessionOutBridge struct {
	session              *Session
	protocol             *minecraft.Version
	server               *connect.Server
	conn                 net.Conn
	connCodec            *packet.PacketConnCodec
	pipeline             *packet.PacketPipeline
	compressionThreshold int

	remoteIp      string
	remotePort    string
	state         SessionState
	disconnectErr error
}

func NewSessionOutBridge(session *Session, server *connect.Server, conn net.Conn) (this *SessionOutBridge) {
	this = new(SessionOutBridge)
	this.session = session
	this.protocol = this.session.protocol
	this.server = server
	this.conn = conn
	this.compressionThreshold = -1
	this.remoteIp, this.remotePort, _ = net.SplitHostPort(conn.RemoteAddr().String())
	this.state = STATE_DISCONNECTED
	return
}

func (this *SessionOutBridge) Serve() {
	this.session.activeServersLock.Lock()
	if _, ok := this.session.activeServers[this.server.Name]; ok {
		this.conn.Close()
		this.session.activeServersLock.Unlock()
		return
	}
	this.session.activeServers[this.server.Name] = struct{}{}
	this.session.activeServersLock.Unlock()

	this.pipeline = packet.NewPacketPipeline()
	this.pipeline.AddLast("varIntLength", packet.NewPacketCodecVarIntLength())
	this.pipeline.AddLast("registry", minecraft.HandshakePacketClientCodec)
	this.connCodec = packet.NewPacketConnCodec(this.conn, this.pipeline, 20*time.Second)

	inRemotePort, _ := strconv.ParseUint(this.session.remotePort, 10, 16)
	outRemotePort, _ := strconv.ParseUint(this.remotePort, 10, 16)
	loginPayload := LoginPayload{
		SecurityKey: this.server.SecurityKey,
		Host:        this.session.rawServerAddress,
		RealIp:      this.session.remoteIp,
		RealPort:    int(inRemotePort),
		Name:        this.session.name,
		UUID:        this.session.profile.Id,
		Properties:  make([]LoginPayloadProperty, 0),
	}
	for _, property := range this.session.profile.Properties {
		loginPayload.Properties = append(loginPayload.Properties, LoginPayloadProperty{property.Name, property.Value, property.Signature})
	}
	this.Write(minecraft.NewPacketServerHandshake(this.session.protocolVersion, EncodeLoginPayload(loginPayload), uint16(outRemotePort), 2))

	this.pipeline.Replace("registry", this.protocol.LoginClientCodec)
	this.Write(minecraft.NewPacketServerLoginStart(this.protocol.IdMap, this.session.name))

	this.state = STATE_LOGIN
	go this.connCodec.ReadConn(this)
}

func (this *SessionOutBridge) Write(packet packet.Packet) (err error) {
	event := this.session.server.apiEventBus.fireEventSessionPacket(this.session, &packet, api.PacketStagePre, api.PacketSubjectOutBridge, api.PacketDirectionWrite)
	if event.IsCancelled() {
		return
	}
	err = this.connCodec.Write(packet)
	if err != nil {
		return
	}
	this.session.server.apiEventBus.fireEventSessionPacket(this.session, &packet, api.PacketStageMonitor, api.PacketSubjectOutBridge, api.PacketDirectionWrite)
	return
}

func (this *SessionOutBridge) EnsureCompression() {
	this.SetCompression(this.compressionThreshold)
}

func (this *SessionOutBridge) SetCompression(threshold int) {
	if this.state == STATE_INIT || this.state == STATE_CONNECTED {
		this.session.SetCompression(threshold)
	}
	if this.compressionThreshold == threshold {
		return
	}
	this.compressionThreshold = threshold
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

func (this *SessionOutBridge) HandlePacket(packet packet.Packet) (err error) {
	if !this.session.Authenticated() {
		this.conn.Close()
		return
	}
	event := this.session.server.apiEventBus.fireEventSessionPacket(this.session, &packet, api.PacketStagePre, api.PacketSubjectOutBridge, api.PacketDirectionRead)
	if event.IsCancelled() {
		return
	}
	err = this.handlePacket(packet)
	if err != nil {
		return
	}
	this.session.server.apiEventBus.fireEventSessionPacket(this.session, &packet, api.PacketStageMonitor, api.PacketSubjectOutBridge, api.PacketDirectionRead)
	return
}

func (this *SessionOutBridge) handlePacket(packet packet.Packet) (err error) {
	switch this.state {
	case STATE_LOGIN:
		if packet.Id() == this.protocol.IdMap.PacketClientLoginSuccess {
			this.session.redirectMutex.Lock()
			this.state = STATE_INIT
			this.session.redirecting = true
			this.EnsureCompression()
			this.pipeline.Replace("registry", this.protocol.PlayClientCodec)
		} else if packet.Id() == this.protocol.IdMap.PacketClientLoginDisconnect {
			this.session.DisconnectJson(packet.(*minecraft.PacketClientLoginDisconnect).Json)
			this.conn.Close()
		} else if packet.Id() == this.protocol.IdMap.PacketClientLoginSetCompression {
			this.SetCompression(packet.(*minecraft.PacketClientLoginSetCompression).Threshold)
		} else {
			if this.session.Initializing() {
				this.session.Disconnect(fmt.Sprintf("Error: Outbound Protocol Mismatch: %d", packet.Id()))
			}
			this.conn.Close()
		}
	case STATE_INIT:
		if packet.Id() == this.protocol.IdMap.PacketClientPlayerPositionandLook {
			this.session.outBridge = this
			this.session.redirecting = false
			this.session.state = STATE_CONNECTED
			this.state = STATE_CONNECTED
			this.session.redirectMutex.Unlock()
		}
		fallthrough
	case STATE_CONNECTED:
		if this.state == STATE_CONNECTED {
			this.session.redirectMutex.Lock()
			if this.session.outBridge != this {
				this.session.outBridge.EnsureCompression()
				this.conn.Close()
				this.session.redirectMutex.Unlock()
				break
			}
			this.session.redirectMutex.Unlock()
		}
		id := packet.Id()
		if id == this.protocol.IdMap.PacketClientJoinGame {
			joinGamePacket := packet.(*minecraft.PacketClientJoinGame)
			if this.session.mcBrand != nil {
				this.Write(this.session.mcBrand)
			}
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
				this.session.Write(minecraft.NewPacketClientRespawn(this.protocol.IdMap, swapDimension, 2, 0, "DEFAULT"))
				this.session.Write(minecraft.NewPacketClientRespawn(this.protocol.IdMap, int32(joinGamePacket.Dimension), joinGamePacket.Difficulty, joinGamePacket.Gamemode, joinGamePacket.LevelType))
				if len(this.session.playerList) > 0 {
					if this.session.protocolVersion <= mc17.VersionNum {
						for player, _ := range this.session.playerList {
							this.session.Write(mc17.NewPacketClientPlayerListRemove(player))
						}
					} else {
						items := make([]minecraft.PacketClientPlayerListItem, 0, len(this.session.playerList))
						for uuidString, _ := range this.session.playerList {
							uuid, _ := uuid.FromBytes([]byte(uuidString))
							items = append(items, minecraft.PacketClientPlayerListItem{UUID: uuid})
						}
						this.session.Write(minecraft.NewPacketClientPlayerList(this.protocol.IdMap, minecraft.PACKET_CLIENT_PLAYER_LIST_ACTION_REMOVE, items))
					}
					this.session.playerList = make(map[string]struct{})
				}
				if len(this.session.scoreboards) > 0 {
					for scoreboard, _ := range this.session.scoreboards {
						this.session.Write(minecraft.NewPacketClientScoreboardObjectiveRemove(this.protocol.IdMap, scoreboard))
					}
					this.session.scoreboards = make(map[string]struct{})
				}
				if len(this.session.teams) > 0 {
					for team, _ := range this.session.teams {
						this.session.Write(minecraft.NewPacketClientTeamsRemove(this.protocol.IdMap, team))
					}
					this.session.teams = make(map[string]struct{})
				}
				if len(this.session.pluginChannels) > 0 {
					channels := make([][]byte, 0, len(this.session.pluginChannels))
					for channel, _ := range this.session.pluginChannels {
						channels = append(channels, []byte(channel))
					}
					this.Write(minecraft.NewPacketServerPluginMessage(this.protocol.IdMap, "REGISTER", bytes.Join(channels, []byte{0})))
				}
				if len(this.session.bossBars) > 0 {
					for uuidString, _ := range this.session.bossBars {
						uuid, _ := uuid.FromBytes([]byte(uuidString))
						this.session.Write(mc19.NewPacketClientBossBarRemove(uuid))
					}
				}
				return
			}
		} else if id == this.protocol.IdMap.PacketClientPlayerList {
			if this.session.protocolVersion <= mc17.VersionNum {
				playerListPacket := packet.(*mc17.PacketClientPlayerList)
				if playerListPacket.Online {
					this.session.playerList[playerListPacket.Name] = struct{}{}
				} else {
					delete(this.session.playerList, playerListPacket.Name)
				}
			} else {
				playerListPacket := packet.(*minecraft.PacketClientPlayerList)
				if playerListPacket.Action == minecraft.PACKET_CLIENT_PLAYER_LIST_ACTION_ADD {
					for _, item := range playerListPacket.Items {
						this.session.playerList[string(item.UUID[:])] = struct{}{}
					}
				} else if playerListPacket.Action == minecraft.PACKET_CLIENT_PLAYER_LIST_ACTION_REMOVE {
					for _, item := range playerListPacket.Items {
						delete(this.session.playerList, string(item.UUID[:]))
					}
				}
			}
		} else if id == this.protocol.IdMap.PacketClientScoreboardObjective {
			scoreboardPacket := packet.(*minecraft.PacketClientScoreboardObjective)
			if scoreboardPacket.Action == minecraft.PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_ADD {
				this.session.scoreboards[scoreboardPacket.Name] = struct{}{}
			} else if scoreboardPacket.Action == minecraft.PACKET_CLIENT_SCOREBOARD_OBJECTIVE_ACTION_REMOVE {
				delete(this.session.scoreboards, scoreboardPacket.Name)
			}
		} else if id == this.protocol.IdMap.PacketClientTeams {
			teamPacket := packet.(*minecraft.PacketClientTeams)
			if teamPacket.Action == minecraft.PACKET_CLIENT_TEAMS_ACTION_ADD {
				this.session.teams[teamPacket.Name] = struct{}{}
			} else if teamPacket.Action == minecraft.PACKET_CLIENT_TEAMS_ACTION_REMOVE {
				delete(this.session.teams, teamPacket.Name)
			}
		} else if id == this.protocol.IdMap.PacketClientBossBar {
			bossBarPacket := packet.(*mc19.PacketClientBossBar)
			if bossBarPacket.Action == mc19.PACKET_CLIENT_BOSS_BAR_ACTION_ADD {
				this.session.bossBars[string(bossBarPacket.UUID[:])] = struct{}{}
			} else if bossBarPacket.Action == mc19.PACKET_CLIENT_BOSS_BAR_ACTION_REMOVE {
				delete(this.session.bossBars, string(bossBarPacket.UUID[:]))
			}
		} else if id == this.protocol.IdMap.PacketClientDisconnect {
			this.state = STATE_DISCONNECTED
			this.session.DisconnectJson(packet.(*minecraft.PacketClientDisconnect).Json)
			return
		} else if id == this.protocol.IdMap.PacketClientSetCompression {
			this.SetCompression(packet.(*minecraft.PacketClientSetCompression).Threshold)
			return
		} else {
			if genericPacket, ok := packet.(*minecraft.PacketGeneric); ok {
				genericPacket.SwapEntities(this.session.clientEntityId, this.session.serverEntityId, true)
			}
		}
		this.session.Write(packet)
	}
	return
}

func (this *SessionOutBridge) ErrorCaught(err error) {
	this.session.activeServersLock.Lock()
	delete(this.session.activeServers, this.server.Name)
	this.session.activeServersLock.Unlock()
	if this.state == STATE_INIT {
		this.session.redirecting = false
		this.session.redirectMutex.Unlock()
	}
	this.disconnectErr = err
	if this.state != STATE_DISCONNECTED && this.session.outBridge == this {
		this.session.Disconnect(minecraft.Colorize(this.session.server.localizer.LocaleLostConn()))
	}
	this.session = nil
	this.server = nil
	this.state = STATE_DISCONNECTED
	this.conn.Close()
}
