package pages

import "ChatInput/internal/service"

type Register struct {
	app     *App
	input   *Input
	tts     *TTS
	options *Options
}

func NewRegister() Register {
	return Register{
		app:     NewAPP(),
		input:   NewInput(),
		options: NewOptions(),
		tts:     NewTTS(),
	}
}

func (p *Register) Binds() []interface{} {
	return []interface{}{
		p.app,
		p.input,
		p.options,
		p.tts,
	}
}

type content struct {
	srv *service.Service
	reg *Register
}

func (p *Register) Startup(srv *service.Service) {
	c := &content{
		srv: srv,
		reg: p,
	}
	p.app.Startup(c)
	p.input.Startup(c)
	p.options.Startup(c)
	p.tts.Startup(c)
}
