package connect

import "github.com/LilyPad/GoLilyPad/packet/connect"

type RequestRecord struct {
	request connect.Request
	callback RequestCallback
}

type RequestCallback func(uint8, connect.Result)
