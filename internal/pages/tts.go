package pages

import (
	"ChatInput/internal/tts"
	"ChatInput/pkg/aivis"
)

type TTS struct {
	*content
	Startup func(c *content)
}

func NewTTS() *TTS {
	o := &TTS{}
	o.Startup = o.startup
	return o
}

func (o *TTS) startup(srv *content) {
	o.content = srv
}

func (o *TTS) GetSpackers() ([]aivis.Speaker, error) {
	return o.srv.ChatBox.TTS.GetSpeakers()
}

func (o *TTS) SaveSpacker(s int64) (err error) {
	o.srv.Option.TTSOption.NowSpacker = s
	err = o.srv.Option.Save()
	if err != nil {
		return err
	}
	return nil
}

type SelectedSpackerRsp struct {
	SpackerId int64
}

func (o *TTS) SelectedSpacker() SelectedSpackerRsp {
	return SelectedSpackerRsp{
		SpackerId: o.srv.Option.TTSOption.NowSpacker,
	}
}

func (o *TTS) GetAudioDevice() ([]tts.Device, error) {
	return o.srv.ChatBox.TTS.GetDevices()
}

func (o *TTS) SaveAudioDevice(id string) error {
	err := o.srv.ChatBox.TTS.SetDevice(id)
	if err != nil {
		return err
	}
	return o.srv.Option.Save()
}

func (o *TTS) ReadLog() (string, error) {
	if !o.srv.Option.TTS {
		return "not running", nil
	}
	if r, _ := o.srv.ChatBox.IsStarted(); !r {
		return "not running", nil
	}
	return o.srv.ChatBox.Logs(100)
}

func (o *TTS) IsRunning() bool {
	r, _ := o.srv.ChatBox.TTS.IsStarted()
	return r
}
