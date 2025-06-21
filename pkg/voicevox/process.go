package voicevox

import (
	"bufio"
	"errors"
	log "github.com/sirupsen/logrus"
	"io"
	"strings"
	"time"
)

func (v *VoiceVox) runExec() {
	if err := v.process.Start(); err != nil {
		log.WithError(err).Error("Failed to start VoiceVox process")
		close(v.closeBack)
		v.Close()
		return
	}
	d := make(chan struct{})
	c := make(chan struct{})
	go v.readToLogLoop(d, c)
	select {
	case <-d:
	case <-v.closed:
		v.process.Wait()
		close(v.closeBack)
		return
	case <-time.After(time.Minute * 2):
		log.Error("VoiceVox process timeout")
		close(v.closeBack)
		v.Close()
		return
	}
	select {
	case <-v.closed:
		v.process.Wait()
		close(v.closeBack)
		close(c)
	}
}

func (v *VoiceVox) readToLogLoop(done, closeR chan struct{}) {
	br := bufio.NewReader(v.reader)
	for {
		select {
		case <-closeR:
			return
		default:
		}
		line, err := br.ReadString('\n')
		if errors.Is(err, io.ErrClosedPipe) {
			return
		}
		if err != nil {
			log.WithError(err).Error("Read log failed")
		}
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "running") {
			if !v.complete.Load() {
				v.complete.Store(true)
				close(done)
			}
		}
	}
}
