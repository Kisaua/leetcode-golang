package main

import "testing"

func Test_canReach(t *testing.T) {
	type args struct {
		s       string
		minJump int
		maxJump int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case one",
			args: args{
				s:       "01101000",
				minJump: 2,
				maxJump: 3,
			},
			want: true,
		}, {
			name: "test case two",
			args: args{
				s:       "01101110",
				minJump: 2,
				maxJump: 3,
			},
			want: false,
		}, {
			name: "test case three",
			args: args{
				s:       "00111010",
				minJump: 3,
				maxJump: 5,
			},
			want: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canReach(tt.args.s, tt.args.minJump, tt.args.maxJump); got != tt.want {
				t.Errorf("canReach() = %v, want %v", got, tt.want)
			}
		})
	}
}
