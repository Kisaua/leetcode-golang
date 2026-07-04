package main

import "testing"

func Test_maxBuilding(t *testing.T) {
	type args struct {
		n            int
		restrictions [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case one",
			args: args{
				n:            5,
				restrictions: [][]int{{2, 1}, {4, 1}},
			},
			want: 2,
		}, {
			name: "test case two",
			args: args{
				n:            6,
				restrictions: [][]int{},
			},
			want: 5,
		}, {
			name: "test case three",
			args: args{
				n:            10,
				restrictions: [][]int{{5, 3}, {2, 5}, {7, 4}, {10, 3}},
			},
			want: 5,
			// }, {
			// 	name: "test case four",
			// 	args: args{
			// 		n:            1000000000,
			// 		restrictions: testCase4,
			// 	},
			// 	want: 10933396,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxBuilding(tt.args.n, tt.args.restrictions); got != tt.want {
				t.Errorf("maxBuilding() = %v, want %v", got, tt.want)
			}
		})
	}
}
