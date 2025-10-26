package service

import (
	"fmt"
	"github.com/SayukiDev/Beep"
	"time"
)

func (s *Service) SendChatboxMsg(text string, tts bool, disableSfx bool) error {
	err := s.OSC.ChatBoxInput(text, true, !disableSfx)
	if err != nil {
		return err
	}
	if s.Option.TTS && text != "" && tts {
		if !s.VV.Complete() {
			return nil // voicevox not ready
		}
		b, err := s.VV.TTS(text)
		if err != nil {
			return err
		}

		// voice toggle, but working is so bad, so disabled for now.
		/*err = u.OSC.SendRaw(OSC.NewMessage("/input/Voice", int32(0)))
		if err != nil {
			return err
		}*/
		err = Beep.Play(b, "wav", func() {
			/*err = u.OSC.SendRaw(OSC.NewMessage("/input/Voice", int32(1)))
			if err != nil {
				log.Println("Warn: [ Send voice toggle error:", err, "]")
			}*/
		})
		if err != nil {
			return err
		}
	}
	if s.Option.MsgKeeping {
		s.ChatBoxKeepingMsg.Store(text)
	} else {
		if s.ChatBoxKeepingMsg.Load() != "" {
			s.ChatBoxKeepingMsg.Store("")
		}
	}
	return nil
}

func (s *Service) keepingMsgTask(c chan struct{}) {
	for range time.Tick(time.Second * 3) {
		select {
		case <-c:
			return
		default:
		}
		if s.ChatBoxKeepingMsg.Load() != "" {
			err := s.SendChatboxMsg(s.ChatBoxKeepingMsg.Load(), false, true)
			if err != nil {
				fmt.Println("Warn: [ Chatbox keeping msg error:", err, "]")
			}
		}
	}
	return
}
