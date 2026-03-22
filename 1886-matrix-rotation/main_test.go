package main

import "testing"

func Test_findRotation(t *testing.T) {
	type args struct {
		mat    [][]int
		target [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case one",
			args: args{
				mat:    [][]int{{0, 1}, {1, 0}},
				target: [][]int{{1, 0}, {0, 1}},
			},
			want: true,
		},
		{
			name: "test case two",
			args: args{
				mat:    [][]int{{0, 1}, {1, 1}},
				target: [][]int{{1, 0}, {0, 1}},
			},
			want: false,
		},
		{
			name: "test case three",
			args: args{
				mat:    [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}},
				target: [][]int{{1, 1, 1}, {0, 1, 0}, {0, 0, 0}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRotation(tt.args.mat, tt.args.target); got != tt.want {
				t.Errorf("findRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
