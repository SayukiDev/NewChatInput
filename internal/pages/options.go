package pages

import (
	"ChatInput/options"
)

type Options struct {
	*content
	Startup func(c *content)
}

func NewOptions() *Options {
	o := &Options{}
	o.Startup = o.startup
	return o
}

func (o *Options) startup(srv *content) {
	o.content = srv
}

func (o *Options) Load() *options.Config {
	return o.srv.Option.Config
}

func (o *Options) Save(opt *options.Config) error {
	o.srv.Option.Config = opt
	if err := o.srv.Option.Save(); err != nil {
		return err
	}
	err := o.srv.Option.Updated()
	if err != nil {
		return err
	}
	return nil
}
