package connect

import (
	"github.com/suedadam/GoLilyPad/packet/connect"
)

type RequestRecord struct {
	request connect.Request
	callback RequestCallback
}

func NewRequestRecord(request connect.Request, callback RequestCallback) (this *RequestRecord) {
	this = new(RequestRecord)
	this.request = request
	this.callback = callback
	return
}

type RequestCallback func(statusCode uint8, result connect.Result)
