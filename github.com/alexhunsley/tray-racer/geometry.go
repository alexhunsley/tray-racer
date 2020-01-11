package main

type ray struct {
	// a start point for the ray
	start vec3
	// direction of the ray
	direction vec3
}

func (r ray) coord(lambda float64) vec3 {
	return r.start.add(r.direction.mult(lambda))
}
