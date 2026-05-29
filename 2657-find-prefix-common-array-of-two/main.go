package main

func findThePrefixCommonArray(A []int, B []int) []int {
	result := make([]int, len(A))
	mapA := map[int]struct{}{
		A[0]: {},
	}
	mapB := map[int]struct{}{
		B[0]: {},
	}

	result[0] = 0
	if A[0] == B[0] {
		result[0] = 1
	}

	for i := 1; i < len(A); i++ {
		if A[i] == B[i] {
			result[i] = result[i-1] + 1
			continue
		}
		s := 2
		if _, ok := mapA[B[i]]; !ok {
			s -= 1
			mapB[B[i]] = struct{}{}
		}

		if _, ok := mapB[A[i]]; !ok {
			s -= 1
			mapA[A[i]] = struct{}{}
		}
		result[i] = result[i-1] + s
	}

	return result
}
