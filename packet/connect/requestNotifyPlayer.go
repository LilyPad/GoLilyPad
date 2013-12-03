package connect

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type RequestNotifyPlayer struct {
	Add bool
	Player string
}

func (this *RequestNotifyPlayer) Id() int {
	return REQUEST_NOTIFY_PLAYER
}

type RequestNotifyPlayerCodec struct {
	
}

func (this *RequestNotifyPlayerCodec) Decode(reader io.Reader, util []byte) (request Request, err error) {
	requestNotifyPlayer := &RequestNotifyPlayer{}
	requestNotifyPlayer.Add, err = packet.ReadBool(reader, util)
	if err != nil {
		return
	}
	requestNotifyPlayer.Player, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	return requestNotifyPlayer, nil
}

func (this *RequestNotifyPlayerCodec) Encode(writer io.Writer, util []byte, request Request) (err error) {
	requestNotifyPlayer := request.(*RequestNotifyPlayer)
	err = packet.WriteBool(writer, util, requestNotifyPlayer.Add)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, util, requestNotifyPlayer.Player)
	return
}

type ResultNotifyPlayer struct {

}

func (this *ResultNotifyPlayer) Id() int {
	return REQUEST_NOTIFY_PLAYER
}

type ResultNotifyPlayerCodec struct {

}

func (this *ResultNotifyPlayerCodec) Decode(reader io.Reader, util []byte) (result Result, err error) {
	return &ResultNotifyPlayer{}, nil
}

func (this *ResultNotifyPlayerCodec) Encode(writer io.Writer, util []byte, result Result) (err error) {
	return
}