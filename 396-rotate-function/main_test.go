package main

import "testing"

func Test_maxRotateFunction(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "test case one",
			args: []int{4, 3, 2, 6},
			want: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRotateFunction(tt.args); got != tt.want {
				t.Errorf("maxRotateFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}
