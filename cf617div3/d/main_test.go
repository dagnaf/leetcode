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
			args: args{r: strings.NewReader(`6 2 3 3
7 10 50 12 1 8`)},
			wantW: `5`,
		},
		{
			args: args{r: strings.NewReader(`1 1 100 99
100`)},
			wantW: `1`,
		},
		{
			args: args{r: strings.NewReader(`7 4 2 1
1 3 5 4 2 7 6`)},
			wantW: `6`,
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

func Test_getK(t *testing.T) {
	type args struct {
		a  int
		b  int
		hi int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{2, 3, 7}, want: 0},
		{args: args{2, 3, 10}, want: 2},
		{args: args{2, 3, 50}, want: 2},
		{args: args{2, 3, 12}, want: 0},
		{args: args{2, 3, 1}, want: 0},
		{args: args{2, 3, 8}, want: 1},
		{args: args{4, 2, 1}, want: 0},
		{args: args{4, 2, 3}, want: 0},
		{args: args{4, 2, 5}, want: 1},
		{args: args{4, 2, 4}, want: 0},
		{args: args{4, 2, 2}, want: 0},
		{args: args{4, 2, 7}, want: 0},
		{args: args{4, 2, 6}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getK(tt.args.a, tt.args.b, tt.args.hi); got != tt.want {
				t.Errorf("getK() = %v, want %v", got, tt.want)
			}
		})
	}
}
