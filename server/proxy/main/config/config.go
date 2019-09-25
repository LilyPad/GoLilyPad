package config

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

type Config struct {
	Connect ConfigConnect `yaml:"connect"`
	Proxy   ConfigProxy   `yaml:"proxy"`
}

func (this *Config) Route(domain string) (val []string) {
	if route, ok := this.Proxy.routes[strings.ToLower(domain)]; ok {
		if route.Servers != nil {
			val = make([]string, len(route.Servers))
			copy(val, route.Servers)
			return
		} else if len(route.Server) > 0 {
			val = []string{route.Server}
			return
		}
	}
	if domain != "" {
		val = this.Route("")
	} else {
		val = []string{}
	}
	return
}

func (this *Config) RouteMotds(domain string) (val []string) {
	if route, ok := this.Proxy.routes[strings.ToLower(domain)]; ok {
		if route.Motds != nil {
			val = make([]string, len(route.Motds))
			copy(val, route.Motds)
			return
		} else if len(route.Motd) > 0 {
			val = []string{route.Motd}
			return
		}
	}
	if domain != "" {
		val = this.RouteMotds("")
	} else {
		val = []string{this.Proxy.Motd}
	}
	return
}

func (this *Config) RouteIcons(domain string) (val []string) {
	if route, ok := this.Proxy.routes[strings.ToLower(domain)]; ok {
		if route.Icons != nil {
			val = make([]string, len(route.Icons))
			copy(val, route.Icons)
			return
		} else if len(route.Icon) > 0 {
			val = []string{route.Icon}
			return
		}
	}
	if domain != "" {
		val = this.RouteIcons("")
	} else {
		val = []string{"server-icon.png"}
	}
	return
}

func (this *Config) RouteSample(domain string) (val string) {
	if route, ok := this.Proxy.routes[strings.ToLower(domain)]; ok && route.Sample != "" {
		val = route.Sample
	} else if domain != "" {
		val = this.RouteSample("")
	} else {
		val = "sample.txt"
	}
	return
}

func (this *Config) routeBake() {
	if this.Proxy.routes == nil {
		this.Proxy.routes = make(map[string]ConfigProxyRoute)
		for _, route := range this.Proxy.Routes {
			if route.Domains != nil {
				for _, domain := range route.Domains {
					this.Proxy.routes[strings.ToLower(domain)] = route
				}
			}
			if route.Domains == nil || route.Domain != "" {
				this.Proxy.routes[strings.ToLower(route.Domain)] = route
			}
		}
	}
}

func (this *Config) LocaleFull() (val string) {
	val = this.Proxy.Locale.Full
	return
}

func (this *Config) LocaleOffline() (val string) {
	val = this.Proxy.Locale.Offline
	return
}

func (this *Config) LocaleLoggedIn() (val string) {
	val = this.Proxy.Locale.LoggedIn
	return
}

func (this *Config) LocaleLostConn() (val string) {
	val = this.Proxy.Locale.LostConn
	return
}

func (this *Config) LocaleShutdown() (val string) {
	val = this.Proxy.Locale.Shutdown
	return
}

type ConfigConnect struct {
	Address     string                   `yaml:"address"`
	Credentials ConfigConnectCredentials `yaml:"credentials"`
}

type ConfigConnectCredentials struct {
	Username string
	Password string
}

type ConfigProxy struct {
	Bind           string             `yaml:"bind"`
	Routes         []ConfigProxyRoute `yaml:"routes"`
	routes         map[string]ConfigProxyRoute
	Locale         ConfigProxyLocale `yaml:"locale"`
	Motd           string            `yaml:"motd"`
	MaxPlayers     uint16            `yaml:"maxPlayers"`
	SyncMaxPlayers bool              `yaml:"syncMaxPlayers"`
	Authenticate   bool              `yaml:"authenticate"`
}

type ConfigProxyLocale struct {
	Full     string `yaml:"full"`
	Offline  string `yaml:"offline"`
	LoggedIn string `yaml:"loggedIn"`
	LostConn string `yaml:"lostConn"`
	Shutdown string `yaml:"shutdown"`
}

type ConfigProxyRoute struct {
	Domain  string   `yaml:"domain,omitempty"`
	Domains []string `yaml:"domains,omitempty"`
	Server  string   `yaml:"server,omitempty"`
	Servers []string `yaml:"servers,omitempty"`
	Motd    string   `yaml:"motd,omitempty"`
	Motds   []string `yaml:"motds,omitempty"`
	Icon    string   `yaml:"icon,omitempty"`
	Icons   []string `yaml:"icons,omitempty"`
	Sample  string   `yaml:"sample,omitempty"`
}

func DefaultConfig() (config *Config) {
	config = new(Config)
	config.Connect = ConfigConnect{
		Address: "127.0.0.1:5091",
		Credentials: ConfigConnectCredentials{
			Username: "example",
			Password: "example",
		},
	}
	config.Proxy = ConfigProxy{
		Bind: ":25565",
		Routes: []ConfigProxyRoute{
			{"", nil, "example", nil, "", nil, "", nil, ""},
			{"example.com", nil, "", []string{"hub1", "hub2"}, "Example Custom MOTD", nil, "", nil, ""},
			{"hub.example.com", nil, "hub", nil, "", []string{"Example MOTD 1", "Example MOTD 2"}, "", nil, ""},
			{"icon.example.com", nil, "hub", nil, "", nil, "icon.png", []string{"icon1.png", "icon2.png", "icons/icon3.png"}, ""},
		},
		Locale: ConfigProxyLocale{
			Full:     "The server seems to be currently full. Try again later!",
			Offline:  "The requested server is currently offline. Try again later!",
			LoggedIn: "You seem to be logged in already. Try again later!",
			LostConn: "Lost connection... Please try to reconnect",
			Shutdown: "The server is being restarted. Please try to reconnect",
		},
		Motd:         "A LilyPad Server",
		MaxPlayers:   1,
		Authenticate: true,
	}
	return
}

func LoadConfig(file string) (config *Config, err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	config = new(Config)
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return
	}
	config.routeBake()
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
