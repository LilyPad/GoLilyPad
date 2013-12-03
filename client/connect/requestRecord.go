package connect

import "github.com/LilyPad/GoLilyPad/packet/connect"

type RequestRecord struct {
	request connect.Request
	callback RequestCallback
}

type RequestCallback func(statusCode uint8, result connect.Result)
