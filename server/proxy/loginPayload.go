package proxy

import (
	"encoding/json"
)

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

func EncodeLoginPayload(payload LoginPayload) (val string) {
	bytes, _ := json.Marshal(payload)
	val = string(bytes)
	return
}

func DecodeLoginPayload(val string) (payload LoginPayload) {
	json.Unmarshal([]byte(val), &payload)
	return
}
