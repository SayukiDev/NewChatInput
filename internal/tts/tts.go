package tts

import (
	"ChatInput/pkg/aivis"
	malgo2 "ChatInput/pkg/malgo"
	"bufio"
	"errors"
	"os"
	"strings"
)

type TTS struct {
	o  *Option
	a  *aivis.Aivis
	ae *aivis.Engine
}

type Option struct {
	Baseurl    string
	Run        bool     `json:"run"`
	Log        string   `json:"log"`
	Path       string   `json:"path"`
	Args       []string `json:"args"`
	NowSpacker int64    `json:"now_spacker"`
	Device     string   `json:"device"`
	Cache      bool     `json:"cache"`
	CachePath  string   `json:"cache_path"`
}

func NewOption() Option {
	return Option{
		Baseurl: "http://localhost:10101",
		Run:     true,
		Log:     "./tts_engine.log",
		Path:    "./AivisSpeech/AivisSpeech-Engine/run.exe",
		Args: []string{
			"--host", "127.0.0.1",
			"--port", "10101",
			"--use_gpu",
		},
		NowSpacker: -1,
		Cache:      true,
		CachePath:  "./cache",
	}
}

func NewTTS(o *Option) (*TTS, error) {
	if o.Cache {
		os.MkdirAll(o.CachePath, 0755)
	}

	t := &TTS{
		o: o,
		a: aivis.New(o.Baseurl),
	}
	err := t.SetDevice(o.Device)
	if err != nil {
		return nil, err
	}
	if o.Run {
		ae, err := aivis.NewEngine(o.Log, o.Path, o.Args...)
		if err != nil {
			return nil, err
		}
		t.ae = ae
	}
	return t, nil
}

func (t *TTS) TTS(text string) (err error) {
	text = dictionaryTrans(text)
	var b []byte
	var isCached bool
	if t.o.Cache {
		b, err = t.readCache(text)
		if err != nil {
			return err
		}
		if b != nil {
			isCached = true
		}
	}
	if !isCached {
		if t.o.NowSpacker == -1 {
			ss, err := t.a.Speakers()
			if err != nil {
				return err
			}
			if len(ss) == 0 {
				return errors.New("no speakers")
			}
			if len(ss[0].Styles) == 0 {
				return errors.New("no styles")
			}
			t.o.NowSpacker = ss[0].Styles[0].ID
		}
		so, err := t.a.AudioQuery(text, t.o.NowSpacker)
		if err != nil {
			return err
		}
		b, err = t.a.Synthesis(text, t.o.NowSpacker, so)
		if err != nil {
			return err
		}
	}
	err = malgo2.PlayFromBytes(b)
	if err != nil {
		return err
	}
	if !isCached {
		err := t.writeCache(text, b)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *TTS) GetSpeakers() ([]aivis.Speaker, error) {
	return t.a.Speakers()
}

func (t *TTS) SetSpacker(s int64) {
	t.o.NowSpacker = s
}

func (t *TTS) SetDevice(id string) error {
	ds, err := malgo2.GetDevices()
	if err != nil {
		return err
	}

	for _, d := range ds {
		if d.ID.String() == id {
			malgo2.SetDevice(d.ID)
		}
	}
	return nil
}

func (t *TTS) Start() error {
	if t.ae != nil {
		err := t.ae.Start()
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *TTS) Close() error {
	if t.ae != nil {
		err := t.ae.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *TTS) IsStarted() (bool, error) {
	if running, err := t.ae.IsRunning(); err != nil {
		return false, err
	} else if !running {
		return false, nil
	}
	l, err := t.Logs(0)
	if err != nil {
		return false, nil
	}
	if l == "" {
		return false, nil
	}
	return strings.Contains(l, "Uvicorn running"), nil
}

func (t *TTS) Logs(limit int) (string, error) {
	f, err := os.Open(t.o.Log)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if len(lines) > limit && limit > 0 {
			lines = lines[1:]
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return strings.Join(lines, "\n"), nil
}

func (t *TTS) InstallModel(modelPath string) error {
	return t.ae.InstallModel(modelPath)
}
