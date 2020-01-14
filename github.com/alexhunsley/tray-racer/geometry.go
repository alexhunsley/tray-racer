package main

type ray struct {
	// a start point for the ray
	start vec3
	// direction of the ray
	direction vec3
}

// a collection of rays, each with a relative weight.
// the weight is the total weight for *all* the rays in the corresponding value.
// for example:
//
//  1.0 -> ray1, ray2
//  2.0 -> ray3, ray4, ray5
//
// Here we have ray1 and ray2 cumulatively having weight 1.0 (so they are weighted 0.5 each),
// whereas ray3, ray4 and ray5 have cumulative weight of 2.0, hence each have a weight of 2/3.
type rayBundle struct {
	weightToRaysMap map[float64][]ray
}

func (r ray) coord(lambda float64) vec3 {
	return r.start.add(r.direction.mult(lambda))
}
