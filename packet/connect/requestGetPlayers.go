package connect

import "io"
import "github.com/LilyPad/GoLilyPad/packet"

type RequestGetPlayers struct {
	List bool
}

func (this *RequestGetPlayers) Id() int {
	return REQUEST_GET_PLAYERS
}

type RequestGetPlayersCodec struct {

}

func (this *RequestGetPlayersCodec) Decode(reader io.Reader, util []byte) (request Request, err error) {
	requestGetPlayers := &RequestGetPlayers{}
	requestGetPlayers.List, err = packet.ReadBool(reader, util)
	if err != nil {
		return
	}
	return requestGetPlayers, nil
}

func (this *RequestGetPlayersCodec) Encode(writer io.Writer, util []byte, request Request) (err error) {
	err = packet.WriteBool(writer, util, request.(*RequestGetPlayers).List)
	return
}

type ResultGetPlayers struct {
	List bool
	CurrentPlayers uint16
	MaximumPlayers uint16
	Players []string
}

func (this *ResultGetPlayers) Id() int {
	return REQUEST_GET_PLAYERS
}

type ResultGetPlayersCodec struct {

}

func (this *ResultGetPlayersCodec) Decode(reader io.Reader, util []byte) (result Result, err error) {
	resultGetPlayers := &ResultGetPlayers{}
	resultGetPlayers.List, err = packet.ReadBool(reader, util)
	if err != nil {
		return
	}
	resultGetPlayers.CurrentPlayers, err = packet.ReadUint16(reader, util)
	if err != nil {
		return
	}
	resultGetPlayers.MaximumPlayers, err = packet.ReadUint16(reader, util)
	if err != nil {
		return
	}
	if resultGetPlayers.List {
		resultGetPlayers.Players = make([]string, resultGetPlayers.CurrentPlayers)
		var i uint16
		for i = 0; i < resultGetPlayers.CurrentPlayers; i++ {
			if err != nil {
				return
			}
			resultGetPlayers.Players[i], err = packet.ReadString(reader, util)
		}
	}
	return resultGetPlayers, nil
}

func (this *ResultGetPlayersCodec) Encode(writer io.Writer, util []byte, result Result) (err error) {
	resultGetPlayers := result.(*ResultGetPlayers)
	err = packet.WriteBool(writer, util, resultGetPlayers.List)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, util, resultGetPlayers.CurrentPlayers)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, util, resultGetPlayers.MaximumPlayers)
	if resultGetPlayers.List {
		var i uint16
		for i = 0; i < resultGetPlayers.CurrentPlayers; i++ {
			if err != nil {
				return
			}
			err = packet.WriteString(writer, util, resultGetPlayers.Players[i])
		}
	}
	return
}