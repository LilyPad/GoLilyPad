package proxy

import (
	"github.com/LilyPad/GoLilyPad/server/proxy/api"
)

type apiContext struct {
	server          *Server
	config          api.Config
	sessionRegistry api.SessionRegistry
}

func NewAPIContext(server *Server) *apiContext {
	return &apiContext{
		server:          server,
		config:          &apiConfig{server},
		sessionRegistry: &apiSessionRegistry{server.sessionRegistry},
	}
}

func (this *apiContext) Config() api.Config {
	return this.config
}

func (this *apiContext) EventBus() api.EventBus {
	return this.server.apiEventBus
}

func (this *apiContext) SessionRegistry() api.SessionRegistry {
	return this.sessionRegistry
}

type apiConfig struct {
	server *Server
}

func (this *apiConfig) ListenAddr() (val string) {
	return this.server.ListenAddr()
}

func (this *apiConfig) Motd() (val string) {
	return this.server.Motd()
}

func (this *apiConfig) MaxPlayers() (val uint16) {
	return this.server.MaxPlayers()
}

func (this *apiConfig) SyncMaxPlayers() (val bool) {
	return this.server.SyncMaxPlayers()
}

func (this *apiConfig) Authenticate() (val bool) {
	return this.server.Authenticate()
}
