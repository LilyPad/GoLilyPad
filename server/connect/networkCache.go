package connect

import "sync"

type NetworkCache struct {
	playerToProxy map[string]*Session
	playerToProxyLock *sync.RWMutex
	addresses []string
	ports []uint16
	motds []string
	versions []string
	maxPlayers []uint16
	shownAddress string
	shownPort uint16
	shownMotd string
	shownVersion string
	shownMaxPlayers uint16
	rebuildLock *sync.RWMutex
}

func NewNetworkCache() *NetworkCache {
	return &NetworkCache{
		playerToProxy: make(map[string]*Session),
		playerToProxyLock: &sync.RWMutex{},
		addresses: make([]string, 0),
		ports: make([]uint16, 0),
		motds: make([]string, 0),
		versions: make([]string, 0),
		maxPlayers: make([]uint16, 0),
		rebuildLock: &sync.RWMutex{},
	}
}

func (this *NetworkCache) RegisterProxy(session *Session) {
	this.addresses = append(this.addresses, session.roleAddress)
	this.ports = append(this.ports, session.rolePort)
	this.motds = append(this.motds, session.proxyMotd)
	this.versions = append(this.versions, session.proxyVersion)
	this.maxPlayers = append(this.maxPlayers, session.proxyMaxPlayers)
	this.Rebuild()
}

func (this *NetworkCache) UnregisterProxy(session *Session) {
	for i, address := range this.addresses {
		if address != session.roleAddress {
			continue
		}
		this.addresses = append(this.addresses[:i], this.addresses[i+1:]...)
		break
	}
	for i, port := range this.ports {
		if port != session.rolePort {
			continue
		}
		this.ports = append(this.ports[:i], this.ports[i+1:]...)
		break
	}
	for i, motd := range this.motds {
		if motd != session.proxyMotd {
			continue
		}
		this.motds = append(this.motds[:i], this.motds[i+1:]...)
		break
	}
	for i, version := range this.versions {
		if version != session.proxyVersion {
			continue
		}
		this.versions = append(this.versions[:i], this.versions[i+1:]...)
		break
	}
	for i, maxPlayers := range this.maxPlayers {
		if maxPlayers != session.proxyMaxPlayers {
			continue
		}
		this.maxPlayers = append(this.maxPlayers[:i], this.maxPlayers[i+1:]...)
		break
	}
	this.Rebuild()
	this.RemovePlayersByProxy(session)
}

func (this *NetworkCache) AddPlayer(player string, session *Session) bool {
	this.playerToProxyLock.Lock()
	defer this.playerToProxyLock.Unlock()
	if _, ok := this.playerToProxy[player]; ok {
		return false
	}
	this.playerToProxy[player] = session
	return true
}

func (this *NetworkCache) RemovePlayer(player string) {
	this.playerToProxyLock.Lock()
	defer this.playerToProxyLock.Unlock()
	delete(this.playerToProxy, player)
}

func (this *NetworkCache) RemovePlayersByProxy(session *Session) {
	for player, _ := range session.proxyPlayers {
		this.RemovePlayer(player)
	}
}

func (this *NetworkCache) Players() (players []string) {
	this.playerToProxyLock.RLock()
	defer this.playerToProxyLock.RUnlock()
	players = make([]string, 0, len(this.playerToProxy))
    for player, _ := range this.playerToProxy {
        players = append(players, player)
    }
    return
}

func (this *NetworkCache) ProxyByPlayer(player string) *Session {
	this.playerToProxyLock.RLock()
	defer this.playerToProxyLock.RUnlock()
	return this.playerToProxy[player]
}

func (this *NetworkCache) Rebuild() {
	this.rebuildLock.Lock()
	defer this.rebuildLock.Unlock()
	if len(this.addresses) == 0 {
		this.shownAddress = "0.0.0.0"
	} else {
		this.shownAddress = this.addresses[0]
	}
	if len(this.ports) == 0 {
		this.shownPort = 0
	} else {
		this.shownPort = this.ports[0]
	}
	if len(this.motds) == 0 {
		this.shownMotd = "Unknown"
	} else {
		this.shownMotd = this.motds[0]
	}
	if len(this.versions) == 0 {
		this.shownVersion = "Unknown"
	} else {
		this.shownVersion = this.versions[0]
	}
	this.shownMaxPlayers = 0
	for _, maxPlayers := range this.maxPlayers {
		if maxPlayers <= 1 {
			this.shownMaxPlayers = maxPlayers
			break
		}
		this.shownMaxPlayers += maxPlayers
	}
}

func (this *NetworkCache) Address() string {
	this.rebuildLock.RLock()
	defer this.rebuildLock.RUnlock()
	return this.shownAddress
}

func (this *NetworkCache) Port() uint16 {
	this.rebuildLock.RLock()
	defer this.rebuildLock.RUnlock()
	return this.shownPort
}

func (this *NetworkCache) Motd() string {
	this.rebuildLock.RLock()
	defer this.rebuildLock.RUnlock()
	return this.shownMotd
}

func (this *NetworkCache) Version() string {
	this.rebuildLock.RLock()
	defer this.rebuildLock.RUnlock()
	return this.shownVersion
}

func (this *NetworkCache) MaxPlayers() uint16 {
	this.rebuildLock.RLock()
	defer this.rebuildLock.RUnlock()
	return this.shownMaxPlayers
}