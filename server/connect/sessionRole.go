package connect

type SessionRole int

const (
	UNAUTHORIZED SessionRole = iota
	AUTHORIZED
	ROLE_PROXY
	ROLE_SERVER
)