package proxy

import "encoding/json"

type LoginPayload struct {
	SecurityKey string `json:"s"`
	Host string `json:"h"`
	RealIp string `json:"rIp"`
	RealPort int `json:"rP"`
	Name string `json:"n"`
	UUID string `json:"u"`
	Properties []LoginPayloadProperty `json:"p"`
}

type LoginPayloadProperty struct {
	Name string `json:"n"`
	Value string `json:"v"`
	Signature string `json:"s"`
}

func EncodeLoginPayload(payload LoginPayload) (s string) {
	bytes, _ := json.Marshal(payload)
	s = string(bytes)
	return
}

func DecodeLoginPayload(s string) (payload LoginPayload) {
	json.Unmarshal([]byte(s), &payload)
	return
}
