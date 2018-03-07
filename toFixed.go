// Package toFixed formats a number using fixed-point notation.
package toFixed

import "math"

// ToFixed formats a number using fixed-point notation.
func ToFixed(n float64, precision int) float64 {
	scale := math.Pow(10, float64(precision))

	return math.Round(n*scale) / scale
}
