package main

import "testing"

func Test_processStr(t *testing.T) {
	type args struct {
		s string
		k int64
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "test case one",
			args: args{
				s: "a#b%*",
				k: 1,
			},
			want: byte('a'),
		}, {
			name: "test case two",
			args: args{
				s: "cd%#*#",
				k: 3,
			},
			want: byte('d'),
		}, {
			name: "test case three",
			args: args{
				s: "%edx#n#lkc####uom##qg#%#b#ek%##%%ocr#m%#fv%i%%#n#u%%#n#q%v#rwvd##t###%#%%%o*##r#gr*gz#dm%ez",
				k: 4780,
			},
			want: byte('d'),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("processStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
