package main

import "testing"

func Test_minimumCost(t *testing.T) {
	tests := []struct {
		name string
		cost []int
		want int
	}{
		{
			name: "test case one",
			cost: []int{1, 2, 3},
			want: 5,
		}, {
			name: "test case two",
			cost: []int{6, 5, 7, 9, 2, 2},
			want: 23,
		}, {
			name: "test case two",
			cost: []int{5, 5},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumCost(tt.cost); got != tt.want {
				t.Errorf("minimumCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
