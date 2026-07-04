package main

func numberOfSubstrings(s string) int {
	result := 0

	a, b, c := 0, 0, 0
	i := 0
	j := 0
	for {
		if i >= len(s) && (a == 0 || b == 0 || c == 0) {
			return result
		}
		if a > 0 && b > 0 && c > 0 {
			result += len(s) - i + 1
			switch s[j] {
			case 'a':
				a--
			case 'b':
				b--
			case 'c':
				c--
			}
			j++
			continue

		}
		switch s[i] {
		case 'a':
			a++
		case 'b':
			b++
		case 'c':
			c++
		}
		i++
	}

	return result
}
