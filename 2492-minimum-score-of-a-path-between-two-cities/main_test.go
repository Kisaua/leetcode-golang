package main

import "testing"

func Test_minScore(t *testing.T) {
	type args struct {
		n     int
		roads [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case one",
			args: args{
				n:     4,
				roads: [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}},
			},
			want: 5,
		}, {
			name: "test case two",
			args: args{
				n:     4,
				roads: [][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}},
			},
			want: 2,
		}, {
			name: "test case three",
			args: args{
				n:     14,
				roads: [][]int{{2, 9, 2308}, {2, 5, 2150}, {12, 3, 4944}, {13, 5, 5462}, {2, 10, 2187}, {2, 12, 8132}, {2, 13, 3666}, {4, 14, 3019}, {2, 4, 6759}, {2, 14, 9869}, {1, 10, 8147}, {3, 4, 7971}, {9, 13, 8026}, {5, 12, 9982}, {10, 9, 6459}},
			},
			want: 2150,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minScore(tt.args.n, tt.args.roads); got != tt.want {
				t.Errorf("minScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
