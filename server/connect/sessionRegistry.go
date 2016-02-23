package connect

import (
	"sync"
)

type SessionRegistry struct {
	sync.RWMutex
	sessionsByName map[string]*Session
}

func NewSessionRegistry() (this *SessionRegistry) {
	this = new(SessionRegistry)
	this.Clear()
	return
}

func (this *SessionRegistry) Register(session *Session) {
	this.Lock()
	this.sessionsByName[session.Id()] = session
	this.Unlock()
}

func (this *SessionRegistry) Unregister(session *Session) {
	this.Lock()
	delete(this.sessionsByName, session.Id())
	this.Unlock()
}

func (this *SessionRegistry) HasId(id string) (val bool) {
	this.RLock()
	_, val = this.sessionsByName[id]
	this.RUnlock()
	return
}

func (this *SessionRegistry) GetById(id string) (session *Session) {
	this.RLock()
	session = this.sessionsByName[id]
	this.RUnlock()
	return
}

func (this *SessionRegistry) GetAll() (sessions []*Session) {
	this.RLock()
	sessions = make([]*Session, 0, len(this.sessionsByName))
	for _, session := range this.sessionsByName {
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
	this.Lock()
	this.sessionsByName = make(map[string]*Session)
	this.Unlock()
}
