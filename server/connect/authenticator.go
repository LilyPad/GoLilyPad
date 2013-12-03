package connect

type Authenticator interface {
	Authenticate(username string, password string, passwordSalt string) (ok bool, err error)
}