package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_rotate(t *testing.T) {
	tests := []struct {
		name string
		args [][]int
		want [][]int
	}{
		{
			name: "test case one",
			args: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		}, {
			name: "test case one",
			args: [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			want: [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args)
			if !cmp.Equal(tt.args, tt.want) {
				t.Errorf("maxRotateFunction() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}
