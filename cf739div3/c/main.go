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
	var n int32
	fmt.Fscan(r, &t)
	for ti := 0; ti < t; ti++ {
		fmt.Fscan(r, &n)
		w1, w2 := do(n)
		fmt.Fprint(w, w1, w2)
		fmt.Fprintln(w)
	}
}

func do(n int32) (int32, int32) {
	r := int32(0)
	for {
		if r*r < n && n <= (r+1)*(r+1) {
			r += 1
			break
		}
		r++
	}
	a := (r-1)*(r-1) + 1
	b := r * r
	c := (a + b) / 2
	if n < c {
		return (n - a) + 1, r
	}
	return r, (b - n) + 1
}
