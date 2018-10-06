package main

import (
	"bufio"
	"fmt"
	"github.com/LilyPad/GoLilyPad/server/proxy"
	"github.com/LilyPad/GoLilyPad/server/proxy/connect"
	"github.com/LilyPad/GoLilyPad/server/proxy/main/config"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var VERSION string

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

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
			if err == io.EOF {
				time.Sleep(100 * time.Millisecond)
				continue
			}
			if err != nil {
				stdinErr <- err
			}
			stdinString <- str
		}
	}()

	connectDone := make(chan bool)
	bindAddr := strings.Split(cfg.Proxy.Bind, ":")[0]
	bindPort, _ := strconv.ParseInt(strings.Split(cfg.Proxy.Bind, ":")[1], 10, 8)
	proxyConfig := connect.NewProxyConfig(bindAddr, uint16(bindPort), &cfg.Proxy.Motd, proxy.MinecraftVersion(), &cfg.Proxy.MaxPlayers)
	proxyConnect := connect.NewProxyConnect(&cfg.Connect.Address, &cfg.Connect.Credentials.Username, &cfg.Connect.Credentials.Password, proxyConfig, connectDone)

	serverErr := make(chan error, 1)
	var server *proxy.Server
	go func() {
		var err error
		server, err = proxy.NewServer(&cfg.Proxy.Bind, &cfg.Proxy.Motd, &cfg.Proxy.MaxPlayers, &cfg.Proxy.SyncMaxPlayers, &cfg.Proxy.Authenticate, cfg, cfg, proxyConnect)
		if err != nil {
			serverErr <- err
			return
		}
		serverErr <- server.ListenAndServe()
	}()

	closeAll := func() {
		close(connectDone)
		os.Stdin.Close()
		if server != nil {
			server.Close()
		}
	}

	fmt.Println("Proxy server started, version:", strings.Replace(VERSION, "_", " ", -1))
	for {
		select {
		case str := <-stdinString:
			str = strings.TrimSpace(str)
			switch str {
			case "reload":
				fmt.Println("Reloading config...")
				newCfg, err := config.LoadConfig("proxy.yml")
				if err != nil {
					fmt.Println("Error during reloading config", err)
					continue
				} else {
					fmt.Println("Reloaded config")
				}
				*cfg = *newCfg
			case "debug":
				fmt.Println("runtime.NumCPU:", runtime.NumCPU())
				fmt.Println("runtime.NumGoroutine:", runtime.NumGoroutine())
				memStats := new(runtime.MemStats)
				runtime.ReadMemStats(memStats)
				fmt.Println("runtime.MemStats.Alloc:", memStats.Alloc, "bytes")
				fmt.Println("runtime.MemStats.TotalAlloc:", memStats.TotalAlloc, "bytes")
			case "exit", "stop", "halt":
				fmt.Println("Stopping...")
				closeAll()
				return
			case "help":
				fmt.Println("LilyPad Proxy - Help")
				fmt.Println("reload - Reloads the proxy.yml")
				fmt.Println("debug  - Prints out CPU, Memory, and Routine stats")
				fmt.Println("stop   - Stops the process. (Aliases: 'exit', 'halt')")
			default:
				fmt.Println("Command not found. Use \"help\" to view available commands.")
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
