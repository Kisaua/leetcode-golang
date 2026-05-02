package main

import "testing"

func Test_rotatedDigits(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{
			name: "test case one",
			args: 10,
			want: 4,
		}, {
			name: "test case two",
			args: 2,
			want: 1,
		}, {
			name: "test case three",
			args: 100,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotatedDigits(tt.args); got != tt.want {
				t.Errorf("rotatedDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
