package proxy

type SessionState int

const (
	STATE_DISCONNECTED SessionState = iota
	STATE_STATUS
	STATE_LOGIN
	STATE_LOGIN_ENCRYPT
	STATE_INIT
	STATE_CONNECTED
)
