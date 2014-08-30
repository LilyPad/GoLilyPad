package connect

import (
	"sync"
)

type SessionRegistry struct {
	sync.RWMutex
	sessions map[string]*Session
}

func NewSessionRegistry() (this *SessionRegistry) {
	this = new(SessionRegistry)
	this.sessions = make(map[string]*Session)
	return
}

func (this *SessionRegistry) Register(session *Session) {
	this.Lock()
	defer this.Unlock()
	this.sessions[session.Id()] = session
}

func (this *SessionRegistry) Unregister(session *Session) {
	this.Lock()
	defer this.Unlock()
	delete(this.sessions, session.Id())
}

func (this *SessionRegistry) HasId(id string) (val bool) {
	this.RLock()
	defer this.RUnlock()
	val = this.sessions[id] != nil
	return
}

func (this *SessionRegistry) GetById(id string) (session *Session) {
	this.RLock()
	defer this.RUnlock()
	session = this.sessions[id]
	return
}

func (this *SessionRegistry) GetAll() (sessions []*Session) {
	this.RLock()
	defer this.RUnlock()
	sessions = make([]*Session, 0, len(this.sessions))
	for  _, session := range this.sessions {
		sessions = append(sessions, session)
	}
	return
}

func (this *SessionRegistry) Len() (val int) {
	this.RLock()
	defer this.RUnlock()
	val = len(this.sessions)
	return
}

func (this *SessionRegistry) Clear() {
	this.Lock()
	defer this.Unlock()
	this.sessions = make(map[string]*Session)
}
