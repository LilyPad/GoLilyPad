package auth

import "crypto/tls"
import "crypto/x509"
import "encoding/json"
import "errors"
import "fmt"
import "net/http"

type authenticateJson struct {
	Id string `json:"id"`
}

func Authenticate(name string, serverId string, sharedSecret []byte, publicKey []byte) (uuid string, err error) {
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
	decoder := json.NewDecoder(response.Body)
	responseJson := &authenticateJson{}
	err = decoder.Decode(responseJson)
	if err != nil {
		return
	}
	if len(responseJson.Id) != 32 {
		err = errors.New(fmt.Sprintf("Id is not 32 characters: %d", len(responseJson.Id)))
		return
	}
	uuid = responseJson.Id
	return
}
