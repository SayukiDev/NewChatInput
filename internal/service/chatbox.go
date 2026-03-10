package service

import (
	"ChatInput/internal/tts"
	"ChatInput/options"
	"fmt"
	"time"

	"github.com/SayukiDev/VRCOSC"
	"go.uber.org/atomic"
)

type ChatBox struct {
	keepingMsg atomic.String
	TTS        *tts.TTS
	osc        *VRCOSC.VRCOsc
	p          *Service
}

func newChatBox(opt *options.Options, s *Service) (*ChatBox, error) {
	cs := &ChatBox{
		p: s,
	}
	err := cs.initTTS(opt)
	if err != nil {
		return nil, err
	}
	err = cs.initOsc(opt)
	if err != nil {
		return nil, err
	}
	return cs, nil
}

func (s *ChatBox) initTTS(opt *options.Options) error {
	if !opt.TTS {
		return nil
	}
	if s.TTS != nil {
		err := s.TTS.Close()
		if err != nil {
			return err
		}
	}
	t, err := tts.NewTTS(&opt.TTSOption)
	if err != nil {
		return err
	}
	s.TTS = t
	return nil
}

func (s *ChatBox) initOsc(opt *options.Options) error {
	s.osc = VRCOSC.New(&VRCOSC.Options{
		SendPort: opt.SendPort,
		RecvPort: opt.RecvPort,
	})
	keepingInit := func(o *options.Options) error {
		s.osc = VRCOSC.New(&VRCOSC.Options{
			SendPort: o.SendPort,
			RecvPort: o.RecvPort,
		})
		if o.MsgKeeping || o.EnableTypingMsg {
			if !s.p.Tasks.Exist("osc_msg_keeping") {
				err := s.p.Tasks.Add("osc_msg_keeping", s.keepingMsgTask)
				if err != nil {
					return err
				}
			}
		} else {
			if s.p.Tasks.Exist("osc_msg_keeping") {
				err := s.p.Tasks.Remove("osc_msg_keeping")
				if err != nil {
					return err
				}
			}
		}
		return nil
	}
	err := keepingInit(opt)
	if err != nil {
		return err
	}
	opt.AddHook(keepingInit)
	return nil
}

func (s *ChatBox) SendChatboxMsg(text string, tts bool, disableSfx bool) error {
	err := s.osc.ChatBoxInput(text, true, !disableSfx)
	if err != nil {
		return err
	}
	if s.p.Option.TTS && text != "" && tts {
		err := s.TTS.TTS(text)
		if err != nil {
			return err
		}
	}
	if s.p.Option.MsgKeeping {
		s.keepingMsg.Store(text)
	} else {
		if s.keepingMsg.Load() != "" {
			s.keepingMsg.Store("")
		}
	}
	return nil
}

func (s *ChatBox) SetTyping(typing bool) {
	if typing {
		s.keepingMsg.Store("入力中...")
	} else {
		s.keepingMsg.Store("")
	}
}

func (s *ChatBox) keepingMsgTask(c chan struct{}) {
	for range time.Tick(time.Second * 3) {
		select {
		case <-c:
			return
		default:
		}
		if s.keepingMsg.Load() != "" {
			err := s.SendChatboxMsg(s.keepingMsg.Load(), false, true)
			if err != nil {
				fmt.Println("Warn: [ Chatbox keeping msg error:", err, "]")
			}
		}
	}
	return
}

func (s *ChatBox) IsStarted() (bool, error) {
	r, err := s.TTS.IsStarted()
	if err != nil {
		return false, err
	}
	return r, nil
}

func (s *ChatBox) Logs(limit int) (string, error) {
	return s.TTS.Logs(limit)
}

func (s *ChatBox) Start() error {
	if s.TTS == nil {
		return nil
	}
	err := s.TTS.Start()
	if err != nil {
		return err
	}
	return nil
}

func (s *ChatBox) Close() error {
	if s.TTS == nil {
		return nil
	}
	err := s.TTS.Close()
	if err != nil {
		return err
	}
	return nil
}
