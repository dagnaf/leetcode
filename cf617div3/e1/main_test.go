package main

import (
	"bytes"
	"io"
	"reflect"
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
			args: args{r: strings.NewReader(`9
abacbecfd
`)},
			wantW: `YES
001010101
`,
		},
		{
			args: args{r: strings.NewReader(`8
aaabbcbb
`)},
			wantW: `YES
00000011
`, // 01011011
		},
		{
			args: args{r: strings.NewReader(`7
abcdedc
`)},
			wantW: `NO
`,
		},
		{
			args: args{r: strings.NewReader(`5
abcde
`)},
			wantW: `YES
00000
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

func Test_getSt(t *testing.T) {
	type args struct {
		n int
		s []byte
	}
	tests := []struct {
		name string
		args args
		want st
	}{
		{
			args: args{
				n: 9,
				//         012345678
				s: []byte(`abacbecfd`)},
			want: st{
				s: []byte(`aabbccdef`),
				i: []int{0, 2, 1, 4, 3, 6, 8, 5, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSt(tt.args.n, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRe(t *testing.T) {
	type args struct {
		n  int
		st st
	}
	tests := []struct {
		name   string
		args   args
		wantRe [][2]int
	}{
		{
			args: args{
				n: 9,
				st: st{
					i: []int{0, 2, 1, 4, 3, 6, 8, 5, 7}},
			},
			wantRe: [][2]int{
				{1, 2}, {3, 4}, {5, 6}, {5, 8}, {7, 8},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRe := getRe(tt.args.n, tt.args.st); !reflect.DeepEqual(gotRe, tt.wantRe) {
				t.Errorf("getRe() = %v, want %v", gotRe, tt.wantRe)
			}
		})
	}
}
