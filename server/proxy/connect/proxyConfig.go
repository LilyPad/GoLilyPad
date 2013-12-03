package connect

type ProxyConfig struct {
	Address string
	Port uint16
	Motd *string
	Version string
	Maxplayers *uint16
}
