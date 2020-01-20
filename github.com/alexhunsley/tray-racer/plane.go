package main

type plane struct {
	// the normal of the plane
	orientation vec3
	// a point on the plane
	surfacePoint vec3
}

// returns hit true/false and lambda for any intersection
func (p plane) intersect(r ray) (bool, float64) {
	planeSurfaceToRayStart := p.surfacePoint.sub(r.start)

	topDot := planeSurfaceToRayStart.dot(p.orientation)
	bottomDot := r.direction.dot(p.orientation)

	if bottomDot == 0.0 {
		return false, 0.0
	}
	lambda := topDot / bottomDot
	if lambda <= 0.0 {
		return false, 0.0
	}
	return true, lambda
}
