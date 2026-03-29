package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(canBeEqual("abcd", "cdab"))
	fmt.Println(canBeEqual("abcd", "dacb"))
}

func canBeEqual(s1 string, s2 string) bool {
	arrS1 := strings.Split(s1, "")
	arrS2 := strings.Split(s2, "")

	return isEq(arrS1[0], arrS2[0], arrS1[2], arrS2[2]) &&
		isEq(arrS1[1], arrS2[1], arrS1[3], arrS2[3])
}

func isEq(a, b, a1, b1 string) bool {
	return (a == b && a1 == b1) || (a == b1 && a1 == b)
}
