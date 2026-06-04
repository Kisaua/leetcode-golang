package main

import (
	"sort"
)

type ByKey struct {
	keys   []int
	values []int
}

func (pair ByKey) Len() int { return len(pair.keys) }
func (pair ByKey) Less(i, j int) bool {
	return pair.keys[i]+pair.values[i] < pair.keys[j]+pair.values[j]
}

func (pair ByKey) Swap(i, j int) {
	pair.keys[i], pair.keys[j] = pair.keys[j], pair.keys[i]
	pair.values[i], pair.values[j] = pair.values[j], pair.values[i]
}

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	result := max(landStartTime[0]+landDuration[0], waterStartTime[0]) + waterDuration[0]

	sort.Sort(ByKey{keys: landStartTime, values: landDuration})
	sort.Sort(ByKey{keys: waterStartTime, values: waterDuration})

	for j := range waterStartTime {
		result = min(result, max(landStartTime[0]+landDuration[0], waterStartTime[j])+waterDuration[j])
	}

	for j := range landStartTime {
		result = min(result, max(waterStartTime[0]+waterDuration[0], landStartTime[j])+landDuration[j])
	}

	return result
}
