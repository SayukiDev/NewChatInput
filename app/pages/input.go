package pages

import "ChatInput/internal/service"

type Input struct {
	srv *service.Service
}

func NewInput(srv *service.Service) *Input {
	return &Input{
		srv: srv,
	}
}

func (i *Input) SendMessage(message string) error {
	return i.srv.SendChatboxMsg(message, i.srv.Option.TTS, false)
}
