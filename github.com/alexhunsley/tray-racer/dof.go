package main

import (
	"fmt"
	"math"
)

// returns a list of rays based on the given ray,
// spread out horizontally. The resulting rays have slightly different start points and directions,
// such that they all converge at r.start + r.direction.
func makeDofRays(r ray, newRayCount int, focalDistance float64, spread float64) []ray {
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

func MakeRayOffsetsForDof(spread float64) rayOffsetBundle {
	// How many points we divide each circle into. the first item
	// should always be 1 since it's a zero radius circle (i.e. the original ray)
	pointsInCircles := []int{1, 1, 1}

	// relative weights of each circle's contribution
	circleRelativeWeights := []float64{1.0, 2.0, 1.0}

	totalWeight := 0.0
	for _, circRelWeight := range circleRelativeWeights {
		totalWeight += circRelWeight
	}

	rayOffsets := MakeRayOffsetBundle()

	for circleIndex, pointCount := range pointsInCircles {
		fmt.Println("idx = ", circleIndex, " Point count = ", pointCount)

		if circleIndex == 0 {
			rayOffsets.weightToRayOffsetMap[circleRelativeWeights[0] / totalWeight] = []vec3{zeroVec}
			continue
		}
		weightEachRay := circleRelativeWeights[circleIndex] / totalWeight / float64(pointsInCircles[circleIndex])

		circleRadius := spread * (1.0 + float64(circleIndex - 1)) / float64(len(pointsInCircles) - 1)

		// only insert the weight key into the map if it's not there already (we don't want to clobber any existing
		// weight key; they can exist already depending on weight setup given)
		if _, ok := rayOffsets.weightToRayOffsetMap[weightEachRay]; !ok {
			rayOffsets.weightToRayOffsetMap[weightEachRay] = []vec3{}
		}

		for angleIndex := 0; angleIndex < pointsInCircles[circleIndex]; angleIndex++ {
			angle := math.Pi * 2.0 * float64(angleIndex) / float64(len(pointsInCircles))
			coord := vec3{circleRadius * math.Cos(angle), circleRadius * math.Sin(angle), 0.0}
			fmt.Println("Made coord: ", coord)
			rayOffsets.weightToRayOffsetMap[weightEachRay] = append(rayOffsets.weightToRayOffsetMap[weightEachRay], coord)
		}
	}
	return rayOffsets
}
