package main

import (
	app2 "ChatInput/app"
	"ChatInput/internal/service"
	o "ChatInput/options"
	"embed"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
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
	app := app2.NewApp(srv)
	ao := &options.App{
		Title:  "ChatInput",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
	}
	ao.Bind = app.Binds()
	// Create application with options
	err = wails.Run(ao)
	if err != nil {
		panic(err)
	}
}
