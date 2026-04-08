package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(solutionTwo(1, 4))
}

func getHappyString(n int, k int) string {
	if n == 0 || 3*int(math.Pow(2, float64(n-1))) < k {
		return ""
	}

	h := happyStings([]string{}, n)

	return h[k-1]
}

func happyStings(happy []string, n int) []string {
	defaultStrings := []string{"a", "b", "c"}

	if n == 1 {
		return defaultStrings
	}

	unhappy := happyStings(happy, n-1)

	resultStrings := make([]string, len(unhappy)*2)
	for i := range unhappy {
		switch {
		case strings.HasSuffix(unhappy[i], "a"):
			resultStrings[2*i] = unhappy[i] + "b"
			resultStrings[2*i+1] = unhappy[i] + "c"
		case strings.HasSuffix(unhappy[i], "b"):
			resultStrings[2*i] = unhappy[i] + "a"
			resultStrings[2*i+1] = unhappy[i] + "c"
		case strings.HasSuffix(unhappy[i], "c"):
			resultStrings[2*i] = unhappy[i] + "a"
			resultStrings[2*i+1] = unhappy[i] + "b"
		}
	}
	return resultStrings
}

// solutionTwo calculate by the incoming number K
// ["aba", "abc", "aca", "acb", "bab", "bac", "bca", "bcb", "cab", "cac", "cba", "cbc"]
// as the list ordered, 0-4 starts always from a, 4-8 from b and 8-12 from c,
// so 9 / 4 gives 3 - > firs letter is "c", reminder 1
//
//	1 / 2 gives  0 -> we need to take  myMap["c"] from prev result and take element 0, reminder is 1
//
// repeating, 1/2 gives  0 -> we need to take myMap["a"][0], "a" is from prev steep
func solutionTwo(n, k int) string {
	if n == 0 || 3*int(math.Pow(2, float64(n-1))) < k {
		return ""
	}

	var result string

	defaultString := []string{"a", "b", "c"}

	myMap := map[string][]string{
		"a": {"b", "c"},
		"b": {"a", "c"},
		"c": {"a", "b"},
	}

	rest := (k - 1) / int(math.Pow(2, float64(n-1)))
	reminder := (k - 1) % int(math.Pow(2, float64(n-1)))
	result += defaultString[rest]
	key := result
	for i := 1; i < n; i++ {
		rest = reminder / int(math.Pow(2, float64(n-1-i)))
		reminder = reminder % int(math.Pow(2, float64(n-1-i)))
		result += myMap[key][rest]
		key = myMap[key][rest]
	}

	return result
}
