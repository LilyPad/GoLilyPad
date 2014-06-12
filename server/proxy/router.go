package proxy

type Router interface {
	Route(domain string) (servers []string)
	RouteMotds(domain string) (motds []string)
	RouteIcons(domain string) (iconPath []string)
	RouteSample(domain string) (samplePath string)
}
