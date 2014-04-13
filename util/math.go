package util

import (
	"math/rand"
)

func Randf(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func Clamp(v, min, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
