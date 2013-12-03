package connect

import "net"
import "time"
import "strconv"
import "sync"
import clientConnect "github.com/LilyPad/GoLilyPad/client/connect"
import packetConnect "github.com/LilyPad/GoLilyPad/packet/connect"

type ProxyConnect struct {
	client clientConnect.Connect
	servers map[string]*Server
	localPlayers map[string]bool
	localPlayersMutex *sync.RWMutex
	players uint16
	maxPlayers uint16
}

func NewProxyConnect(addr *string, user *string, pass *string, proxy *ProxyConfig, done chan bool) (connect *ProxyConnect) {
	client := clientConnect.NewConnect()
	connect = &ProxyConnect{
		client: client,
		localPlayers: make(map[string]bool),
		localPlayersMutex: &sync.RWMutex{},
	}
	client.RegisterEvent("preconnect", func(event clientConnect.Event) {
		connect.servers = make(map[string]*Server)
		connect.players = 0
		connect.maxPlayers = 0
	})
	client.RegisterEvent("authenticate", func(event clientConnect.Event) {
		client.RequestLater(&packetConnect.RequestAsProxy{proxy.Address, proxy.Port, *proxy.Motd, proxy.Version, *proxy.Maxplayers}, func(statusCode uint8, result packetConnect.Result) {
			if result == nil {
				return
			}
			client.DispatchEvent("roled", nil)
		})
	})
	client.RegisterEvent("server", func(event clientConnect.Event) {
		serverEvent := event.(*clientConnect.EventServer)
		if serverEvent.Add {
			var address string
			if serverEvent.Address == "127.0.0.1" || serverEvent.Address == "localhost" {
				address, _, _ = net.SplitHostPort(*addr)
			} else {
				address = serverEvent.Address
			}
			connect.servers[serverEvent.Server] = &Server{serverEvent.Server, address + ":" + strconv.FormatInt(int64(serverEvent.Port), 10), serverEvent.SecurityKey}
		} else {
			delete(connect.servers, serverEvent.Server)
		}
	})
	client.RegisterEvent("roled", func(event clientConnect.Event) {
		connect.ResendLocalPlayers()
	})
	clientConnect.AutoAuthenticate(client, user, pass)
	go clientConnect.AutoConnect(client, addr, done)
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-ticker.C:
				if !client.Connected() {
					break
				}
				connect.QueryRemotePlayers()
			case <-done:
				ticker.Stop()
				return
			}
		}
	}()
	return connect
}

func (this *ProxyConnect) AddLocalPlayer(player string) int {
	statusCode, _, err := this.client.Request(&packetConnect.RequestNotifyPlayer{true, player})
	if err != nil {
		return -1
	}
	if statusCode == packetConnect.STATUS_SUCCESS {
		this.localPlayersMutex.Lock()
		defer this.localPlayersMutex.Unlock()
		this.localPlayers[player] = true
		return 1
	} else if statusCode == packetConnect.STATUS_ERROR_GENERIC {
		return 0
	}
	return -1
}

func (this *ProxyConnect) RemoveLocalPlayer(player string) {
	this.localPlayersMutex.Lock()
	defer this.localPlayersMutex.Unlock()
	delete(this.localPlayers, player)
	this.client.RequestLater(&packetConnect.RequestNotifyPlayer{false, player}, nil)
}

func (this *ProxyConnect) ResendLocalPlayers() {
	this.localPlayersMutex.RLock()
	defer this.localPlayersMutex.RUnlock()
	for player, _ := range this.localPlayers {
		this.client.RequestLater(&packetConnect.RequestNotifyPlayer{true, player}, nil)
	}
}

func (this *ProxyConnect) QueryRemotePlayers() {
	this.client.RequestLater(&packetConnect.RequestGetPlayers{false}, func(statusCode uint8, result packetConnect.Result) {
		if result == nil {
			return
		}
		getPlayersResult := result.(*packetConnect.ResultGetPlayers)
		this.players = getPlayersResult.CurrentPlayers
		this.maxPlayers = getPlayersResult.MaximumPlayers
	})
}

func (this *ProxyConnect) Server(name string) *Server {
	if server, ok := this.servers[name]; ok {
		return server
	}
	return nil
}

func (this *ProxyConnect) Players() uint16 {
	return this.players
}

func (this *ProxyConnect) MaxPlayers() uint16 {
	return this.maxPlayers
}

func (this *ProxyConnect) OnRedirect(handler RedirectHandler) {
	this.client.RegisterEvent("redirect", func(event clientConnect.Event) {
		redirectEvent := event.(*clientConnect.EventRedirect)
		handler(redirectEvent.Server, redirectEvent.Player)
	})
}