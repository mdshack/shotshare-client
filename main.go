package main

import (
	"context"
	"embed"
	"fmt"

	"github.com/kbinani/screenshot"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:      "shotshare-client-sdfsdf",
		Frameless:  true,
		Fullscreen: true,
		Mac: &mac.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
		},
		Linux: &linux.Options{
			WindowIsTranslucent: true,
		},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
		// AlwaysOnTop: true,
		AlwaysOnTop: false,
		StartHidden: true,
		OnDomReady:  domReady,
		MaxWidth:    10000,
		MaxHeight:   10000,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func domReady(ctx context.Context) {
	var minX, minY, maxX, maxY int

	n := screenshot.NumActiveDisplays()
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		minX = min(minX, bounds.Min.X)
		maxX = max(maxX, bounds.Max.X)

		minY = min(minY, bounds.Min.Y)
		maxY = max(maxY, bounds.Max.Y)
		fmt.Println(bounds.Max.Y)
	}

	runtime.WindowSetSize(ctx, abs(maxX)+abs(minX), abs(maxY)+abs(minY))
	runtime.WindowSetPosition(ctx, 0, 0)
	runtime.WindowShow(ctx)
}

func abs(x int) int {
	if x < 0 {
		return 0 - x
	}
	return x - 0
}
