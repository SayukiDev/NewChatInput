package app

import (
	"ChatInput/app/pages"
	"ChatInput/internal/service"
	"context"
)

// App struct
type App struct {
	ctx   context.Context
	srv   *service.Service
	input *pages.Input
}

// NewApp creates a new App application struct
func NewApp(srv *service.Service) *App {
	return &App{
		srv:   srv,
		input: pages.NewInput(srv),
	}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Binds() []interface{} {
	return []interface{}{
		a.input,
	}
}
