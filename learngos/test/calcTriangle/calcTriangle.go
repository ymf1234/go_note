package calcTriangle

import "math"

func CalcTriangle(a, b int) int {
	c := int(math.Sqrt(float64(a * a + b * b)))
	return c
}