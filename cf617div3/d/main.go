package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

var _, _, _ = main, rw, do

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	rw(r, w)
}

func rw(r io.Reader, w io.Writer) {
	var n, a, b, k int
	h := make([]int, 2e5+1)
	fmt.Fscanf(r, "%d", &n)
	fmt.Fscanf(r, "%d", &a)
	fmt.Fscanf(r, "%d", &b)
	fmt.Fscanf(r, "%d", &k)
	fmt.Fscanln(r)
	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%d", &h[i])
	}
	w1 := do(n, a, b, k, h[:n])
	fmt.Fprintf(w, "%d", w1)
}

func do(n, a, b, k int, h []int) int {
	for i, hi := range h {
		h[i] = getK(a, b, hi)
	}
	sort.Ints(h)
	for i, hi := range h {
		k -= hi
		if k < 0 {
			return i
		}
	}
	return n
}

func getK(a, b, hi int) int {
	hi -= a
	if hi <= 0 {
		return 0
	}
	res := hi % (b + a)
	if res == 0 {
		return 0
	}
	if b < res {
		return 0
	}
	return (res + a - 1) / a
}
