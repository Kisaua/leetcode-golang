package main

func canReach(arr []int, start int) bool {
	q := []int{start}
	visited := map[int]struct{}{}

	for len(q) > 0 {
		el := q[0]
		q = q[1:]
		if arr[el] == 0 {
			return true
		}

		if _, ok := visited[el]; ok {
			continue
		}
		if el-arr[el] >= 0 {
			visited[el] = struct{}{}
			q = append(q, el-arr[el])
		}

		if el+arr[el] < len(arr) {
			visited[el] = struct{}{}
			q = append(q, el+arr[el])
		}
	}

	return false
}
