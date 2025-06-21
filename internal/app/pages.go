package app

import (
	"ChatInput/internal/pages"
	"ChatInput/internal/service"
)

type Pages struct {
	input   *pages.Input
	options *pages.Options
}

func NewPages() Pages {
	return Pages{
		input:   pages.NewInput(),
		options: pages.NewOptions(),
	}
}

func (p *Pages) Binds() []interface{} {
	return []interface{}{
		p.input,
		p.options,
	}
}

func (p *Pages) Startup(srv *service.Service) {
	p.input.Startup(srv)
	p.options.Startup(srv)
}
