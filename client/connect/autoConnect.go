package connect

import "fmt"
import "time"
import "github.com/LilyPad/GoLilyPad/packet/connect"

func AutoConnect(connectClient Connect, addr *string, done chan bool) {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			if connectClient.Connected() {
				break
			}
			connectClient.DispatchEvent("preconnect", nil)
			err := connectClient.Connect(*addr)
			if err != nil {
				fmt.Println("Connect client, Couldn't connect to remote host:", err)
			}
			connectClient.DispatchEvent("connect", nil)
		case <-done:
			ticker.Stop()
			return
		}
	}
}

func AutoAuthenticate(connectClient Connect, user *string, pass *string) {
	connectClient.RegisterEvent("connect", func(event Event) {
		connectClient.RequestLater(&connect.RequestGetSalt{}, func(statusCode uint8, result connect.Result) {
			if result == nil {
				fmt.Println("Connect client, Connection timed out while keying")
				return
			}
			connectClient.RequestLater(&connect.RequestAuthenticate{*user, PasswordAndSaltHash(*pass, result.(*connect.ResultGetSalt).Salt)}, func(statusCode uint8, result connect.Result) {
				if statusCode == connect.STATUS_ERROR_GENERIC {
					fmt.Println("Connect client, Invalid username or password")
					return
				}
				if result == nil {
					fmt.Println("Connect client, Connection timed out while authenticating")
					return
				}
				connectClient.DispatchEvent("authenticate", nil)
				fmt.Println("Connect client, authenticated to remote host")
			})
		});
	});
}
