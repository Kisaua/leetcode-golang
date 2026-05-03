package main

// Other possible solution is concatenate s + s and use strings.Contains goal in the string
func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	for i := range len(s) {
		if s[i:]+s[:i] == goal {
			return true
		}
	}

	return false
}
