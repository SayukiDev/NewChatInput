package app

import (
	"ChatInput/internal/pages"
	"ChatInput/internal/service"
	o "ChatInput/options"
	malgo2 "ChatInput/pkg/malgo"
	"context"
	"os"

	"github.com/gentlemanautomaton/winobj/winmutex"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx   context.Context
	srv   *service.Service
	pages pages.Register
	m     *winmutex.Mutex
}

// NewApp creates a new App application struct
func NewApp() (*App, error) {
	m, err := winmutex.New("SayukiDev/ChatInput")
	if err != nil {
		return nil, err
	}
	if !m.TryLock() {
		logrus.Fatal("Another instance is already running")
	}
	return &App{
		pages: pages.NewRegister(),
		m:     m,
	}, nil
}

func (a *App) Panic(err error, title, message string) {
	logrus.WithError(err).Error(message)
	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Title:   title,
		Message: message + "\n\n" + err.Error(),
		Type:    runtime.ErrorDialog,
	})
	os.Exit(1)
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	var err error
	defer func() {
		if err != nil {
			a.Panic(err, "Error", "Failed to start app")
		} else {
			logrus.Println("App started successfully")
		}
	}()
	opt := o.NewOptions("./data.json")
	err = opt.Load()
	if err != nil {
		return
	}
	srv, err := service.New(opt)
	if err != nil {
		return
	}
	go func() {
		runtime.EventsEmit(ctx, "serviceLoad")
		err = srv.Start(ctx)
		if err != nil {
			a.Panic(err, "Error", "Failed to start service")
			return
		}
		runtime.EventsEmit(ctx, "serviceReady")
	}()
	a.ctx = ctx
	a.srv = srv
	a.pages.Startup(srv)
}

func (a *App) Shutdown(_ context.Context) {
	if a.srv != nil {
		if err := a.srv.Close(); err != nil {
			logrus.WithError(err).Error("Failed to close service")
		}
	}
	malgo2.Shutdown()
	a.m.Unlock()
	a.m.Close()
	logrus.Println("App shutdown successfully")
}

func (a *App) Binds() []interface{} {
	return a.pages.Binds()
}
