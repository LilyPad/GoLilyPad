package proxy

import (
	cryptoRand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"github.com/LilyPad/GoLilyPad/server/proxy/api"
	"github.com/LilyPad/GoLilyPad/server/proxy/connect"
	"io/ioutil"
	"net"
	"path/filepath"
	"plugin"
	"strings"
	"time"
)

type Server struct {
	listener        net.Listener
	listenerAddr    string
	SessionRegistry *SessionRegistry

	apiContext  *apiContext
	apiEventBus *eventBus

	motd           *string
	maxPlayers     *uint16
	syncMaxPlayers *bool
	authenticate   *bool
	router         Router
	localizer      Localizer
	connect        *connect.ProxyConnect
	privateKey     *rsa.PrivateKey
	publicKey      []byte
}

func NewServer(motd *string, maxPlayers *uint16, syncMaxPlayers *bool, authenticate *bool, router Router, localizer Localizer, connect *connect.ProxyConnect) (this *Server, err error) {
	this = new(Server)
	this.SessionRegistry = NewSessionRegistry()
	this.apiContext = NewAPIContext(this)
	this.apiEventBus = NewEventBus()
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
	this.loadPlugins()
	return
}

func (this *Server) ListenAndServe(addr string) (err error) {
	this.listenerAddr = addr
	this.listener, err = net.Listen("tcp", this.listenerAddr)
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

func (this *Server) loadPlugins() {
	var context api.Context
	context = this.apiContext
	fileDir := "./plugins/"
	files, err := ioutil.ReadDir(fileDir)
	if err != nil {
		return
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), ".so") {
			continue
		}
		pluginOpen, err := plugin.Open(filepath.Join(fileDir, file.Name()))
		if err != nil {
			fmt.Println("Plugin load error, file:", file.Name(), "error:", err)
			continue
		}
		pluginHandle, err := pluginOpen.Lookup("Plugin")
		if err != nil {
			fmt.Println("Plugin init error, file:", file.Name(), "error:", err)
			continue
		}
		pluginHandle.(api.Plugin).Init(context)
	}
}

func (this *Server) Close() {
	if this.listener != nil {
		this.listener.Close()
	}
}

func (this *Server) ListenAddr() (val string) {
	val = this.listenerAddr
	return
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
