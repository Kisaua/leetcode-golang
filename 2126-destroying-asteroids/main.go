package main

import (
	"slices"
)

func asteroidsDestroyed(mass int, asteroids []int) bool {
	slices.Sort(asteroids)

	for i := range asteroids {
		if mass < asteroids[i] {
			return false
		}
		mass += asteroids[i]
	}

	return true
}
