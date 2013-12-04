package config

import "io/ioutil"
import yaml "launchpad.net/goyaml"

type Config struct {
	Connect ConfigConnect `yaml:"connect"`
	Proxy ConfigProxy `yaml:"proxy"`
}

func (this *Config) Route(domain string) string {
	if this.Proxy.routes == nil {
		this.Proxy.routes = make(map[string]string);
		for _, route := range this.Proxy.Routes {
			this.Proxy.routes[route.Domain] = route.Server
		}
	}
	if server, ok := this.Proxy.routes[domain]; ok {
		return server
	}
	if domain != "" {
		return this.Route("")
	}
	return ""
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
	routes map[string]string
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
	Server string `yaml:"server"`
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
				ConfigProxyRoute{"", "example"},
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
