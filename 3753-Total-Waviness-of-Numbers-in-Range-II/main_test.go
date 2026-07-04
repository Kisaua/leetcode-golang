package main

import "testing"

func Test_totalWaviness(t *testing.T) {
	type args struct {
		num1 int64
		num2 int64
	}
	tests := []struct {
		name string
		args args
		want int64
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
				num1: 4848,
				num2: 4848,
			},
			want: 2,
		}, {
			name: "test case four",
			args: args{
				num1: 2549294942,
				num2: 5067104447,
			},
			want: 10871250585,
		}, {
			name: "test case five",
			args: args{
				num1: 8900,
				num2: 9532,
			},
			want: 794,
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
