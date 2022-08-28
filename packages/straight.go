package main

import "math"

//Starting with uppercase letter means PUBLIC (visible inside and out of the package)
//Starting with lowercase letter means PRIVATE (visible only inside the package)

//For instance
// Point - it will generate something publics
// point  or _Point - it will generate something private

// Point represents a coordinate on cartesian plane
//public
type Point struct {
	x float64
	y float64
}

//func private only for the escope
func peccaries(a, b Point) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

//Distance is responsible for calculate the linear distance between two points
func Distance(a, b Point) float64 {
	cx, cy := peccaries(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
