package main

type plane struct {
	// the normal of the plane
	orientation vec3
	// a point on the plane
	surfacePoint vec3
}

// returns lambda for the intersection
func (p plane) intersect(r ray) float64 {
	planeSurfaceToRayStart := p.surfacePoint.sub(r.start)

	topDot := planeSurfaceToRayStart.dot(p.orientation)
	bottomDot := r.direction.dot(p.orientation)

	if bottomDot == 0.0 {
		return -9999
	}
	return topDot / bottomDot
}
