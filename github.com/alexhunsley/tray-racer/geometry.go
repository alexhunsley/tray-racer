package main

type ray struct {
	// a start point for the ray
	start vec3
	// direction of the ray
	direction vec3
}

// a collection of rays, each with a relative weight.
// the weight is the weight for *each* the rays in the corresponding value.
// for example:
//
//  1.0 -> ray1, ray2
//  2.0 -> ray3, ray4, ray5
//
// ray1 and 2 have weight 1 each.
// rays 3, 4, 5 have weight 2.0 each.
// Note that usually all weights should total 1.0!
type rayOffsetBundle struct {
	weightToRayOffsetMap map[float64][]vec3
}

func (r ray) coord(lambda float64) vec3 {
	return r.start.add(r.direction.mult(lambda))
}
