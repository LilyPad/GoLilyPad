package config

import "io/ioutil"
import "sync"
import yaml "launchpad.net/goyaml"

type Config struct {
	Connect ConfigConnect `yaml:"connect"`
	Proxy ConfigProxy `yaml:"proxy"`
}

func (this *Config) Route(domain string) []string {
	this.Proxy.routesMutex.Lock()
	if this.Proxy.routes == nil {
		this.Proxy.routes = make(map[string]ConfigProxyRoute);
		for _, route := range this.Proxy.Routes {
			this.Proxy.routes[route.Domain] = route
		}
	}
	this.Proxy.routesMutex.Unlock()
	if route, ok := this.Proxy.routes[domain]; ok {
		if route.Servers == nil {
			if len(route.Server) == 0 {
				return []string{}
			} else {
				return []string{route.Server}
			}
		} else {
			return route.Servers
		}
	}
	if domain != "" {
		return this.Route("")
	}
	return []string{}
}

func (this *Config) RouteMotd(domain string) (motd string) {
	this.Proxy.routesMutex.Lock()
	if this.Proxy.routes == nil {
		this.Proxy.routes = make(map[string]ConfigProxyRoute);
		for _, route := range this.Proxy.Routes {
			this.Proxy.routes[route.Domain] = route
		}
	}
	this.Proxy.routesMutex.Unlock()
	if route, ok := this.Proxy.routes[domain]; ok {
		if len(route.Motd) > 0 {
			return route.Motd
		}
	}
	if domain != "" {
		return this.RouteMotd("")
	}
	return this.Proxy.Motd
}

func (this *Config) LocaleFull() string {
	return this.Proxy.Locale.Full
}

func (this *Config) LocaleOffline() string {
	return this.Proxy.Locale.Offline
}

func (this *Config) LocaleLoggedIn() string {
	return this.Proxy.Locale.LoggedIn
}

func (this *Config) LocaleLostConn() string {
	return this.Proxy.Locale.LostConn
}

func (this *Config) LocaleShutdown() string {
	return this.Proxy.Locale.Shutdown
}

type ConfigConnect struct {
	Address string `yaml:"address"`
	Credentials ConfigConnectCredentials `yaml:"credentials"`
}

type ConfigConnectCredentials struct {
	Username string
	Password string
}

type ConfigProxy struct {
	Bind string `yaml:"bind"`
	Routes []ConfigProxyRoute `yaml:"routes"`
	routes map[string]ConfigProxyRoute
	routesMutex sync.RWMutex
	Locale ConfigProxyLocale `yaml:"locale"`
	Motd string `yaml:"motd"`
	MaxPlayers uint16 `yaml:"maxPlayers"`
	Authenticate bool `yaml:"authenticate"`
}

type ConfigProxyLocale struct {
	Full string `yaml:"full"`
	Offline string `yaml:"offline"`
	LoggedIn string `yaml:"loggedIn"`
	LostConn string `yaml:"lostConn"`
	Shutdown string `yaml:"shutdown"`
}

type ConfigProxyRoute struct {
	Domain string `yaml:"domain"`
	Server string `yaml:"server,omitempty"`
	Servers []string `yaml:"servers,omitempty"`
	Motd string `yaml:"motd,omitempty"`
}

func DefaultConfig() (config *Config) {
	return &Config{
		Connect: ConfigConnect {
			Address: "127.0.0.1:5091",
			Credentials: ConfigConnectCredentials{
				Username: "example",
				Password: "example",
			},
		},
		Proxy: ConfigProxy{
			Bind: ":25565",
			Routes: []ConfigProxyRoute{
				ConfigProxyRoute{"", "example", nil, ""},
				ConfigProxyRoute{"example.com", "", []string{"hub1", "hub2"}, "Example Custom MOTD"},
			},
			Locale: ConfigProxyLocale{
				Full: "The server seems to be currently full. Try again later!",
				Offline: "The requested server is currently offline. Try again later!",
				LoggedIn: "You seem to be logged in already. Try again later!",
				LostConn: "Lost connection... Please try to reconnect",
				Shutdown: "The server is being restarted. Please try to reconnect",
			},
			Motd: "A LilyPad Server",
			MaxPlayers: 1,
			Authenticate: true,
		},
	}
}

func LoadConfig(file string) (config *Config, err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	var cfg Config
	config = &cfg
	err = yaml.Unmarshal(data, config)
	return
}

func SaveConfig(file string, config *Config) (err error) {
	data, err := yaml.Marshal(config)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(file, data, 0644)
	return
}
