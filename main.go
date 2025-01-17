package main

import (
	"embed"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"launcher/pkg/filehandler"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

// http://localhost:34115/
func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "launcher",
		DisableResize: true,
		Frameless:     true,
		AlwaysOnTop:   true,
		MaxWidth:      Width,
		MaxHeight:     Height,
		MinWidth:      Width,
		MinHeight:     Height,
		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: filehandler.Handler(),
		},
		Linux: &linux.Options{
			WindowIsTranslucent: true,
			ProgramName:         "launcher",
		},
		//BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:  app.startup,
		OnDomReady: app.domReady,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
