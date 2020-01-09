package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math"
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
	renderWidth float64
	renderHeight float64
	fov float64
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
	m := image.NewRGBA(image.Rect(0, 0, int(renderConfig.renderWidth), int(renderConfig.renderHeight))) //*NRGBA (image.Image interface)

	// fill m in blue
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	halfViewportWidth := renderConfig.renderWidth / 2.0
	halfViewportHeight := renderConfig.renderHeight / 2.0

	distToViewPort := halfViewportWidth / math.Atan(renderConfig.fov)

	fmt.Println("dist to viewport: ", distToViewPort)

	vecFromEyeToTopLeftOfViewport := vec3{x: -halfViewportWidth, y: halfViewportHeight, z: distToViewPort}

	planeYCoord := -100.0
	planeStripeWidth := 100.0

	rayStart := vec3{0.0, 0.0, 0.0}

	// draw a line
	for y := 0.0; y < renderConfig.renderHeight; y++ {
		rayDirn := vecFromEyeToTopLeftOfViewport.add(vec3{0.0, - y, 0})
		for x := 0.0; x < renderConfig.renderWidth; x++ {
			rayDirn = rayDirn.add(vec3{1.0, 0.0, 0.0})
			//fmt.Println("Made ray: start, dirn = ", rayStart, rayDirn)

			intersectLambda := (planeYCoord - rayStart.y) / rayDirn.y

			if intersectLambda >= 0 {
				planeIntersection := rayStart.add(rayDirn.mult(intersectLambda))

				if planeIntersection.x < 0 {
					planeIntersection.x -= planeStripeWidth
				}
				if int(planeIntersection.x / planeStripeWidth) % 2 == 0 {
					m.Set(int(x), int(y), white)
				}
			}
		}
		//fmt.Println("============== end row")
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
