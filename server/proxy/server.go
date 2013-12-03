package proxy

import "crypto/rand"
import "crypto/rsa"
import "net"
import "github.com/LilyPad/GoLilyPad/server/proxy/connect"

type Server struct {
	listener net.Listener
	sessionRegistry *SessionRegistry

	motd *string
	maxPlayers *uint16
	authenticate *bool
	router Router
	localizer Localizer
	connect *connect.ProxyConnect
	privateKey *rsa.PrivateKey
}

func NewServer(motd *string, maxPlayers *uint16, authenticate *bool, router Router, localizer Localizer, connect *connect.ProxyConnect) (server *Server, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}
	server = &Server {
		sessionRegistry: NewSessionRegistry(),
		motd: motd,
		maxPlayers: maxPlayers,
		authenticate: authenticate,
		router: router,
		localizer: localizer,
		connect: connect,
		privateKey: privateKey,
	}
	connect.OnRedirect(func(serverName string, player string) {
		session := server.SessionRegistry().GetByName(player)
		if session == nil {
			return
		}
		server := connect.Server(serverName)
		if server == nil {
			return
		}
		session.Redirect(server)
	})
	return
}

func (this *Server) ListenAndServe(addr string) (err error) {
	defer this.Close()
	this.listener, err = net.Listen("tcp", addr)
	if err != nil {
		return
	}
	var conn net.Conn
	for {
		conn, err = this.listener.Accept()
		if err != nil {
			return
		}
		NewSession(this, conn).Serve()
	}
	return
}

func (this *Server) Close() {
	this.listener.Close()
}

func (this *Server) SessionRegistry() *SessionRegistry {
	return this.sessionRegistry
}

func (this *Server) Motd() string {
	return *this.motd
}

func (this *Server) MaxPlayers() uint16 {
	return *this.maxPlayers
}

func (this *Server) Authenticate() bool {
	return *this.authenticate
}

func (this *Server) Router() Router {
	return this.router
}

func (this *Server) Localizer() Localizer {
	return this.localizer
}

func (this *Server) Connect() *connect.ProxyConnect {
	return this.connect
}

func (this *Server) PrivateKey() *rsa.PrivateKey {
	return this.privateKey
}