package main

import (
	"context"
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) TakeScreenshot(startX int, startY int, endX int, endY int) string {
	var minX, minY int

	n := screenshot.NumActiveDisplays()
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		minX = min(minX, bounds.Min.X)
		minY = min(minY, bounds.Min.Y)
	}

	output := ""

	bounds := image.Rectangle{
		Min: image.Point{X: startX - abs(minX), Y: startY - abs(minY)},
		Max: image.Point{X: endX - abs(minX), Y: endY - abs(minY)},
	}

	fmt.Println(bounds)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}
	fileName := fmt.Sprintf("%d_%dx%d.png", 1, bounds.Dx(), bounds.Dy())
	file, _ := os.Create(fileName)
	defer file.Close()
	png.Encode(file, img)

	output = output + "\n" + fileName

	return output
}
