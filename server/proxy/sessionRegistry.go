package proxy

import "sync"

type SessionRegistry struct {
	sync.RWMutex
	sessionByName map[string]*Session
}

func NewSessionRegistry() *SessionRegistry {
	return &SessionRegistry{
		sessionByName: make(map[string]*Session),
	}
}

func (this *SessionRegistry) Register(session *Session) {
	this.Lock()
	defer this.Unlock()
	this.sessionByName[session.Name()] = session
}

func (this *SessionRegistry) Unregister(session *Session) {
	this.Lock()
	defer this.Unlock()
	delete(this.sessionByName, session.Name())
}

func (this *SessionRegistry) HasName(name string) bool {
	this.RLock()
	defer this.RUnlock()
	return this.sessionByName[name] != nil
}

func (this *SessionRegistry) GetByName(name string) *Session {
	this.RLock()
	defer this.RUnlock()
	return this.sessionByName[name]
}

func (this *SessionRegistry) GetAll() (sessions []*Session) {
	this.RLock()
	defer this.RUnlock()
	sessions = make([]*Session, 0, len(this.sessionByName))
	for  _, session := range this.sessionByName {
   		sessions = append(sessions, session)
	}
	return
}

func (this *SessionRegistry) Len() int {
	return len(this.sessionByName)
}

func (this *SessionRegistry) Clear() {
	this.sessionByName = make(map[string]*Session)
}