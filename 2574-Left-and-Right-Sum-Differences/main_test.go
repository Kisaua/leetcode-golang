package main

import (
	"reflect"
	"testing"
)

func Test_leftRightDifference(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "test case one",
			nums: []int{10, 4, 8, 3},
			want: []int{15, 1, 11, 22},
		}, {
			name: "test case one",
			nums: []int{1},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leftRightDifference(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("leftRightDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
