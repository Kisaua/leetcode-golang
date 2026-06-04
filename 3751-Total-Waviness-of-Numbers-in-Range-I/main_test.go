package main

import "testing"

func Test_totalWaviness(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case one",
			args: args{
				num1: 120,
				num2: 130,
			},
			want: 3,
		}, {
			name: "test case two",
			args: args{
				num1: 198,
				num2: 202,
			},
			want: 3,
		}, {
			name: "test case three",
			args: args{
				num1: 198,
				num2: 202,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalWaviness(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("totalWaviness() = %v, want %v", got, tt.want)
			}
		})
	}
}
