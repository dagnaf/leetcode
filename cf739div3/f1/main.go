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
	var k int
	fmt.Fscan(r, &t)
	for ti := 0; ti < t; ti++ {
		fmt.Fscan(r, &s, &k)
		w1 := do(s, k)
		fmt.Fprint(w, w1)
		fmt.Fprintln(w)
	}
}

func do(t string, k int) string {
	s := make([]byte, len(t))
	for i := range t {
		s[i] = t[i] - '0'
	}
	v := make([]byte, 10)
	w := make([]byte, len(t))
	dfs(s, w, v, k, true)
	for i := range w {
		w[i] += '0'
	}
	return string(w)
}

func dfs(t []byte, w []byte, v []byte, k int, strict bool) bool {
	if k < 0 {
		return false
	}
	if len(t) == 0 {
		return true
	}
	var b byte
	if strict {
		b = t[0]
	}
	for i := b; i <= 9; i++ {
		w[0] = i
		o := 0
		if v[i] == 0 {
			o = 1
		}
		v[i]++
		if dfs(t[1:], w[1:], v, k-o, strict && i == t[0]) {
			return true
		}
		v[i]--
	}
	return false
}
