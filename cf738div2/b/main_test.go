package main

import (
	"bytes"
	"fmt"
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
			args: args{strings.NewReader(`5
7
?R???BR
7
???R???
1
?
1
B
10
?R??RB??B?
`)},
			wantW: `BRRBRBR
BRBRBRB
B
B
BRRBRBBRBR`,
		},
	}
	g := func(s string) (r int) {
		for i := 1; i < len(s); i++ {
			if s[i] == s[i-1] {
				r++
			}
		}
		return
	}
	f := func(s string) string {
		r := []int{}
		for _, s := range strings.Split(s, "\n") {
			r = append(r, g(s))
		}
		return fmt.Sprint(r)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			rw(tt.args.r, w)
			if gotW := w.String(); f(gotW) != f(tt.wantW) {
				t.Errorf("rw() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
