package main

func minJumps(arr []int) int {
	l := len(arr)
	dp := make([]int, l)

	arrMap := make(map[int][]int)

	for i := range arr {
		dp[i] = l
		if _, ok := arrMap[arr[i]]; !ok {
			arrMap[arr[i]] = []int{i}
			continue
		}
		arrMap[arr[i]] = append(arrMap[arr[i]], i)
	}
	dp[0] = 0

	for i := range arr {
		for j := i; j < l; j++ {

			if dp[j] > i {
				continue
			}

			if i+1 < l {
				dp[i+1] = min(dp[i+1], dp[i]+1)
			}
			if arr[i] == arr[j] {
				dp[j] = min(dp[j], dp[i]+1)
			}

			// if dp[j] <= i {
			if j-1 >= 0 {
				dp[j-1] = min(dp[j-1], dp[j]+1)
			}
			if j+1 < len(arr) {
				dp[j+1] = min(dp[j+1], dp[j]+1)
			}

			for _, v := range arrMap[arr[j]] {
				dp[v] = min(dp[v], dp[j]+1)
			}
			// }

			// if dp[len(arr)-1] != len(arr) {
			// 	return dp[len(arr)-1]
			// }
			//
		}
	}

	return dp[l-1]
}
