package main

func numberOfSpecialChars(word string) int {
	count := map[byte]struct{}{}

	bytes := []byte(word)

	mem := map[byte]struct{}{}

	for i := range bytes {
		if bytes[i] < 97 {
			_, ok := mem[bytes[i]+32]
			_, ok1 := mem[bytes[i]]
			if ok && !ok1 {
				count[bytes[i]+32] = struct{}{}
			}
			mem[bytes[i]] = struct{}{}
			continue
		}

		mem[bytes[i]] = struct{}{}
		delete(count, bytes[i])

	}

	return len(count)
}
