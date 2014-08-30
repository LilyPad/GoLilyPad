package connect

import (
	"net"
	"time"
	"strconv"
	"sync"
	clientConnect "github.com/LilyPad/GoLilyPad/client/connect"
	packetConnect "github.com/LilyPad/GoLilyPad/packet/connect"
)

type ProxyConnect struct {
	client clientConnect.Connect
	servers map[string]*Server
	serversMutex sync.RWMutex
	localPlayers map[string]struct{}
	localPlayersMutex sync.RWMutex
	players uint16
	maxPlayers uint16
}

func NewProxyConnect(addr *string, user *string, pass *string, proxy *ProxyConfig, done chan bool) (this *ProxyConnect) {
	this = new(ProxyConnect)
	this.client = clientConnect.NewConnectImpl()
	this.servers = make(map[string]*Server)
	this.localPlayers = make(map[string]struct{})
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
			} else {
				address = serverEvent.Address
			}
			server := new(Server)
			server.Name = serverEvent.Server
			server.Addr = address + ":" + strconv.FormatInt(int64(serverEvent.Port), 10)
			server.SecurityKey = serverEvent.SecurityKey
			this.serversMutex.Lock()
			defer this.serversMutex.Unlock()
			this.servers[serverEvent.Server] = server
		} else {
			this.serversMutex.Lock()
			defer this.serversMutex.Unlock()
			delete(this.servers, serverEvent.Server)
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

func (this *ProxyConnect) AddLocalPlayer(player string) (ok int) {
	statusCode, _, err := this.client.Request(packetConnect.NewRequestNotifyPlayerAdd(player))
	if err != nil {
		ok = -1
	} else if statusCode == packetConnect.STATUS_SUCCESS {
		this.localPlayersMutex.Lock()
		defer this.localPlayersMutex.Unlock()
		this.localPlayers[player] = struct{}{}
		ok = 1
	} else if statusCode == packetConnect.STATUS_ERROR_GENERIC {
		ok = 0
	}
	return
}

func (this *ProxyConnect) RemoveLocalPlayer(player string) {
	this.localPlayersMutex.Lock()
	defer this.localPlayersMutex.Unlock()
	delete(this.localPlayers, player)
	this.client.RequestLater(packetConnect.NewRequestNotifyPlayerRemove(player), nil)
}

func (this *ProxyConnect) ResendLocalPlayers() {
	this.localPlayersMutex.RLock()
	defer this.localPlayersMutex.RUnlock()
	for player, _ := range this.localPlayers {
		this.client.RequestLater(packetConnect.NewRequestNotifyPlayerAdd(player), nil)
	}
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

func (this *ProxyConnect) HasServer(name string) (val bool) {
	this.serversMutex.RLock()
	defer this.serversMutex.RUnlock()
	_, val = this.servers[name]
	return
}

func (this *ProxyConnect) Server(name string) (val *Server) {
	this.serversMutex.RLock()
	defer this.serversMutex.RUnlock()
	val, _ = this.servers[name]
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
