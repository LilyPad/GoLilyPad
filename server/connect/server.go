package connect

import "net"

type Server struct {
	listener net.Listener
	keepaliveDone chan bool

	authenticator Authenticator
	sessionRegistries map[SessionRole]*SessionRegistry
	networkCache *NetworkCache
}

func NewServer(authenticator Authenticator) *Server {
	return &Server{
		authenticator: authenticator,
		sessionRegistries: make(map[SessionRole]*SessionRegistry),
		networkCache: NewNetworkCache(),
	}
}

func (this *Server) ListenAndServe(addr string) (err error) {
	defer this.Close()
	this.listener, err = net.Listen("tcp", addr)
	if err != nil {
		return
	}
	this.keepaliveDone = make(chan bool)
	go Keepalive(this.SessionRegistry(AUTHORIZED), this.keepaliveDone)
	var conn net.Conn
	var session *Session
	for {
		conn, err = this.listener.Accept()
		if err != nil {
			return
		}
		session, err = NewSession(this, conn)
		if err != nil {
			return
		}
		session.Serve()
	}
	return
}

func (this *Server) Close() {
	this.listener.Close()
	if this.keepaliveDone != nil {
		close(this.keepaliveDone)
	}
}

func (this *Server) Authenticator() Authenticator {
	return this.authenticator
}

func (this *Server) SessionRegistry(sessionRole SessionRole) (sessionRegistry *SessionRegistry) {
	sessionRegistry = this.sessionRegistries[sessionRole]
	if sessionRegistry == nil {
		this.sessionRegistries[sessionRole] = NewSessionRegistry()
		sessionRegistry = this.sessionRegistries[sessionRole]
	}
	return
}

func (this *Server) NetworkCache() *NetworkCache {
	return this.networkCache
}