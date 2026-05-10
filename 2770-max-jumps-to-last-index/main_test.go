package main

import "testing"

func Test_maximumJumps(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case one",
			args: args{
				nums:   []int{1, 3, 6, 4, 1, 2},
				target: 2,
			},
			want: 3,
		}, {
			name: "test case two:",
			args: args{
				nums:   []int{1, 3, 6, 4, 1, 2},
				target: 3,
			},
			want: 5,
		}, {
			name: "test case three",
			args: args{
				nums:   []int{1, 3, 6, 4, 1, 2},
				target: 0,
			},
			want: -1,
		}, {
			name: "test case four",
			args: args{
				nums:   []int{0, 2, 1, 3},
				target: 1,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumJumps(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("maximumJumps() = %v, want %v", got, tt.want)
			}
		})
	}
}
