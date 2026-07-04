package main

import "fmt"

func minScore(n int, roads [][]int) int {
	roadMap := map[int][][]int{}

	for i := range roads {
		roadMap[roads[i][0]] = append(roadMap[roads[i][0]], roads[i])
		roadMap[roads[i][1]] = append(roadMap[roads[i][1]], roads[i])
	}

	q := make([][]int, 0)
	q = append(q, roadMap[1]...)
	res := roadMap[1][0][2]
	delete(roadMap, 1)

	// fmt.Println(roadMap)

	for len(q) > 0 {
		fmt.Println(q)
		el := q[len(q)-1]
		q = q[:len(q)-1]
		res = min(res, el[2])
		q = append(q, roadMap[el[1]]...)
		q = append(q, roadMap[el[0]]...)
		delete(roadMap, el[1])
		delete(roadMap, el[0])
	}

	return res
}
