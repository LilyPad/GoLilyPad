package proxy

type Router interface {
	Route(domain string) (server string)
}
