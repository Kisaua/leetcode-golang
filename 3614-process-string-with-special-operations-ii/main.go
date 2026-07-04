package main

func processStr(s string, k int64) byte {
	l := int64(0)
	for i := range s {
		switch s[i] {
		case 42:
			if l != 0 {
				l--
			}
		case 37:
		case 35:
			l *= 2
		default:
			l += 1
		}
	}
	if l <= k {
		return '.'
	}
	for i := range s {
		switch s[len(s)-i-1] {
		case 42:
			l++
		case 37:
			k = l - 1 - k
		case 35:
			if k+1 > (l+1)/2 {
				k -= l / 2
			}
			l = (l + 1) / 2

		default:
			if k == l {
				return s[len(s)-i-1]
			}
			l--
		}
	}

	return '.'
}
