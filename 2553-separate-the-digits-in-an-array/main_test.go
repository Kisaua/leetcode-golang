package main

import (
	"reflect"
	"testing"
)

func Test_separateDigits(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "test case one",
			args: []int{13, 25, 83, 77},
			want: []int{1, 3, 2, 5, 8, 3, 7, 7},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := separateDigits(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("separateDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
