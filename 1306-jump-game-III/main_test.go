package main

import "testing"

func Test_canReach(t *testing.T) {
	type args struct {
		arr   []int
		start int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case one",
			args: args{
				arr:   []int{4, 2, 3, 0, 3, 1, 2},
				start: 5,
			},
			want: true,
		}, {
			name: "test case two",
			args: args{
				arr:   []int{4, 2, 3, 0, 3, 1, 2},
				start: 0,
			},
			want: true,
		}, {
			name: "test case three",
			args: args{
				arr:   []int{3, 0, 2, 1, 2},
				start: 0,
			},
			want: false,
		}, {
			name: "test case four",
			args: args{
				arr:   []int{4, 4, 1, 3, 0, 3},
				start: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canReach(tt.args.arr, tt.args.start); got != tt.want {
				t.Errorf("canReach() = %v, want %v", got, tt.want)
			}
		})
	}
}
