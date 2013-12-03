package connect

import "time"

func Keepalive(sessionRegistry *SessionRegistry, done chan bool) {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			for _, session := range sessionRegistry.GetAll() {
				session.Keepalive()
			}
		case <-done:
			ticker.Stop()
			return
		}
	}
}