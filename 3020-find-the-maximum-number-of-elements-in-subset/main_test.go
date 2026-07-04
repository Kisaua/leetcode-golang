package main

import "testing"

func Test_maximumLength(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "test case one",
			nums: []int{5, 4, 1, 2, 2},
			want: 3,
		}, {
			name: "test case one",
			nums: []int{1, 3, 2, 4},
			want: 1,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumLength(tt.nums); got != tt.want {
				t.Errorf("maximumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
