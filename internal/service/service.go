package service

import (
	"ChatInput/options"
	"ChatInput/pkg/tasks"
	"ChatInput/pkg/voicevox"
	"context"
	"github.com/SayukiDev/VRCOSC"
)

type Service struct {
	OSC    *VRCOSC.VRCOsc
	VV     *voicevox.VoiceVox
	Option *options.Options
	Tasks  *tasks.Tasks
	AppCtx context.Context
}

func New(opt *options.Options) *Service {
	s := &Service{
		Option: opt,
		Tasks:  tasks.New(),
	}
	return s
}

func (s *Service) Start(ctx context.Context) error {
	s.initOsc(s.Option)
	err := s.initVoiceVox(s.Option)
	if err != nil {
		return err
	}
	s.AppCtx = ctx
	return nil
}

func (s *Service) Close() error {
	return s.VV.Close()
}
