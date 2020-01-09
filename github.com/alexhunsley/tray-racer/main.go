package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"os/exec"
)

var (
	white color.Color = color.RGBA{255, 255, 255, 255}
	black color.Color = color.RGBA{0, 0, 0, 255}
	blue  color.Color = color.RGBA{0, 0, 255, 255}
)

type float float64

type renderConfig struct {
	renderWidth int
	renderHeight int
	fov float
}

func main() {
	fmt.Println("Hello, Plane")

	renderConfig := renderConfig{
		renderWidth:  480,
		renderHeight: 320,
		fov:          90,
	}

	createImage(renderConfig)
}

func createImage(renderConfig renderConfig) {
	m := image.NewRGBA(image.Rect(0, 0, renderConfig.renderWidth, renderConfig.renderHeight)) //*NRGBA (image.Image interface)

	// fill m in blue
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	// draw a line
	for y := 0; y < renderConfig.renderHeight; y++ {
		for x := 0; x < renderConfig.renderWidth; x++ {
			if (x + y) % 10 == 0 || (x - y) % 10 == 0 {
				m.Set(x, y, white)
			}
		}
	}

	w, _ := os.Create("testImage.png")
	defer w.Close()
	png.Encode(w, m) //Encode writes the Image m to w in PNG format.

	Show(w.Name())
}

// show a specified file by Preview.app for OS X(darwin)
func Show(name string) {
	previewCommandPath := "/System/Applications/Preview.app/Contents/MacOS/Preview"

	if _, err := os.Stat(previewCommandPath); os.IsNotExist(err) {
		previewCommandPath = "/Applications/Preview.app/Contents/MacOS/Preview"
	}
	command := "open"
	arg1 := "-a"
	//arg2 :=
	cmd := exec.Command(command, arg1, previewCommandPath, name)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
