package proxy

import "github.com/LilyPad/GoLilyPad/packet/minecraft"

type apiOutBridge struct {
	session *SessionOutBridge
}

func (this *apiOutBridge) Version() *minecraft.Version {
	return this.session.protocol
}


