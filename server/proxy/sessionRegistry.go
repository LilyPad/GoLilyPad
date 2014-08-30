package proxy

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
	this.sessions[session.name] = session
}

func (this *SessionRegistry) Unregister(session *Session) {
	this.Lock()
	defer this.Unlock()
	delete(this.sessions, session.name)
}

func (this *SessionRegistry) HasName(name string) (val bool) {
	this.RLock()
	defer this.RUnlock()
	val = this.sessions[name] != nil
	return
}

func (this *SessionRegistry) GetByName(name string) (session *Session) {
	this.RLock()
	defer this.RUnlock()
	session = this.sessions[name]
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
	this.sessions = make(map[string]*Session)
}
