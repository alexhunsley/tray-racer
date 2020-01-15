package main

import "math"

type vec3 struct {
	x float64
	y float64
	z float64
}

var (
	zeroVec vec3 = vec3{0.0, 0.0, 0.0}
	xUnitVec vec3 = vec3{1.0, 0.0, 0.0}
	yUnitVec vec3 = vec3{0.0, 1.0, 0.0}
	zUnitVec vec3 = vec3{0.0, 0.0, 1.0}
	xUnitVecNegative vec3 = vec3{-1.0, 0.0, 0.0}
	yUnitVecNegative vec3 = vec3{0.0, -1.0, 0.0}
	zUnitVecNegative vec3 = vec3{0.0, 0.0, -1.0}
)

func (v vec3) size() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y + v.z * v.z)
}


func (v vec3) sizeSquared() float64 {
	return v.x * v.x + v.y * v.y + v.z * v.z
}

func (v vec3) add(v2 vec3) vec3 {
	return vec3{v.x + v2.x, v.y + v2.y, v.z + v2.z}
}

func (v vec3) sub(v2 vec3) vec3 {
	return vec3{v.x - v2.x, v.y - v2.y, v.z - v2.z}
}

func (v vec3) mult(m float64) vec3 {
	return vec3{m * v.x , m * v.y, m * v.z}
}

func (v vec3) dot(v2 vec3) float64 {
	return v.x * v2.x + v.y * v2.y + v.z * v2.z
}

func (v vec3) cross(v2 vec3) vec3 {
	return vec3{v.y * v2.z - v.z * v2.y , v.x * v2.z - v.z * v2.x, v.x * v2.y - v.y * v2.x}
}

func (v vec3) unit() vec3 {
	return v.mult(1.0 / v.size())
}
