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
			args: args{strings.NewReader(`3 998244353
`)},
			wantW: `5
`,
		},
		{
			args: args{strings.NewReader(`5 998244353
`)},
			wantW: `25
`,
		},
		{
			args: args{strings.NewReader(`42 998244353
`)},
			wantW: `793019428
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

func Test_tle(t *testing.T) {
	t.Log(do(2e5, 1e8+7))
}

func Test_tle2(t *testing.T) {
	t.Log(do(4e6, 1e8+7))
}
