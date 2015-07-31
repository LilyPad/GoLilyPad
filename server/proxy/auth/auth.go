package auth

import (
	json "github.com/pquerna/ffjson/ffjson"
	"errors"
	"fmt"
	"net/http"
	"bytes"
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
	response, err := http.Get(fmt.Sprintf(URL, name, MojangSha1Hex([]byte(serverId), sharedSecret, publicKey)))
	if err != nil {
		return
	}
	jsonDecoder := json.NewDecoder()
	profile = GameProfile{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	s := buf.Bytes()
	err = jsonDecoder.Decode(s, &profile)
	response.Body.Close()
	if err != nil {
		return
	}
	if len(profile.Id) != 32 {
		err = errors.New(fmt.Sprintf("Id is not 32 characters: %d", len(profile.Id)))
	}
	return
}
