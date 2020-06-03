package gamestate

import "math/rand"

// UpdateDay is used to update the passing day
func UpdateDay(currentDay int) int {
	var nextDay int

	nextDay = currentDay + rand.Intn(7)

	return nextDay
}
