package main

import "testing"

func Test_numberOfSubstrings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "test case one",
			s:    "abcabc",
			want: 10,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSubstrings(tt.s); got != tt.want {
				t.Errorf("numberOfSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
