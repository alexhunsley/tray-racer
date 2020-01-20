package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
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
		renderWidth:  1024,
		renderHeight: 768,
		//renderWidth:  10,
		//renderHeight: 10,
		fov:          80,
		focalPlaneDistFromViewport: 300,
		sampleCount:  1,
	}

	createImage(renderConfig)
}

func createImage(renderConfig renderConfig) {
	m := image.NewRGBA(image.Rect(0, 0, int(renderConfig.renderWidth), int(renderConfig.renderHeight))) //*NRGBA (image.Image interface)

	halfViewportWidth := renderConfig.renderWidth / 2.0
	halfViewportHeight := renderConfig.renderHeight / 2.0

	distToViewPort := halfViewportWidth / math.Atan(math.Pi * renderConfig.fov / 180.0)

	fmt.Println("fov, dist to viewport: ", renderConfig.fov, distToViewPort)

	vecFromEyeToTopLeftOfViewport := vec3{x: -halfViewportWidth, y: halfViewportHeight, z: distToViewPort}

	//plane := plane{orientation: vec3{0.1, 1.0, -0.5}, surfacePoint: vec3{0.0, -10.0, 0.0}}

	planeStripeWidth := 100.0

	rayStart := vec3{0.0, 0.0, -distToViewPort}
	//bodge := 400

	plane1 := sceneobject{plane{orientation: vec3{0.0, 1.0, -0.3}, surfacePoint: vec3{0.0, -100.0, 0.0}}, material{vec3{40.0, 255.0, 0.0}}}
	plane2 := sceneobject{plane{orientation: vec3{1.0, -0.4, 0.0}, surfacePoint: vec3{250.0, 0.0, 0.0}}, material{vec3{255.0, 255.0, 0.0}}}

	objects := []sceneobject{plane1, plane2}

	for y := 0.0; y < renderConfig.renderHeight; y++ {
		rayDirn := vecFromEyeToTopLeftOfViewport.add(vec3{0.0, - y, 0})
		for x := 0.0; x < renderConfig.renderWidth; x++ {
			rayDirn = rayDirn.add(vec3{1.0, 0.0, 0.0})
			//fmt.Println("Made ray: start, dirn = ", rayStart, rayDirn)

			aggregateResultColour := vec3{0.0, 0.0, 0.0}

			r := ray{start: rayStart, direction: rayDirn}

			// TODO make an interface for ray bundle producers
			//dofRays := makeDofRays(r, int(renderConfig.sampleCount), 50, 20)
			dofRays := []ray{r}

			for _, r := range dofRays {

				resultColour := vec3{0.0, 0.0, 0.0}

				//perturbedRayDirn := rayDirn.add(vec3{rand.Float64(), rand.Float64(), rand.Float64()})
				//perturbedRayDirn := rayDirn

				//r := ray{start: rayStart, direction: perturbedRayDirn}
				//intersectLambda := (planeYCoord - rayStart.y) / perturbedRayDirn.y

				//intersectLambdas := []float64{}

				closestObjectHitLambda := 0.0
				var closestObjectHit sceneobject

				for _, object := range objects {

					didHit, intersectionLambda := object.primitive.intersect(r)
					//intersectLambdas = append(intersectLambdas, object.intersect(r))

					//fmt.Println("didHit, lambda = ", didHit, intersectionLambda)
					if didHit && (closestObjectHit.primitive == nil || intersectionLambda < closestObjectHitLambda) {
						closestObjectHit = object
						closestObjectHitLambda = intersectionLambda
						intersectionLambda = closestObjectHitLambda
					}
				}
				if closestObjectHit.primitive != nil {
					objectIntersection := r.coord(closestObjectHitLambda)

					//if int(x) % bodge == 0 && int(y) % bodge == 0 {
					//	fmt.Println("For ray ", r, " plane ix = ", objectIntersection)
					//}

					// TODO this detail finding should be in the object too
					if objectIntersection.x * objectIntersection.x + objectIntersection.z * objectIntersection.z < 1000 {
						// red dot so we can see origin
						resultColour = vec3{200.0, 0.0, 0.0}
					} else {
						if objectIntersection.x < 0 {
							objectIntersection.x -= planeStripeWidth
						}
						if objectIntersection.z < 0 {
							objectIntersection.z -= planeStripeWidth
						}

						if (int(objectIntersection.x/planeStripeWidth)+int(objectIntersection.z/planeStripeWidth))%2 == 0 {
							resultColour = closestObjectHit.material.colour
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
