package service

import (
	"ChatInput/options"
	"github.com/SayukiDev/VRCOSC"
)

func (s *Service) initOsc(opt *options.Options) {
	s.OSC = VRCOSC.New(&VRCOSC.Options{
		SendPort: opt.SendPort,
		RecvPort: opt.RecvPort,
	})
	opt.AddHook(func(o *options.Options) error {
		s.OSC = VRCOSC.New(&VRCOSC.Options{
			SendPort: o.SendPort,
			RecvPort: o.RecvPort,
		})
		if o.MsgKeeping {
			if !s.Tasks.Exist("osc_msg_keeping") {
				err := s.Tasks.Add("osc_msg_keeping", s.keepingMsgTask)
				if err != nil {
					return err
				}
			}
		} else {
			if s.Tasks.Exist("osc_msg_keeping") {
				err := s.Tasks.Remove("osc_msg_keeping")
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}
