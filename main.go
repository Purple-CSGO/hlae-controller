package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"hlae-controller/backend"

	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// App instance
	app := backend.NewApp(appName, appDeveloper, version)

	// Create and run
	err := wails.Run(&options.App{
		Title:             appName,
		Width:             Width,
		Height:            Height,
		MinWidth:          MinWidth,
		MinHeight:         MinHeight,
		MaxWidth:          MaxWidth,
		MaxHeight:         MaxHeight,
		DisableResize:     DisableResize,
		Fullscreen:        Fullscreen,
		Frameless:         Frameless,
		StartHidden:       StartHidden,
		HideWindowOnClose: HideWindowOnClose,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 0},
		AssetServer:       &assetserver.Options{Assets: assets},
		LogLevel:          logger.DEBUG,
		OnStartup:         app.Startup,
		OnDomReady:        app.DomReady,
		OnBeforeClose:     app.BeforeClose,
		OnShutdown:        app.Shutdown,
		Bind: []interface{}{
			app,
		},

		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               true,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
		},

		// Mac platform specific options
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarHiddenInset(),
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "HLAE Controller",
				Message: "Externel Control App for HLAE",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		println(err)
	}
}
