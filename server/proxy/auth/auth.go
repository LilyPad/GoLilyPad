package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type GameProfile struct {
	Id         string                `json:"id"`
	Properties []GameProfileProperty `json:"properties"`
}

type GameProfileProperty struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Signature string `json:"signature"`
}

func Authenticate(name string, serverId string, sharedSecret []byte, publicKey []byte, ip string) (profile GameProfile, err error) {
	//escape the ip to it correctly parse IPv6 addresses
	httpUrl := fmt.Sprintf(URL, name, MojangSha1Hex([]byte(serverId), sharedSecret, publicKey), url.QueryEscape(ip))
	response, err := http.Get(httpUrl)
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
	}
	return
}
