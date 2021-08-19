package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var _, _, _ = main, rw, do

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	rw(r, w)
}

func rw(r io.Reader, w io.Writer) {
	var t int
	var s string
	fmt.Fscan(r, &t)
	for ti := 0; ti < t; ti++ {
		fmt.Fscan(r, &s)
		w1, w2, w3 := do(s)
		if w3 {
			fmt.Fprint(w, w1, " ", w2)
		} else {
			fmt.Fprint(w, -1)
		}
		fmt.Fprintln(w)
	}
}

func do(t string) (string, string, bool) {
	m := map[byte]int{}
	var d string
	for i := len(t) - 1; i >= 0; i-- {
		m[t[i]]++
		if m[t[i]] == 1 {
			d = string(t[i]) + d
		}
	}
	l := 0
	for i := range d {
		if m[d[i]]%(i+1) != 0 {
			return "", "", false
		}
		l += m[d[i]] / (i + 1)
	}
	s := t[:l]
	a := 0
	b := l
	for i := range d {
		l -= (m[d[i]] / (i + 1))
		c := b + l
		if !chk(t[a:b], t[b:c], d[i]) {
			return "", "", false
		}
		a, b = b, c
	}
	return s, d, true
}

func chk(a, b string, c byte) bool {
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] == c {
			i++
			continue
		}
		if a[i] != b[j] {
			return false
		}
		i++
		j++
	}
	for i < len(a) && a[i] == c {
		i++
	}
	return i >= len(a) && j >= len(b)
}
