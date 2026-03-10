package malgo

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/gen2brain/malgo"
	"github.com/hajimehoshi/go-mp3"
	"github.com/youpy/go-wav"
)

var selectedDevice atomic.Pointer[malgo.DeviceID]

func GetDevices() ([]malgo.DeviceInfo, error) {
	ctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = ctx.Uninit()
		ctx.Free()
	}()

	infos, err := ctx.Devices(malgo.Playback)
	if err != nil {
		return nil, err
	}
	return infos, nil
}

func SetDevice(name malgo.DeviceID) {
	selectedDevice.Store(&name)
	// Reset the global player so it's recreated with the new device on next play.
	globalPlayerMu.Lock()
	if globalPlayer != nil {
		globalPlayer.close()
		globalPlayer = nil
	}
	globalPlayerMu.Unlock()
}

type ReaderAt interface {
	io.ReaderAt
	io.Reader
}

// audioFormat identifies the playback format of a device.
type audioFormat struct {
	channels   uint32
	sampleRate uint32
}

// audioState holds the audio source for the current playback.
type audioState struct {
	reader io.Reader
	done   chan struct{}
}

// player manages a persistent audio playback device.
// The device runs continuously; silence is emitted when nothing is queued.
type player struct {
	ctx    *malgo.AllocatedContext
	device *malgo.Device
	format audioFormat
	state  atomic.Pointer[audioState]
	closed chan struct{}
}

func newPlayer(format audioFormat) (*player, error) {
	ctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, func(_ string) {})
	if err != nil {
		return nil, err
	}

	p := &player{
		ctx:    ctx,
		format: format,
		closed: make(chan struct{}),
	}

	deviceConfig := malgo.DefaultDeviceConfig(malgo.Playback)
	if selectedDevice.Load() != nil {
		deviceConfig.Playback.DeviceID = selectedDevice.Load().Pointer()
	}
	deviceConfig.Playback.Format = malgo.FormatS16
	deviceConfig.Playback.Channels = format.channels
	deviceConfig.SampleRate = format.sampleRate
	deviceConfig.Alsa.NoMMap = 1

	device, err := malgo.InitDevice(ctx.Context, deviceConfig, malgo.DeviceCallbacks{
		Data: p.onSamples,
	})
	if err != nil {
		_ = ctx.Uninit()
		ctx.Free()
		return nil, err
	}
	p.device = device

	if err := device.Start(); err != nil {
		device.Uninit()
		_ = ctx.Uninit()
		ctx.Free()
		return nil, err
	}

	return p, nil
}

// onSamples is called by the audio thread to fill the output buffer.
// It reads from the current audioState, or emits silence when idle.
func (p *player) onSamples(output, _ []byte, _ uint32) {
	s := p.state.Load()
	if s == nil {
		clear(output)
		return
	}

	n, err := io.ReadFull(s.reader, output)
	if err != nil {
		// Zero out the unfilled tail of the buffer.
		clear(output[n:])
		p.state.Store(nil)
		close(s.done)
	}
}

// play queues audio data and blocks until playback completes.
func (p *player) play(reader io.Reader) {
	s := &audioState{
		reader: reader,
		done:   make(chan struct{}),
	}
	p.state.Store(s)
	select {
	case <-s.done:
	case <-p.closed:
	}
}

func (p *player) close() {
	close(p.closed)
	p.device.Uninit()
	_ = p.ctx.Uninit()
	p.ctx.Free()
}

// Global persistent player.
var (
	globalPlayer   *player
	globalPlayerMu sync.Mutex
)

func getPlayer(format audioFormat) (*player, error) {
	globalPlayerMu.Lock()
	defer globalPlayerMu.Unlock()

	if globalPlayer != nil {
		if globalPlayer.format == format {
			return globalPlayer, nil
		}
		// Format changed: recreate the device.
		globalPlayer.close()
		globalPlayer = nil
	}

	p, err := newPlayer(format)
	if err != nil {
		return nil, err
	}
	globalPlayer = p
	return p, nil
}

// Shutdown cleanly stops the persistent player. Call on application exit.
func Shutdown() {
	globalPlayerMu.Lock()
	defer globalPlayerMu.Unlock()
	if globalPlayer != nil {
		globalPlayer.close()
		globalPlayer = nil
	}
}

func PlayFromFile(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	return Play(f, strings.TrimPrefix(strings.ToLower(path.Ext(file)), "."))
}

func PlayFromBytes(data []byte) error {
	r := bytes.NewReader(data)
	return Play(r, "wav")
}

func Play(audio ReaderAt, format string) error {
	var reader io.Reader
	var channels, sampleRate uint32

	switch format {
	case "wav":
		w := wav.NewReader(audio)
		f, err := w.Format()
		if err != nil {
			return err
		}
		reader = w
		channels = uint32(f.NumChannels)
		sampleRate = f.SampleRate

	case "mp3":
		m, err := mp3.NewDecoder(audio)
		if err != nil {
			return err
		}
		reader = m
		channels = 2
		sampleRate = uint32(m.SampleRate())

	default:
		return fmt.Errorf("unsupported format: %s", format)
	}

	p, err := getPlayer(audioFormat{channels: channels, sampleRate: sampleRate})
	if err != nil {
		return err
	}

	p.play(reader)
	return nil
}
