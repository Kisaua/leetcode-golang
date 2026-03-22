package main

import "testing"

func Test_countSubmatrices(t *testing.T) {
	type args struct {
		grid [][]int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case one",
			args: args{
				grid: [][]int{{7, 6, 3}, {6, 6, 1}},
				k:    18,
			},
			want: 4,
		},
		{
			name: "test case two",
			args: args{
				grid: [][]int{{7, 2, 9}, {1, 5, 0}, {2, 6, 6}},
				k:    20,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubmatrices(tt.args.grid, tt.args.k); got != tt.want {
				t.Errorf("countSubmatrices() = %v, want %v", got, tt.want)
			}
		})
	}
}
