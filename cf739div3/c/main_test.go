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
			args: args{r: strings.NewReader(`7
11
14
5
4
1
2
1000000000
`)},
			wantW: `2 4
4 3
1 3
2 1
1 1
1 2
31623 14130
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
