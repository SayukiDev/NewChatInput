package service

import (
	"ChatInput/options"
	"ChatInput/pkg/voicevox"
	"fmt"
)

func (s *Service) initVoiceVox(opt *options.Options) error {
	s.VV = voicevox.New(opt.VoiceVox.Path, opt.VoiceVox.LineLimit, &opt.VoiceVox)
	s.VV.SetSpeaker(opt.VoiceVox.Selected)
	opt.AddHook(func(o *options.Options) error {
		if o.VoiceVox.IsRemote {
			s.VV.Close()
			s.VV = voicevox.New(opt.VoiceVox.Path, opt.VoiceVox.LineLimit, &opt.VoiceVox)
		}
		s.VV.SetSpeaker(opt.VoiceVox.Selected)
		return nil
	})
	if opt.VoiceVox.AutoStart {
		if err := s.VV.Start(); err != nil {
			return fmt.Errorf("failed to start VoiceVox: %w", err)
		}
	}
	return nil
}
