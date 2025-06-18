package pages

import (
	"ChatInput/internal/service"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

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

func (i *Input) SetFullInputMode(mode bool) {
	if mode {
		w, h := runtime.WindowGetSize(i.srv.AppCtx)
		runtime.WindowSetSize(i.srv.AppCtx, w, h-100)
	} else {
		w, h := runtime.WindowGetSize(i.srv.AppCtx)
		runtime.WindowSetSize(i.srv.AppCtx, w, h+100)
	}
}
