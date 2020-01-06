package main

import (
	"fmt"
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
	"image"
	"image/color"
	"image/draw"
)

var (
	white color.Color = color.RGBA{255, 255, 255, 255}
	black color.Color = color.RGBA{0, 0, 0, 255}
	blue  color.Color = color.RGBA{0, 0, 255, 255}
)

func main() {
	fmt.Println("Hello, Plain")

	gl.StartDriver(appMain)
}

func appMain(driver gxui.Driver) {
	width, height := 640, 480
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	// The themes create the content. Currently only a dark theme is offered for GUI elements.
	theme := dark.CreateTheme(driver)
	img := theme.CreateImage()
	window := theme.CreateWindow(width, height, "Image viewer")
	texture := driver.CreateTexture(m, 1.0)
	img.SetTexture(texture)
	window.AddChild(img)
	window.OnClose(driver.Terminate)
}

//func createImage() {
//	m := image.NewRGBA(image.Rect(0, 0, 640, 480)) //*NRGBA (image.Image interface)
//
//	// fill m in blue
//	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)
//
//	// draw a line
//	for i := m.Bounds().Min.X; i < m.Bounds().Max.X; i++ {
//		m.Set(i, m.Bounds().Max.Y/2, white) // to change a single pixel
//	}
//
//	//w, _ := os.Create("testImage.png")
//	//defer w.Close()
//	//png.Encode(w, m) //Encode writes the Image m to w in PNG format.
//}
