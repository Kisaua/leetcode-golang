package main

func stringIndices(wordsContainer []string, wordsQuery []string) []int {
	result := make([]int, 0, len(wordsQuery))

	sMap, defaultIndex := createSuffixe(wordsContainer)
loop:
	for i := range wordsQuery {
		for j := range wordsQuery[i] {
			suff := wordsQuery[i][j:len(wordsQuery[i])]
			if indexes, ok := sMap[suff]; ok {
				result = append(result, indexes)
				continue loop
			}
		}
		result = append(result, defaultIndex)
	}
	return result
}

func createSuffixe(wordsContainer []string) (map[string]int, int) {
	suffixMap := map[string]int{}
	l := len(wordsContainer[0])
	ind := 0
	for i := range wordsContainer {
		if len(wordsContainer[i]) < l {
			l = len(wordsContainer[i])
			ind = i
		}
		for j := range wordsContainer[i] {
			suff := wordsContainer[i][j:len(wordsContainer[i])]
			if k, ok := suffixMap[suff]; ok {
				if len(wordsContainer[i]) < len(wordsContainer[k]) {
					suffixMap[suff] = i
				}
				continue
			}
			suffixMap[suff] = i
		}
	}

	return suffixMap, ind
}
