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
	focalPlaneDistFromViewport float64
	sampleCount float64
}

func main() {
	fmt.Println("Hello, Plane")

	//r := ray{start: vec3{0, 0, 0}, direction: vec3{0, 0, 100}}
	//dofRays := makeDofRays(r, 4, 50, 1)
	//fmt.Println("dofRays = ", dofRays)
	//
	//fmt.Println("============ Intersect test: ")
	//p := plane{surfacePoint: vec3{0, 0,0}, orientation: vec3{0, 0,-1 }}
	//r := ray{start: vec3{1, 2, -1}, direction: vec3{10, 20, 1}}
	//
	//lambda := p.intersect(r)
	//fmt.Println("lambda = ", lambda)
	//ixPoint := r.coord(lambda)
	//fmt.Println("lambda = ", lambda, " ix = ", ixPoint)

	renderConfig := renderConfig{
		renderWidth:  1280,
		renderHeight: 1024,
		fov:          80,
		focalPlaneDistFromViewport: 300,
		sampleCount:  32,
	}

	createImage(renderConfig)
}

func createImage(renderConfig renderConfig) {
	m := image.NewRGBA(image.Rect(0, 0, int(renderConfig.renderWidth), int(renderConfig.renderHeight))) //*NRGBA (image.Image interface)

	// fill m in blue
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	halfViewportWidth := renderConfig.renderWidth / 2.0
	halfViewportHeight := renderConfig.renderHeight / 2.0

	distToViewPort := halfViewportWidth / math.Atan(math.Pi * renderConfig.fov / 180.0)

	fmt.Println("fov, dist to viewport: ", renderConfig.fov, distToViewPort)

	vecFromEyeToTopLeftOfViewport := vec3{x: -halfViewportWidth, y: halfViewportHeight, z: distToViewPort}

	//plane := plane{orientation: vec3{0.1, 1.0, -0.5}, surfacePoint: vec3{0.0, -10.0, 0.0}}
	plane := plane{orientation: vec3{0.0, 1.0, -0.3}, surfacePoint: vec3{0.0, -100.0, 0.0}}

	planeStripeWidth := 100.0

	rayStart := vec3{0.0, 0.0, -distToViewPort}
	//bodge := 400

	// draw a line
	for y := 0.0; y < renderConfig.renderHeight; y++ {
		rayDirn := vecFromEyeToTopLeftOfViewport.add(vec3{0.0, - y, 0})
		for x := 0.0; x < renderConfig.renderWidth; x++ {
			rayDirn = rayDirn.add(vec3{1.0, 0.0, 0.0})
			//fmt.Println("Made ray: start, dirn = ", rayStart, rayDirn)

			aggregateResultColour := vec3{0.0, 0.0, 0.0}

			r := ray{start: rayStart, direction: rayDirn}
			dofRays := makeDofRays(r, int(renderConfig.sampleCount), 50, 20)

			//if int(x) % bodge == 0 && int(y) % bodge == 0 {
			//	fmt.Println("ray, dof rays = ", r, dofRays)
			//}
			for _, r := range dofRays {

				resultColour := vec3{0.0, 0.0, 0.0}

				//perturbedRayDirn := rayDirn.add(vec3{rand.Float64(), rand.Float64(), rand.Float64()})
				//perturbedRayDirn := rayDirn

				//r := ray{start: rayStart, direction: perturbedRayDirn}
				//intersectLambda := (planeYCoord - rayStart.y) / perturbedRayDirn.y

				intersectLambda := plane.intersect(r)

				if intersectLambda >= 0 {
					planeIntersection := r.start.add(r.direction.mult(intersectLambda))

					//if int(x) % bodge == 0 && int(y) % bodge == 0 {
					//	fmt.Println("For ray ", r, " plane ix = ", planeIntersection)
					//}

					if planeIntersection.x * planeIntersection.x + planeIntersection.z * planeIntersection.z < 1000 {
						resultColour = vec3{200.0, 0.0, 0.0}
					} else {
						if planeIntersection.x < 0 {
							planeIntersection.x -= planeStripeWidth
						}
						if planeIntersection.z < 0 {
							planeIntersection.z -= planeStripeWidth
						}

						if (int(planeIntersection.x/planeStripeWidth)+int(planeIntersection.z/planeStripeWidth))%2 == 0 {
							resultColour = vec3{200.0, 200.0, 200.0}
						}
					}
				}
				aggregateResultColour = aggregateResultColour.add(resultColour.mult(1.0 / float64(len(dofRays))))
			}
			m.Set(int(x), int(y), color.RGBA{uint8(aggregateResultColour.x), uint8(aggregateResultColour.y), uint8(aggregateResultColour.z), 255})
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
