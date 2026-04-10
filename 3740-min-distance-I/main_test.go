package main

import "testing"

func Test_minimumDistance(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "tets case one",
			nums: []int{1, 2, 1, 1, 3},
			want: 6,
		}, {
			name: "tets case two",
			nums: []int{1, 1, 2, 3, 2, 1, 2},
			want: 8,
		}, {
			name: "tets case three",
			nums: []int{1},
			want: -1,
		}, {
			name: "tets case four",
			nums: []int{1, 1, 1, 1},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDistance(tt.nums); got != tt.want {
				t.Errorf("minimumDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
