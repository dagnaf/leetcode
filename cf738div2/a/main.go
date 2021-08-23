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
	var n int
	var a = make([]int, 105)
	fmt.Fscan(r, &t)
	for ti := 0; ti < t; ti++ {
		fmt.Fscan(r, &n)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &a[i])
		}
		w1 := do(n, a[:n])
		fmt.Fprint(w, w1)
		fmt.Fprintln(w)
	}
}

func do(n int, a []int) int {
	r := a[0]
	for i := 1; i < n; i++ {
		r &= a[i]
	}
	return r
}
