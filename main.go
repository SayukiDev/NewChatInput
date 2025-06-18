package main

import (
	app2 "ChatInput/app"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := app2.NewApp()
	ao := &options.App{
		Title:         "ChatInput",
		Width:         600,
		Height:        560,
		DisableResize: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
	}
	ao.Bind = app.Binds()
	// Create application with options
	err := wails.Run(ao)
	if err != nil {
		panic(err)
	}
}
