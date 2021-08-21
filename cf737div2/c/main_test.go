package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func Test_rw(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			args: args{strings.NewReader(`3
3 1
2 1
4 0
`)},
			wantW: `5
2
1
`,
		},
		{
			args: args{strings.NewReader(`5
200000 200000
24567 23423
1 200000
200000 1
200000 0
`)},
			wantW: `226490044
470587519
175895282
87947641
1
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			rw(tt.args.r, w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("rw() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_getGrEq(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		want1 int64
	}{
		{args: args{1}},
		{args: args{2}},
		{args: args{3}},
		{args: args{4}},
		{args: args{10}},
		{args: args{11}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getGrEq(tt.args.n)
			tt.want, tt.want1 = getGrEq2(tt.args.n)
			if got != tt.want {
				t.Errorf("getGrEq() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getGrEq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func getGrEq2(n int) (int64, int64) {
	var gr, eq int64
	for i := 0; i < (1 << n); i++ {
		var and, xor int = 1, 0
		for j := 0; j < n; j++ {
			b := (i >> j) & 1
			and &= b
			xor ^= b
		}
		if and == xor {
			eq++
		} else if and > xor {
			gr++
		}
	}
	return gr, eq
}

func Test_getGrEqHack(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		want1 int64
	}{
		{args: args{4}, want: 1, want1: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getGrEq2(tt.args.n)
			if got != tt.want {
				t.Errorf("getGrEq2() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getGrEq2() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
