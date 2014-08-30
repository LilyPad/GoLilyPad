package connect

import (
	"github.com/LilyPad/GoLilyPad/packet/connect"
)

type Connect interface {
	Connect(addr string) (err error)
	Disconnect()
	Connected() bool

	Request(request connect.Request) (statusCode uint8, result connect.Result, err error)
	RequestLater(request connect.Request, callback RequestCallback) (err error)
	DispatchResult(sequenceId int32, statusCode uint8, result connect.Result) (err error)
	RequestIdBySequenceId(sequenceId int32) int

	RegisterEvent(name string, eventHandler EventHandler)
	DispatchEvent(name string, event Event)
}
