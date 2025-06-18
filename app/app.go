package app

import (
	"ChatInput/internal/service"
	o "ChatInput/options"
	"context"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
)

// App struct
type App struct {
	ctx   context.Context
	srv   *service.Service
	pages Pages
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		pages: NewPages(),
	}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	var err error
	defer func() {
		if err != nil {
			logrus.WithError(err).Error("Failed to start app")
			runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
				Title:   "Error",
				Message: "Failed to start app: " + err.Error(),
				Type:    runtime.ErrorDialog,
			})
			os.Exit(1)
		} else {
			logrus.Println("App started successfully")
		}
	}()
	opt := o.NewOptions("./data.json")
	err = opt.Load()
	if err != nil {
		return
	}
	srv := service.New(opt)
	defer func() {
		logrus.Println("Shutdown...")
		err = srv.Close()
		if err != nil {
			panic(err)
		}
		logrus.Println("Done.")
	}()
	err = srv.Start(ctx)
	if err != nil {
		return
	}
	a.ctx = ctx
	a.srv = srv
	a.pages.Startup(srv)
}

func (a *App) Binds() []interface{} {
	return a.pages.Binds()
}
