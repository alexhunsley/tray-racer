package main

// returns a list of rays based on the given ray,
// spread out horizontally. The resulting rays have slightly different start points and directions,
// such that they all converge at r.start + r.direction.
func makeDofRays(r ray, newRayCount int, 	focalDistance float64, spread float64) []ray {
	dofRays := []ray{}

	for rayIndex := 0; rayIndex < newRayCount; rayIndex += 1 {
		xOffset := vec3{(-0.5 + float64(rayIndex) / float64(newRayCount)) * spread, 0, 0}

		dofRay := ray{r.start.add(xOffset), r.direction.sub(xOffset)}

		dofRays = append(dofRays, dofRay)

		///

		yOffset := vec3{0, (-0.5 + float64(rayIndex) / float64(newRayCount)) * spread, 0}

		dofRayY := ray{r.start.add(yOffset), r.direction.sub(yOffset)}

		dofRays = append(dofRays, dofRayY)

		///

		//xyOffset := vec3{(-0.5 + float64(rayIndex) / float64(newRayCount)) * spread, (-0.5 + float64(rayIndex) / float64(newRayCount)) * spread, 0}
		//
		//dofRayXY := ray{r.start.add(xyOffset), r.direction.sub(yOffset)}
		//
		//dofRays = append(dofRays, dofRayXY)
	}
	return dofRays
}
