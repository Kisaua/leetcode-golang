package main

import "testing"

func Test_findMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "test case one",
			nums: []int{1, 3, 5},
			want: 1,
		}, {
			name: "test case two",
			nums: []int{2, 2, 2, 0, 1},
			want: 0,
		}, {
			name: "test case three",
			nums: []int{1, 1},
			want: 1,
		}, {
			name: "test case four",
			nums: []int{10, 1, 10, 10, 10},
			want: 1,
		}, {
			name: "test case five",
			nums: []int{3, 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
