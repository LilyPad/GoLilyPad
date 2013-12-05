package main

import "bufio"
import "fmt"
import "os"
import "github.com/LilyPad/GoLilyPad/server/connect"
import "github.com/LilyPad/GoLilyPad/server/connect/main/config"

var VERSION string

func main() {
	cfg, err := config.LoadConfig("connect.yml")
	if err != nil {
		cfg = config.DefaultConfig()
		err = config.SaveConfig("connect.yml", cfg)
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

	serverErr := make(chan error, 1)
	server := connect.NewServer(cfg)
	go func() {
		serverErr <- server.ListenAndServe(cfg.Bind)
	}()

	closeAll := func() {
		os.Stdin.Close()
		server.Close()
	}

	fmt.Println("Connect server started, version:", VERSION)

	for {
		select {
		case str := <-stdinString:
			if str == "reload\n" {
				fmt.Println("Reloading config...")
				newCfg, err := config.LoadConfig("connect.yml")
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
