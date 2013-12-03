package connect

import "sync"

type SessionRegistry struct {
	sync.RWMutex
	sessionById map[string]*Session
}

func NewSessionRegistry() *SessionRegistry {
	return &SessionRegistry{
		sessionById: make(map[string]*Session),
	}
}

func (this *SessionRegistry) Register(session *Session) {
	this.Lock()
	defer this.Unlock()
	this.sessionById[session.Id()] = session
}

func (this *SessionRegistry) Unregister(session *Session) {
	this.Lock()
	defer this.Unlock()
	delete(this.sessionById, session.Id())
}

func (this *SessionRegistry) HasId(id string) bool {
	this.RLock()
	defer this.RUnlock()
	return this.sessionById[id] != nil
}

func (this *SessionRegistry) GetById(id string) *Session {
	this.RLock()
	defer this.RUnlock()
	return this.sessionById[id]
}

func (this *SessionRegistry) GetAll() (sessions []*Session) {
	this.RLock()
	defer this.RUnlock()
	sessions = make([]*Session, 0, len(this.sessionById))
	for  _, session := range this.sessionById {
   		sessions = append(sessions, session)
	}
	return
}

func (this *SessionRegistry) Clear() {
	this.sessionById = make(map[string]*Session)
}