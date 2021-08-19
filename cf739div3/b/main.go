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
	var t, a, b, c int
	fmt.Fscan(r, &t)
	for ti := 0; ti < t; ti++ {
		fmt.Fscan(r, &a, &b, &c)
		w1 := do(a, b, c)
		fmt.Fprint(w, w1)
		fmt.Fprintln(w)
	}
}

func do(a, b, c int) int {
	n := a - b
	if n < 0 {
		n = -n
	}
	if a > n+n || b > n+n || c > n+n {
		return -1
	}
	if v := n + c; 1 <= v && v <= n+n {
		return v
	}
	if v := -n + c; 1 <= v && v <= n+n {
		return v
	}
	return -1
}
