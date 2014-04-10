package auth

import "crypto/tls"
import "crypto/x509"
import "encoding/json"
import "errors"
import "fmt"
import "net/http"

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
	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM([]byte(Certificate))
	client := &http.Client{
		Transport:  &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: rootCAs,
			},
		},
	}
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
