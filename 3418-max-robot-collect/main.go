package main

import "fmt"

func main() {
}

type robot struct {
	sum, rob1, rob2 int
}

func maximumAmount(coins [][]int) int {
	results := make([][]robot, len(coins[0]))
	for i := 0; i < len(coins); i++ {
		for j := 0; j < len(coins[0]); j++ {
			coin := coins[i][j]
			if i == 0 && j == 0 {
				robo := robot{sum: coin}
				robo = *robo.update(coin)
				results[j] = []robot{robo}
				continue
			}
			if i == 0 {
				results[j] = addCoin(results[j-1], coin)
				continue
			}
			if j == 0 {
				results[j] = addCoin(results[j], coin)
				continue
			}

			results[j] = addCoin(append(results[j], results[j-1]...), coin)
		}
	}
	fmt.Println("==========================")
	fmt.Println(results[len(coins[0])-1])
	return findMax(results[len(coins[0])-1])
}

func addCoin(left []robot, coin int) []robot {
	newRobots := make([]robot, 0)
	robotsMap := map[struct{ x, y int }]int{}

	for _, l := range left {
		nl := l
		nl.sum = l.sum + coin
		nl = *nl.update(coin)
		if r, ok := robotsMap[struct{ x, y int }{nl.rob1, nl.rob2}]; ok {
			if r < nl.sum {
				robotsMap[struct{ x, y int }{nl.rob1, nl.rob2}] = nl.sum
			}
			continue
		}

		robotsMap[struct {
			x int
			y int
		}{nl.rob1, nl.rob2}] = nl.sum
		// if nl.rob1 == 0 || nl.rob2 == 0 {
		// 	newRobots = append(newRobots, nl)
		// 	continue
		// }
		// if maxSum == nil {
		// 	maxSum = &nl
		// 	continue
		// }
		//
		// 	if maxSum.coinSum() < nl.coinSum() {
		// 		maxSum = &nl
		// 		continue
		// 	}
		// 	if maxSum.coinSum() == nl.coinSum() {
		// 		newRobots = append(newRobots, nl)
		// 	}
		// }
		//
		// if maxSum != nil {
		// 	newRobots = append(newRobots, *maxSum)
	}
	for k, v := range robotsMap {
		newRobots = append(newRobots, robot{v, k.x, k.y})
	}
	return newRobots
}

func findMax(r []robot) int {
	re := r[0].coinSum()
	for i := range r {
		if re < r[i].coinSum() {
			re = r[i].coinSum()
		}
	}
	return re
}

func (r *robot) coinSum() int {
	return r.sum - r.rob1 - r.rob2
}

func (r *robot) update(coin int) *robot {
	if r.rob1 > coin {
		r.rob1, coin = coin, r.rob1
	}
	if r.rob2 > coin {
		r.rob2 = coin
		return r
	}

	return r
}
