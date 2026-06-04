package main

func canReach(s string, minJump int, maxJump int) bool {
	if s[len(s)-1] == 49 {
		return false
	}
	stack := make([]r, 0)

	stack = append(stack, r{
		mi: minJump,
		ma: maxJump,
	})

	for len(stack) > 0 {
		l := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if s[l.ma] == 49 {
			if l.ma > l.mi {
				l.ma -= 1
				stack = append(stack, l)
			}
			continue
		}

		if l.ma+maxJump >= len(s)-1 && l.ma+minJump <= len(s)-1 {
			return true
		}

		if l.ma > l.mi {
			l.ma -= 1
			stack = append(stack, l)
		}
		if l.ma+minJump >= len(s) {
			continue
		}

		stack = append(stack, r{
			mi: l.ma + minJump,
			ma: min(l.ma+maxJump, len(s)-1),
		})

	}

	return false
}

type r struct {
	mi, ma int
}

func canReach2(s string, minJump int, maxJump int) bool {
	stack := make([]int, 0)

	stack = append(stack, 0)

	if s[len(s)-1] == 49 {
		return false
	}

	for len(stack) > 0 {
		i := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for j := minJump; j <= maxJump && i+j < len(s); j++ {
			if s[i+j] == 48 {
				if i+j == len(s)-1 {
					return true
				}
				stack = append(stack, i+j)
			}
		}
	}

	return false
}

// canReach1 is a working solution, but TLE
func canReach1(s string, minJump int, maxJump int) bool {
	canJump := make([]int, len(s))

	canJump[0] = 1

	for i := range canJump {
		if canJump[i] == 0 {
			continue
		}
		for j := minJump; j <= maxJump && i+j < len(canJump); j++ {
			if s[i+j] == 48 {
				canJump[i+j] = 1
			}
		}

	}

	return canJump[len(canJump)-1] == 1
}
