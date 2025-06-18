package pages

import (
	"ChatInput/internal/service"
	"ChatInput/options"
)

type Options struct {
	srv     *service.Service
	Startup func(srv *service.Service)
}

func NewOptions() *Options {
	o := &Options{}
	o.Startup = o.startup
	return o
}

func (o *Options) startup(srv *service.Service) {
	o.srv = srv
}

func (o *Options) Load() *options.Options {
	return o.srv.Option
}

func (o *Options) Save(opt *options.Options) error {
	if err := opt.Save(); err != nil {
		return err
	}
	o.srv.Option = opt
	err := o.srv.Option.Updated()
	if err != nil {
		return err
	}
	return nil
}
