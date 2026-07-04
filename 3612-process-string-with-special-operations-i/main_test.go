package main

import "testing"

func Test_processStr(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "test case one",
			s:    "a#b%*",
			want: "ba",
		}, {
			name: "test case two",
			s:    "z*#",
			want: "",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processStr(tt.s); got != tt.want {
				t.Errorf("processStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
