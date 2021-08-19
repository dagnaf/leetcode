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
			args: args{r: strings.NewReader(`4
1 1
221 2
177890 2
998244353 1
`)},
			wantW: `1
221
181111
999999999
`,
		},
		{
			args: args{r: strings.NewReader(`6
2021 3
177890 2
34512 3
724533 4
998244353 1
12345678 10
`)},
			wantW: `2021
181111
34533
724542
999999999
12345678
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
