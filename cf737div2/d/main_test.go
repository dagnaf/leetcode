package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
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
			args: args{strings.NewReader(`3 6
1 1 1
1 7 8
2 7 7
2 15 15
3 1 1
3 15 15
`)},
			wantW: `0
`,
		},
		{
			args: args{strings.NewReader(`5 4
1 2 3
2 4 6
3 3 5
5 1 1
`)},
			wantW: `3
2 4 5`,
		},
		{
			args: args{strings.NewReader(`64 19
2 9775794 295412328
34 22564262 484762546
47 872839411 887635088
49 83594633 817611693
27 604327714 712536348
26 533779349 604921540
27 382068872 557699922
52 2730152 842526047
51 646434297 820428761
2 733997272 815804831
34 111665087 362220147
37 110660604 635463204
50 107786695 903716745
63 539814017 621057900
62 773918219 800449448
42 845633131 939443202
4 920809791 986989515
27 714553514 720830618
45 885332416 909303283
`)},
			wantW: `55
1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 28 29 30 31 32 33 35 36 38 39 40 41 42 43 44 45 46 47 48 53 54 55 56 57 58 59 60 61 63 64`,
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

const file = "case_tle"

func init() {
	createDataFile(file)
}

func createDataFile(file string) {
	f, err := os.Create(file)
	if err != nil {
		return
	}
	defer f.Close()
	var n, m int = 3e5, 3e5
	fmt.Fprint(f, n, m)
	fmt.Fprintln(f)
	for i := 0; i < m; i++ {
		bi := rand.Intn(n) + 1
		l := rand.Intn(1e9) + 1
		r := rand.Intn(1e9) + 1
		if l > r {
			l, r = r, l
		}
		fmt.Fprint(f, bi, l, r)
		fmt.Fprintln(f)
	}
}

// Test_tle n + mlogm
func Test_tle(t *testing.T) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()
	w := &bytes.Buffer{}
	rw(bufio.NewReader(f), w)
	// t.Log(w.String())
	os.Remove(file)
}
