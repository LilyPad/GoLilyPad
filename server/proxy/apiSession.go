package proxy

import (
	"github.com/LilyPad/GoLilyPad/packet"
	"github.com/LilyPad/GoLilyPad/server/proxy/api"
	uuid "github.com/satori/go.uuid"
	"net"
)

type apiSession struct {
	session *Session
}

func (this *apiSession) Conn() net.Conn {
	return this.session.conn
}

func (this *apiSession) Write(packet packet.Packet, subject api.PacketSubject) {
	if subject == api.PacketSubjectClient {
		this.session.Write(packet)
	} else if subject == api.PacketSubjectOutBridge {
		if this.session.outBridge == nil {
			return
		}
		this.session.outBridge.Write(packet)
	}
}

func (this *apiSession) Profile() (name string, uuid uuid.UUID) {
	return this.session.name, this.session.uuid
}

func (this *apiSession) Disconnect(reason string) {
	this.session.Disconnect(reason)
}

func (this *apiSession) DisconnectJson(json string) {
	this.session.DisconnectJson(json)
}

func (this *apiSession) Remote() (ip string, port string) {
	ip, port = this.session.Remote()
	return
}

func (this *apiSession) RemoteOverride(ip string, port string) {
	this.session.RemoteOverride(ip, port)
}

func (this *apiSession) State() api.SessionState {
	return api.SessionState(this.session.state)
}

type apiSessionRegistry struct {
	sessionRegistry *SessionRegistry
}

func (this *apiSessionRegistry) HasName(name string) (val bool) {
	return this.sessionRegistry.HasName(name)
}

func (this *apiSessionRegistry) HasUuid(uuid uuid.UUID) (val bool) {
	return this.sessionRegistry.HasUuid(uuid)
}

func (this *apiSessionRegistry) GetByName(name string) (session api.Session) {
	return this.sessionRegistry.GetByName(name).apiSession
}

func (this *apiSessionRegistry) GetByUuid(uuid uuid.UUID) (session api.Session) {
	return this.sessionRegistry.GetByUuid(uuid).apiSession
}

func (this *apiSessionRegistry) GetAll() (sessions []api.Session) {
	sessionsRaw := this.sessionRegistry.GetAll()
	sessions = make([]api.Session, len(sessionsRaw))
	for i, session := range sessionsRaw {
		sessions[i] = session.apiSession
	}
	return
}

func (this *apiSessionRegistry) Len() (val int) {
	return this.sessionRegistry.Len()
}
