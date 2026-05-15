package main

import "testing"

func Test_isGood(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "test case one",
			nums: []int{2, 1, 3},
			want: false,
		}, {
			name: "test case two",
			nums: []int{1, 3, 3, 2},
			want: true,
		}, {
			name: "test case three",
			nums: []int{1, 1},
			want: true,
		}, {
			name: "test case four",
			nums: []int{3, 4, 4, 1, 2, 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isGood(tt.nums); got != tt.want {
				t.Errorf("isGood() = %v, want %v", got, tt.want)
			}
		})
	}
}
