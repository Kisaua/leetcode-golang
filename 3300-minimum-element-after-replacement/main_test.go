package main

import "testing"

func Test_minElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "test case one",
			nums: []int{10, 12, 13, 14},
			want: 1,
		}, {
			name: "test case two",
			nums: []int{1, 2, 3, 4},
			want: 1,
		}, {
			name: "test case three",
			nums: []int{999, 19, 199},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minElement(tt.nums); got != tt.want {
				t.Errorf("minElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
