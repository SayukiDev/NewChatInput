package voicevox

import (
	"ChatInput/options"
	"ChatInput/pkg/cmd"
	"ChatInput/pkg/portkill"
	"ChatInput/pkg/voicevox/api"
	"io"
	"os/exec"
	"strings"
	"sync/atomic"
)

type VoiceVox struct {
	*api.Api
	isRemote  bool
	host      string
	port      string
	complete  atomic.Bool
	running   atomic.Bool
	closed    chan struct{}
	closeBack chan struct{}
	reader    *io.PipeReader
	writer    *io.PipeWriter
	cmd       []string
	process   *exec.Cmd
}

func New(path string, lineLimit int, c *options.VoiceVox) *VoiceVox {
	if c.IsRemote {
		return &VoiceVox{
			Api:      api.New("http://" + c.Host + ":" + c.Port + "/"),
			isRemote: true,
		}
	}
	newO := make([]string, 0, len(c.Args))
	skipNext := false
	for _, v := range c.Args {
		if skipNext {
			skipNext = false
			continue
		}
		if strings.HasPrefix(strings.ToLower(v), "--host") {
			skipNext = true
			continue
		}
		if strings.HasPrefix(strings.ToLower(v), "--port") {
			skipNext = true
			continue
		}
		newO = append(newO, v)
	}
	args := []string{
		path,
		"--host", c.Host,
		"--port", c.Port,
	}
	args = append(args, newO...)
	l := make([]string, 0, lineLimit)
	l = append(l, strings.Join(args, ""))
	return &VoiceVox{
		Api: api.New("http://" + c.Host + ":" + c.Port + "/"),
		cmd: args,
	}
}

func (v *VoiceVox) Running() bool {
	return v.running.Load()
}

func (v *VoiceVox) Complete() bool {
	return v.complete.Load()
}

func (v *VoiceVox) Start() error {
	if v.running.Load() {
		return nil
	}
	if v.isRemote {
		return nil
	}
	v.running.Store(true)
	if portkill.IsPortOpen(v.port) {
		portkill.KillPort(v.port)
	}
	v.closed = make(chan struct{})
	v.closeBack = make(chan struct{})
	v.reader, v.writer = io.Pipe()
	v.process = exec.Command(v.cmd[0], v.cmd[1:]...)
	v.process.Stdout = v.writer
	v.process.Stderr = v.writer
	v.process.SysProcAttr = cmd.HideWindowAttr()
	go v.runExec()
	return nil
}

func (v *VoiceVox) Close() error {
	if !v.running.Load() {
		return nil
	}
	v.running.Store(false)
	v.complete.Store(false)
	if v.process != nil && v.process.Process != nil {
		_ = v.process.Process.Kill()
	}
	close(v.closed)
	<-v.closeBack
	err := v.reader.Close()
	if err != nil {
		return err
	}
	err = v.writer.Close()
	if err != nil {
		return err
	}
	return nil
}
