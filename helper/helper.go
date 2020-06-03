package helper

import "math/rand"

// GenerateRandomInteger used to generate random integer
// within minimum and maximum range
func GenerateRandomInteger(min int, max int) int {
	return rand.Intn(rand.Intn(max-min) + min)
}
