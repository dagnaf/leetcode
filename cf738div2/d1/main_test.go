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
			args: args{strings.NewReader(`3 2 2
1 2
2 3
1 2
1 3
`)},
			wantW: `0
`,
		},
		{
			args: args{strings.NewReader(`5 3 2
5 4
2 1
4 3
4 3
1 4
`)},
			wantW: `1
2 4
`,
		},
		{
			args: args{strings.NewReader(`8 1 2
1 7
2 6
1 5
`)},
			wantW: `5
5 2
2 3
3 4
4 7
6 8
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			rw(tt.args.r, w)
			if gotW := w.String(); gotW != tt.wantW {
				gotW = gotW[:strings.Index(gotW, "\n")]
				tt.wantW = tt.wantW[:strings.Index(tt.wantW, "\n")]
				if gotW != tt.wantW {
					t.Errorf("rw() = %v, want %v", gotW, tt.wantW)
				}
			}
		})
	}
}
