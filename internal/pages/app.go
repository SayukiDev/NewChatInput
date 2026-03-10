package pages

import (
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	*content
	Startup   func(c *content)
	sizeRadio float64
	exited    bool
}

func NewAPP() *App {
	i := &App{
		sizeRadio: 1.4,
	}
	i.Startup = i.startup
	return i
}

func (i *App) IsLoaded() bool {
	if !i.srv.Option.TTS {
		return true
	}
	r, err := i.srv.ChatBox.IsStarted()
	if err != nil {
		if i.exited {
			return true
		}
		i.exited = true
		runtime.MessageDialog(i.srv.AppCtx, runtime.MessageDialogOptions{
			Title:   "Error",
			Message: err.Error(),
			Type:    runtime.ErrorDialog,
		})
		os.Exit(1)
		return true
	}
	return r
}

func (i *App) GetSizeRatio() float64 {
	return i.sizeRadio
}

func (i *App) SetSizeRatio(ratio float64) {
	i.sizeRadio = ratio
}

func (i *App) startup(srv *content) {
	i.content = srv
}
