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
			nums: []int{3, 4, 5, 1, 2},
			want: 1,
		}, {
			name: "test case two",
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			want: 0,
		}, {
			name: "test case three",
			nums: []int{11, 13, 15, 17},
			want: 11,
		}, {
			name: "test case four",
			nums: []int{2, 1},
			want: 1,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
