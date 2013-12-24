package proxy

type Router interface {
	Route(domain string) (servers []string)
	RouteMotd(domain string) (motd string)
}
