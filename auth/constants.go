package auth

import (
  "os"
)

var URL = "https://sessionserver.mojang.com/session/minecraft/hasJoined"

func init() {
  if value, ok := os.LookupEnv("LILYPAD_MOJANG_SESSIONSERVER_URL"); ok {
    URL = value
  }
  URL = URL + "?username=%s&serverId=%s"
}
