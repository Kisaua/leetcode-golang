package main

import "testing"

func Test_maxDistance(t *testing.T) {
	tests := []struct {
		name string
		arg  []int
		want int
	}{
		{
			name: "test case one",
			arg:  []int{1, 1, 1, 6, 1, 1, 1},
			want: 3,
		},
		{
			name: "test case two",
			arg:  []int{1, 8, 3, 8, 3},
			want: 4,
		},
		{
			name: "test case three",
			arg:  []int{0, 1},
			want: 1,
		},
		{
			name: "test case four",
			arg:  []int{20, 76, 20, 20, 20, 20, 76, 20, 20, 20, 20, 20, 20, 20, 20, 20},
			want: 8,
		},
		{
			name: "test case five",
			arg:  []int{},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistance(tt.arg); got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
