package app

import (
	"ChatInput/internal/service"
	o "ChatInput/options"
	"context"
	"github.com/sirupsen/logrus"
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
	opt := o.NewOptions("./data.json")
	err := opt.Load()
	if err != nil {
		panic(err)
	}
	srv, err := service.New(opt)
	if err != nil {
		panic(err)
	}
	defer func() {
		logrus.Println("Shutdown...")
		err = srv.Close()
		if err != nil {
			panic(err)
		}
		logrus.Println("Done.")
	}()
	a.ctx = ctx
	a.srv = srv
	a.pages.Startup(srv)
}

func (a *App) Binds() []interface{} {
	return a.pages.Binds()
}
