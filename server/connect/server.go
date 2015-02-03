package connect

import (
	"net"
)

type Server struct {
	listener net.Listener
	keepaliveDone chan bool

	authenticator Authenticator
	sessionRegistries map[SessionRole]*SessionRegistry
	networkCache *NetworkCache
}

func NewServer(authenticator Authenticator) (this *Server) {
	this = new(Server)
	this.authenticator = authenticator
	this.sessionRegistries = make(map[SessionRole]*SessionRegistry)
	this.networkCache = NewNetworkCache()
	return
}

func (this *Server) ListenAndServe(addr string) (err error) {
	this.listener, err = net.Listen("tcp", addr)
	if err != nil {
		return
	}
	this.keepaliveDone = make(chan bool)
	go Keepalive(this.SessionRegistry(ROLE_AUTHORIZED), this.keepaliveDone)
	var conn net.Conn
	for {
		conn, err = this.listener.Accept()
		if err != nil {
			return
		}
		go NewSession(this, conn).Serve()
	}
	this.Close()
	return
}

func (this *Server) Close() {
	if this.listener != nil {
		this.listener.Close()
	}
	if this.keepaliveDone != nil {
		close(this.keepaliveDone)
	}
}

func (this *Server) SessionRegistry(sessionRole SessionRole) (sessionRegistry *SessionRegistry) {
	sessionRegistry = this.sessionRegistries[sessionRole]
	if sessionRegistry == nil {
		this.sessionRegistries[sessionRole] = NewSessionRegistry()
		sessionRegistry = this.sessionRegistries[sessionRole]
	}
	return
}
