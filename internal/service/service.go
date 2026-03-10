package service

import (
	"ChatInput/options"
	"ChatInput/pkg/tasks"
	"context"

	"github.com/SayukiDev/VRCOSC"
)

type Service struct {
	OSC    *VRCOSC.VRCOsc
	Option *options.Options
	Tasks  *tasks.Tasks
	AppCtx context.Context

	// Sub Services
	ChatBox *ChatBox

	// Sub Services Control
	servicesControl []subService
}

type subService interface {
	Start() error
	Close() error
}

func New(opt *options.Options) (*Service, error) {
	s := &Service{
		Option: opt,
		Tasks:  tasks.New(),
	}
	c, err := newChatBox(opt, s)
	if err != nil {
		return nil, err
	}
	s.ChatBox = c
	s.servicesControl = append(s.servicesControl, c)
	return s, nil
}

func (s *Service) Start(ctx context.Context) error {
	s.AppCtx = ctx
	for _, c := range s.servicesControl {
		err := c.Start()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) Close() error {
	for _, c := range s.servicesControl {
		_ = c.Close()
	}
	return nil
}
