package main

import "testing"

func Test_asteroidsDestroyed(t *testing.T) {
	type args struct {
		mass      int
		asteroids []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case one",
			args: args{
				mass:      10,
				asteroids: []int{3, 9, 19, 5, 21},
			},
			want: true,
		}, {
			name: "test case two",
			args: args{
				mass:      5,
				asteroids: []int{4, 9, 23, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := asteroidsDestroyed(tt.args.mass, tt.args.asteroids); got != tt.want {
				t.Errorf("asteroidsDestroyed() = %v, want %v", got, tt.want)
			}
		})
	}
}
