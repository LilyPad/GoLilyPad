package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type GameProfile struct {
	Id         string                `json:"id"`
	Name       string                `json:"name"`
	Properties []GameProfileProperty `json:"properties"`
}

type GameProfileProperty struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Signature string `json:"signature"`
}

func (property *GameProfileProperty) HasSignature() bool {
	return len(property.Signature) != 0
}

func Authenticate(name string, serverId string, sharedSecret []byte, publicKey []byte) (profile GameProfile, err error) {
	response, err := http.Get(fmt.Sprintf(URL, name, MojangSha1Hex([]byte(serverId), sharedSecret, publicKey)))
	if err != nil {
		return
	}
	jsonDecoder := json.NewDecoder(response.Body)
	profile = GameProfile{}
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
