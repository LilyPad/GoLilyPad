package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/LilyPad/GoLilyPad/packet/minecraft"
	"net/http"
)

func Authenticate(name string, serverId string, sharedSecret []byte, publicKey []byte) (profile minecraft.GameProfile, err error) {
	response, err := http.Get(fmt.Sprintf(URL, name, MojangSha1Hex([]byte(serverId), sharedSecret, publicKey)))
	if err != nil {
		return
	}
	jsonDecoder := json.NewDecoder(response.Body)
	profile = minecraft.GameProfile{}
	err = jsonDecoder.Decode(&profile)
	response.Body.Close()
	if err != nil {
		return
	}
	if len(profile.Id) != 32 {
		err = errors.New(fmt.Sprintf("Id is not 32 characters: %d", len(profile.Id)))
		return
	}
	if profile.Name != name {
		err = errors.New(fmt.Sprintf("Name mismatch: %s != %s", name, profile.Name))
		return
	}
	return
}
