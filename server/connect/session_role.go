package connect

type SessionRole int

const (
	ROLE_UNAUTHORIZED SessionRole = iota
	ROLE_AUTHORIZED
	ROLE_PROXY
	ROLE_SERVER
)
