package options

import (
	"ChatInput/internal/tts"
	"encoding/json"
	"os"
	"sync"
)

type Options struct {
	// internal
	path    string
	lock    sync.Mutex
	hooks   []HookFunc
	*Config `json:",inline"`
}

type Config struct {
	SendPort        int        `json:"send_port"`
	RecvPort        int        `json:"recv_port"`
	EnableTypingMsg bool       `json:"enable_typing_msg"`
	RealtimeSend    bool       `json:"realtime"`
	MsgKeeping      bool       `json:"msg_keeping"`
	VoiceControl    bool       `json:"voice_control"`
	TTSOption       tts.Option `json:"tts_option"`
}

type HookFunc func(o *Options) error

func NewOptions(p string) *Options {
	return &Options{
		path: p,
		Config: &Config{
			SendPort:     9000,
			RecvPort:     9001,
			RealtimeSend: false,
			VoiceControl: false,
			TTSOption:    tts.NewOption(),
		},
	}
}

func (o *Options) Load() error {
	o.lock.Lock()
	defer o.lock.Unlock()
	file, err := os.Open(o.path)
	if err != nil {
		if os.IsNotExist(err) {
			o.lock.Unlock()
			defer o.lock.Lock()
			return o.Save()
		}
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	return decoder.Decode(o.Config)
}

func (o *Options) Save() error {
	o.lock.Lock()
	defer o.lock.Unlock()
	file, err := os.OpenFile(o.path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	return encoder.Encode(o.Config)
}

func (o *Options) AddHook(f HookFunc) {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.hooks = append(o.hooks, f)
	return
}

func (o *Options) Updated() error {
	o.lock.Lock()
	defer o.lock.Unlock()
	for _, hook := range o.hooks {
		if err := hook(o); err != nil {
			return err
		}
	}
	return nil
}
