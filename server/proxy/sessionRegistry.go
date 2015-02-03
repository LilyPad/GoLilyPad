package proxy

import (
	"sync"
	uuid "code.google.com/p/go-uuid/uuid"
)

type SessionRegistry struct {
	sync.RWMutex
	sessionsByName map[string]*Session
	sessionsByUuid map[string]*Session
}

func NewSessionRegistry() (this *SessionRegistry) {
	this = new(SessionRegistry)
	this.sessionsByName = make(map[string]*Session)
	this.Clear()
	return
}

func (this *SessionRegistry) Register(session *Session) {
	this.Lock()
	this.sessionsByName[session.name] = session
	this.sessionsByUuid[string(session.uuid)] = session
	this.Unlock()
}

func (this *SessionRegistry) Unregister(session *Session) {
	this.Lock()
	delete(this.sessionsByName, session.name)
	delete(this.sessionsByUuid, string(session.uuid))
	this.Unlock()
}

func (this *SessionRegistry) HasName(name string) (val bool) {
	this.RLock()
	_, val = this.sessionsByName[name]
	this.RUnlock()
	return
}

func (this *SessionRegistry) HasUuid(uuid uuid.UUID) (val bool) {
	this.RLock()
	_, val = this.sessionsByUuid[string(uuid)]
	this.RUnlock()
	return
}

func (this *SessionRegistry) GetByName(name string) (session *Session) {
	this.RLock()
	session = this.sessionsByName[name]
	this.RUnlock()
	return
}

func (this *SessionRegistry) GetByUuid(uuid uuid.UUID) (session *Session) {
	this.RLock()
	session = this.sessionsByUuid[string(uuid)]
	this.RUnlock()
	return
}

func (this *SessionRegistry) GetAll() (sessions []*Session) {
	this.RLock()
	sessions = make([]*Session, 0, len(this.sessionsByName))
	for  _, session := range this.sessionsByName {
		sessions = append(sessions, session)
	}
	this.RUnlock()
	return
}

func (this *SessionRegistry) Len() (val int) {
	this.RLock()
	val = len(this.sessionsByName)
	this.RUnlock()
	return
}

func (this *SessionRegistry) Clear() {
	this.sessionsByName = make(map[string]*Session)
	this.sessionsByUuid = make(map[string]*Session)
}
