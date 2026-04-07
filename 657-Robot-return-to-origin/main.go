package main

import "strings"

func main() {
}

func judgeCircle(moves string) bool {
	l := strings.Count(moves, "L")
	r := strings.Count(moves, "R")
	u := strings.Count(moves, "U")
	d := strings.Count(moves, "D")
	if (l-r) == 0 && (u-d) == 0 {
		return true
	}
	return false
}
