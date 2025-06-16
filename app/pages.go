package app

import (
	"ChatInput/internal/pages"
	"ChatInput/internal/service"
)

type Pages struct {
	input *pages.Input
}

func NewPages() Pages {
	return Pages{
		input: pages.NewInput(),
	}
}

func (p *Pages) Binds() []interface{} {
	return []interface{}{
		p.input,
	}
}

func (p *Pages) Startup(srv *service.Service) {
	p.input.Startup(srv)
}
