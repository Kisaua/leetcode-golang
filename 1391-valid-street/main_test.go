package main

import "testing"

func Test_hasValidPath(t *testing.T) {
	tests := []struct {
		name string
		args [][]int
		want bool
	}{
		{
			name: "test case one",
			args: [][]int{{2, 4, 3}, {6, 5, 2}},
			want: true,
		}, {
			name: "test case two",
			args: [][]int{{1, 1, 2}},
			want: false,
		}, {
			name: "test case three",
			args: [][]int{{1, 1, 1, 1, 1, 1, 3}},
			want: true,
		}, {
			name: "test case four",
			args: [][]int{{4, 1}, {6, 1}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasValidPath(tt.args); got != tt.want {
				t.Errorf("hasValidPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
