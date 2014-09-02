package proxy

import (
	"crypto/rand"
	"crypto/rsa"
	"net"
	"github.com/LilyPad/GoLilyPad/server/proxy/connect"
)

type Server struct {
	listener net.Listener
	SessionRegistry *SessionRegistry

	motd *string
	maxPlayers *uint16
	syncMaxPlayers *bool
	authenticate *bool
	router Router
	localizer Localizer
	connect *connect.ProxyConnect
	privateKey *rsa.PrivateKey
}

func NewServer(motd *string, maxPlayers *uint16, syncMaxPlayers *bool, authenticate *bool, router Router, localizer Localizer, connect *connect.ProxyConnect) (this *Server, err error) {
	this = new(Server)
	this.SessionRegistry = NewSessionRegistry()
	this.motd = motd
	this.maxPlayers = maxPlayers
	this.syncMaxPlayers = syncMaxPlayers
	this.authenticate = authenticate
	this.router = router
	this.localizer = localizer
	this.connect = connect
	this.privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}
	connect.OnRedirect(func(serverName string, player string) {
		session := this.SessionRegistry.GetByName(player)
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
		go NewSession(this, conn).Serve()
	}
	return
}

func (this *Server) Close() {
	if this.listener != nil {
		this.listener.Close()
	}
}

func (this *Server) Motd() (val string) {
	val = *this.motd
	return
}

func (this *Server) MaxPlayers() (val uint16) {
	val = *this.maxPlayers
	return
}

func (this *Server) SyncMaxPlayers() (val bool) {
	val = *this.syncMaxPlayers
	return
}

func (this *Server) Authenticate() (val bool) {
	val = *this.authenticate
	return
}
