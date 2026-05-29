package main

func numberOfSpecialChars(word string) int {
	count := map[byte]struct{}{}

	bytes := []byte(word)

	mem := map[byte]struct{}{}

	for i := range bytes {
		mem[bytes[i]] = struct{}{}
		if bytes[i] >= 97 {
			_, ok := mem[bytes[i]-32]
			if ok {
				count[bytes[i]-32] = struct{}{}
			}
			continue
		}
		if _, ok := mem[bytes[i]+32]; ok {
			count[bytes[i]] = struct{}{}
		}

	}

	return len(count)
}
