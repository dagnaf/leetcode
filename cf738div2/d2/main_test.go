package main

import (
	"bytes"
	"io"
	"math/rand"
	"strings"
	"testing"
	"time"
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
				t.Logf("rw() = %v, want %v", gotW, tt.wantW)
				gotW = gotW[:strings.Index(gotW, "\n")]
				tt.wantW = tt.wantW[:strings.Index(tt.wantW, "\n")]
				if gotW != tt.wantW {
					t.Errorf("rw() = %v, want %v", gotW, tt.wantW)
				}
			}
		})
	}
}

func Test_tle(t *testing.T) {
	n, m1, m2 := 100000, 25887, 1242
	f1 := newF(n)
	f2 := newF(n)
	var e1, e2 []edge
	for len(e1) < m1 {
		x := rand.Intn(n) + 1
		y := rand.Intn(n) + 1
		if !same(f1, x, y) {
			union(f1, x, y)
			e1 = append(e1, edge{x, y})
		}
	}
	for len(e2) < m2 {
		x := rand.Intn(n) + 1
		y := rand.Intn(n) + 1
		if !same(f1, x, y) && !same(f2, x, y) {
			union(f2, x, y)
			e2 = append(e2, edge{x, y})
		}
	}
	want := do1(n, e1, e2)
	beg := time.Now()
	got := do2(n, e1, e2)
	t.Log("do2 cost", time.Since(beg))
	if len(want) != len(got) {
		t.Errorf("got = %v, want %v", len(got), len(want))
	}
}
