package main

import "testing"

func Test_earliestFinishTime(t *testing.T) {
	type args struct {
		landStartTime  []int
		landDuration   []int
		waterStartTime []int
		waterDuration  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case one",
			args: args{
				landStartTime:  []int{82, 14},
				landDuration:   []int{42, 30},
				waterStartTime: []int{54, 6},
				waterDuration:  []int{71, 91},
			},
			want: 125,
		}, {
			name: "test case two",
			args: args{
				landStartTime:  []int{49, 62},
				landDuration:   []int{42, 49},
				waterStartTime: []int{40, 99, 50, 78, 26, 96, 87, 84, 51, 34, 52, 35, 10, 74},
				waterDuration:  []int{35, 79, 66, 32, 2, 42, 73, 46, 42, 2, 40, 95, 66, 19},
			},
			want: 91,
		}, {
			name: "test case three",
			args: args{
				landStartTime:  []int{20, 84, 92, 82, 14, 52, 93, 18, 43, 55, 52, 80, 43, 60, 66, 37, 100, 98, 91, 5, 22, 24, 23, 91},
				landDuration:   []int{57, 80, 46, 76, 97, 93, 45, 7, 15, 8, 97, 34, 4, 61, 12, 60, 75, 93, 5, 67, 66, 55, 2, 65},
				waterStartTime: []int{21, 44},
				waterDuration:  []int{38, 41},
			},
			want: 61,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := earliestFinishTime(tt.args.landStartTime, tt.args.landDuration, tt.args.waterStartTime, tt.args.waterDuration); got != tt.want {
				t.Errorf("earliestFinishTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
