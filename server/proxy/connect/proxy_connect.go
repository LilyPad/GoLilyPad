package connect

import (
	clientConnect "github.com/LilyPad/GoLilyPad/client/connect"
	packetConnect "github.com/LilyPad/GoLilyPad/packet/connect"
	uuid "github.com/satori/go.uuid"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ProxyConnect struct {
	client            clientConnect.Connect
	servers           map[string]*Server
	serversMutex      sync.RWMutex
	localPlayers      map[string]uuid.UUID
	localPlayersMutex sync.RWMutex
	players           uint16
	maxPlayers        uint16
}

func NewProxyConnect(addr *string, user *string, pass *string, proxy *ProxyConfig, done chan bool) (this *ProxyConnect) {
	this = new(ProxyConnect)
	this.client = clientConnect.NewConnectImpl()
	this.servers = make(map[string]*Server)
	this.localPlayers = make(map[string]uuid.UUID)
	this.client.RegisterEvent("preconnect", func(event clientConnect.Event) {
		if len(this.servers) > 0 {
			this.serversMutex.Lock()
			this.servers = make(map[string]*Server)
			this.serversMutex.Unlock()
		}
		this.players = 0
		this.maxPlayers = 0
	})
	this.client.RegisterEvent("authenticate", func(event clientConnect.Event) {
		this.client.RequestLater(packetConnect.NewRequestAsProxy(proxy.Address, proxy.Port, *proxy.Motd, proxy.Version, *proxy.Maxplayers), func(statusCode uint8, result packetConnect.Result) {
			if result == nil {
				return
			}
			this.ResendLocalPlayers()
		})
	})
	this.client.RegisterEvent("server", func(event clientConnect.Event) {
		serverEvent := event.(*clientConnect.EventServer)
		if serverEvent.Add {
			var address string
			if serverEvent.Address == "127.0.0.1" || serverEvent.Address == "localhost" {
				address, _, _ = net.SplitHostPort(*addr)
			} else if strings.Contains(serverEvent.Address, ":") {
				address = "[" + serverEvent.Address + "]"
			} else {
				address = serverEvent.Address
			}
			server := new(Server)
			server.Name = serverEvent.Server
			server.Addr = address + ":" + strconv.FormatInt(int64(serverEvent.Port), 10)
			server.SecurityKey = serverEvent.SecurityKey
			this.serversMutex.Lock()
			this.servers[serverEvent.Server] = server
			this.serversMutex.Unlock()
		} else {
			this.serversMutex.Lock()
			delete(this.servers, serverEvent.Server)
			this.serversMutex.Unlock()
		}
	})
	clientConnect.AutoAuthenticate(this.client, user, pass)
	go clientConnect.AutoConnect(this.client, addr, done)
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				if !this.client.Connected() {
					break
				}
				this.QueryRemotePlayers()
			case <-done:
				ticker.Stop()
				return
			}
		}
	}()
	return
}

func (this *ProxyConnect) AddLocalPlayer(player string, uuid uuid.UUID) (ok int) {
	statusCode, _, err := this.client.Request(packetConnect.NewRequestNotifyPlayerAdd(player, uuid))
	if err != nil {
		ok = -1
	} else if statusCode == packetConnect.STATUS_SUCCESS {
		this.localPlayersMutex.Lock()
		this.localPlayers[player] = uuid
		this.localPlayersMutex.Unlock()
		ok = 1
	} else if statusCode == packetConnect.STATUS_ERROR_GENERIC {
		ok = 0
	}
	return
}

func (this *ProxyConnect) RemoveLocalPlayer(player string, uuid uuid.UUID) {
	this.localPlayersMutex.Lock()
	delete(this.localPlayers, player)
	this.localPlayersMutex.Unlock()
	this.client.RequestLater(packetConnect.NewRequestNotifyPlayerRemove(player, uuid), nil)
}

func (this *ProxyConnect) ResendLocalPlayers() {
	this.localPlayersMutex.RLock()
	for player, uuid := range this.localPlayers {
		this.client.RequestLater(packetConnect.NewRequestNotifyPlayerAdd(player, uuid), nil)
	}
	this.localPlayersMutex.RUnlock()
}

func (this *ProxyConnect) QueryRemotePlayers() {
	this.client.RequestLater(packetConnect.NewRequestGetPlayers(), func(statusCode uint8, result packetConnect.Result) {
		if result == nil {
			return
		}
		getPlayersResult := result.(*packetConnect.ResultGetPlayers)
		this.players = getPlayersResult.CurrentPlayers
		this.maxPlayers = getPlayersResult.MaxPlayers
	})
}

func (this *ProxyConnect) Server(name string) (val *Server) {
	this.serversMutex.RLock()
	val, _ = this.servers[name]
	this.serversMutex.RUnlock()
	return
}

func (this *ProxyConnect) HasServer(name string) (val bool) {
	this.serversMutex.RLock()
	_, val = this.servers[name]
	this.serversMutex.RUnlock()
	return
}

func (this *ProxyConnect) Players() (val uint16) {
	val = this.players
	return
}

func (this *ProxyConnect) MaxPlayers() (val uint16) {
	val = this.maxPlayers
	return
}

func (this *ProxyConnect) OnRedirect(handler RedirectHandler) {
	this.client.RegisterEvent("redirect", func(event clientConnect.Event) {
		redirectEvent := event.(*clientConnect.EventRedirect)
		handler(redirectEvent.Server, redirectEvent.Player)
	})
}
