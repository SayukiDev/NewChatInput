package pages

import (
	"ChatInput/internal/service"
	"ChatInput/options"
	"ChatInput/pkg/voicevox/api"
	"errors"
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

func (o *Options) GetSpacker() ([]api.Speaker, error) {
	if !o.srv.VV.Running() {
		return nil, errors.New("not running")
	}
	return o.srv.VV.ListSpeaker()
}

func (o *Options) IsVVRunning() bool {
	if o.srv.VV == nil {
		return false
	}
	return o.srv.VV.Running()
}

func (o *Options) IsVVComplete() bool {
	if o.srv.VV == nil {
		return false
	}
	return o.srv.VV.Complete()
}

func (o *Options) StartVV() error {
	if o.srv.VV == nil {
		return errors.New("voiceVox is not initialized")
	}
	if o.srv.VV.Running() {
		return nil
	}
	if err := o.srv.VV.Start(); err != nil {
		return err
	}
	return nil
}

func (o *Options) StopVV() error {
	if o.srv.VV == nil {
		return errors.New("voiceVox is not initialized")
	}
	if !o.srv.VV.Running() {
		return nil
	}
	if err := o.srv.VV.Close(); err != nil {
		return err
	}
	return nil
}
