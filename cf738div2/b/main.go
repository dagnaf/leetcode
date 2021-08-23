package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var _, _, _ = main, rw, do

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	rw(r, w)
}

func rw(r io.Reader, w io.Writer) {
	var t int
	var n int
	var s string
	fmt.Fscan(r, &t)
	for ti := 0; ti < t; ti++ {
		fmt.Fscan(r, &n, &s)
		w1 := do(n, s)
		if ti != 0 {
			fmt.Fprintln(w)
		}
		fmt.Fprint(w, w1)
	}
}

var m = map[byte]byte{'R': 'B', 'B': 'R'}

func do(n int, s string) string {
	t := []byte(s)
	i := 0
	for i < n && t[i] == '?' {
		i++
	}
	if i == n {
		return strings.Repeat("BR", (n+1)/2)[:n]
	}
	for ; i >= 1; i-- {
		t[i-1] = m[t[i]]
	}
	for i = 1; i < n; i++ {
		if t[i] == '?' {
			t[i] = m[t[i-1]]
		}
	}
	return string(t)
}
