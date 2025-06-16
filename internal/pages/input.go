package pages

import "ChatInput/internal/service"

type Input struct {
	srv     *service.Service
	Startup func(srv *service.Service)
}

func NewInput() *Input {
	i := &Input{}
	i.Startup = i.startup
	return i
}

func (i *Input) startup(srv *service.Service) {
	i.srv = srv
}

func (i *Input) SendMessage(message string) error {
	return i.srv.SendChatboxMsg(message, i.srv.Option.TTS, false)
}
