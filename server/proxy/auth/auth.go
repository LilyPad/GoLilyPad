package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type GameProfile struct {
	Id string `json:"id"`
	Properties []GameProfileProperty `json:"properties"`
}

type GameProfileProperty struct {
	Name string `json:"name"`
	Value string `json:"value"`
	Signature string `json:"signature"`
}

func Authenticate(name string, serverId string, sharedSecret []byte, publicKey []byte) (profile GameProfile, err error) {
	transport := new(http.Transport)
	transport.TLSClientConfig = TLSConfig
	client := new(http.Client)
	client.Transport = transport
	response, err := client.Get(fmt.Sprintf(URL, name, MojangSha1Hex([]byte(serverId), sharedSecret, publicKey)))
	if err != nil {
		return
	}
	defer response.Body.Close()
	jsonDecoder := json.NewDecoder(response.Body)
	profile = GameProfile{}
	err = jsonDecoder.Decode(&profile)
	if err != nil {
		return
	}
	if len(profile.Id) != 32 {
		err = errors.New(fmt.Sprintf("Id is not 32 characters: %d", len(profile.Id)))
	}
	return
}
