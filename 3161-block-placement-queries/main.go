package main

import "fmt"

func getResults(queries [][]int) []bool {
	results := make([]bool, 0)
	d := map[int]int{}

	for i := range queries {
		if queries[i][0] == 1 {
			addPoint(queries[i], d)
			continue
		}
		results = append(results, place(queries[i], d))
	}

	fmt.Println(d)

	return results
}

func addPoint(point []int, d map[int]int) {
	if len(d) == 0 {
		d[0] = point[1]
		return
	}
	if _, ok := d[point[1]]; ok {
		return
	}
	for i := point[1] - 1; i >= 0; i-- {
		if v, ok := d[i]; ok {
			d[i] = point[1] - i
			d[point[1]] = v - point[1]
			return
		}
	}
}

func place(block []int, d map[int]int) bool {
	l := 0
	for k, v := range d {
		if l > block[2] {
			continue
		}
		l = max(l, k)
		if v >= block[1] {
			return true
		}
	}

	return false
}
