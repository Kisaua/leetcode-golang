package main

// a#b%*
// * remove last char
// # duplicate result
// % reverse
func processStr(s string) string {
	result := make([]byte, 0, len(s))
	forward := 1
	for i := range s {
		switch s[i] {
		case 42:
			if len(result) > 0 {
				if forward == 1 {
					result = result[:len(result)-1]
					continue
				}
				result = result[1:]
			}
		case 37:
			forward *= -1
		case 35:
			result = append(result, result...)
		default:
			if forward == 1 {
				result = append(result, s[i])
				continue
			}
			result = append([]byte{s[i]}, result...)

		}
	}

	if forward == -1 {
		for i := 0; i < len(result)/2; i++ {
			result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
		}
	}

	return string(result)
}
