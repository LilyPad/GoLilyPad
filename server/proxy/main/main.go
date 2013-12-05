package main

import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings"
import "github.com/LilyPad/GoLilyPad/server/proxy"
import "github.com/LilyPad/GoLilyPad/server/proxy/connect"
import "github.com/LilyPad/GoLilyPad/server/proxy/main/config"

var VERSION string

func main() {
	cfg, err := config.LoadConfig("proxy.yml")
	if err != nil {
		cfg = config.DefaultConfig()
		err = config.SaveConfig("proxy.yml", cfg)
		if err != nil {
			fmt.Println("Error while saving config", err)
			return
		}
	}

	stdinString := make(chan string, 1)
	stdinErr := make(chan error, 1)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			str, err := reader.ReadString('\n')
			if err != nil {
				stdinErr <- err
			}
			stdinString <- str
		}
	}()

	connectDone := make(chan bool)
	bindAddr := strings.Split(cfg.Proxy.Bind, ":")[0]
	bindPort, _ := strconv.ParseInt(strings.Split(cfg.Proxy.Bind, ":")[1], 10, 8)
	proxyConfig := &connect.ProxyConfig{bindAddr, uint16(bindPort), &cfg.Proxy.Motd, proxy.MinecraftVersion(), &cfg.Proxy.MaxPlayers}
	proxyConnect := connect.NewProxyConnect(&cfg.Connect.Address, &cfg.Connect.Credentials.Username, &cfg.Connect.Credentials.Password, proxyConfig, connectDone)

	serverErr := make(chan error, 1)
	var server *proxy.Server
	go func() {
		var err error
		server, err = proxy.NewServer(&cfg.Proxy.Motd, &cfg.Proxy.MaxPlayers, &cfg.Proxy.Authenticate, cfg, cfg, proxyConnect)
		if err != nil {
			serverErr <- err
			return
		}
		serverErr <- server.ListenAndServe(cfg.Proxy.Bind)
	}()

	closeAll := func() {
		close(connectDone)
		os.Stdin.Close()
		if server != nil {
			server.Close()
		}
	}

	fmt.Println("Proxy server started, version:", VERSION)

	for {
		select {
		case str := <-stdinString:
			if str == "reload\n" {
				fmt.Println("Reloading config...")
				newCfg, err := config.LoadConfig("proxy.yml")
				if err != nil {
					fmt.Println("Error during reloading config", err)
					continue
				}
				*cfg = *newCfg
			} else if str == "exit\n" || str == "stop\n" || str == "halt\n" {
				fmt.Println("Stopping...")
				closeAll()
				return
			}
		case err := <-stdinErr:
			fmt.Println("Error during stdin", err)
			closeAll()
			return
		case err := <-serverErr:
			fmt.Println("Error during listen", err)
			closeAll()
			return
		}
	}
}
