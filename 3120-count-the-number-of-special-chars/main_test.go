package main

import "testing"

func Test_numberOfSpecialChars(t *testing.T) {
	tests := []struct {
		name string
		word string
		want int
	}{
		{
			name: "test case one",
			word: "aaAbcBC",
			want: 3,
		}, {
			name: "test case two",
			word: "abc",
			want: 0,
		}, {
			name: "test case three",
			word: "abBCab",
			want: 1,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSpecialChars(tt.word); got != tt.want {
				t.Errorf("numberOfSpecialChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
