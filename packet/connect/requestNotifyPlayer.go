package connect

import (
	"io"
	uuid "code.google.com/p/go-uuid/uuid"
	"github.com/suedadam/GoLilyPad/packet"
)

type RequestNotifyPlayer struct {
	Add bool
	Player string
	Uuid uuid.UUID
}

func NewRequestNotifyPlayerAdd(player string, uuid uuid.UUID) (this *RequestNotifyPlayer) {
	this = new(RequestNotifyPlayer)
	this.Add = true
	this.Player = player
	this.Uuid = uuid
	return
}

func NewRequestNotifyPlayerRemove(player string, uuid uuid.UUID) (this *RequestNotifyPlayer) {
	this = new(RequestNotifyPlayer)
	this.Add = false
	this.Player = player
	this.Uuid = uuid
	return
}

func (this *RequestNotifyPlayer) Id() int {
	return REQUEST_NOTIFY_PLAYER
}

type requestNotifyPlayerCodec struct {

}

func (this *requestNotifyPlayerCodec) Decode(reader io.Reader) (request Request, err error) {
	requestNotifyPlayer := new(RequestNotifyPlayer)
	requestNotifyPlayer.Add, err = packet.ReadBool(reader)
	if err != nil {
		return
	}
	requestNotifyPlayer.Player, err = packet.ReadString(reader)
	if err != nil {
		return
	}
	requestNotifyPlayer.Uuid, err = packet.ReadUUID(reader)
	if err != nil {
		return
	}
	request = requestNotifyPlayer
	return
}

func (this *requestNotifyPlayerCodec) Encode(writer io.Writer, request Request) (err error) {
	requestNotifyPlayer := request.(*RequestNotifyPlayer)
	err = packet.WriteBool(writer, requestNotifyPlayer.Add)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, requestNotifyPlayer.Player)
	if err != nil {
		return
	}
	err = packet.WriteUUID(writer, requestNotifyPlayer.Uuid)
	return
}

type ResultNotifyPlayer struct {

}

func NewResultNotifyPlayer() (this *ResultNotifyPlayer) {
	this = new(ResultNotifyPlayer)
	return
}

func (this *ResultNotifyPlayer) Id() int {
	return REQUEST_NOTIFY_PLAYER
}

type resultNotifyPlayerCodec struct {

}

func (this *resultNotifyPlayerCodec) Decode(reader io.Reader) (result Result, err error) {
	result = new(ResultNotifyPlayer)
	return
}

func (this *resultNotifyPlayerCodec) Encode(writer io.Writer, result Result) (err error) {
	return
}
