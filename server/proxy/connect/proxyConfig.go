package connect

type ProxyConfig struct {
	Address string
	Port uint16
	Motd *string
	Version string
	Maxplayers *uint16
}

func NewProxyConfig(address string, port uint16, motd *string, version string, maxplayers *uint16) (this *ProxyConfig) {
	this = new(ProxyConfig)
	this.Address = address
	this.Port = port
	this.Motd = motd
	this.Version = version
	this.Maxplayers = maxplayers
	return
}
