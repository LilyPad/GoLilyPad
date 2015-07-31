package proxy

import (
	cryptoRand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"net"
	"time"
	"github.com/suedadam/GoLilyPad/server/proxy/connect"
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
	publicKey []byte
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
	this.privateKey, err = rsa.GenerateKey(cryptoRand.Reader, 2048)
	if err != nil {
		return
	}
	this.publicKey, err = x509.MarshalPKIXPublicKey(&this.privateKey.PublicKey)
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
	this.listener, err = net.Listen("tcp", addr)
	if err != nil {
		return
	}
	var conn net.Conn
	for {
		conn, err = this.listener.Accept()
		if err != nil {
			if neterr, ok := err.(net.Error); ok && neterr.Temporary() {
				time.Sleep(time.Second)
				continue
			}
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
