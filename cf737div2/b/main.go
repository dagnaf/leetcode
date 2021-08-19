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
	var t int
	var n, k int
	a := make([]int, 1e5+5)
	fmt.Fscan(r, &t)
	for ti := 0; ti < t; ti++ {
		fmt.Fscan(r, &n, &k)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &a[i])
		}
		w1 := do(n, k, a[:n])
		if w1 {
			fmt.Fprint(w, "Yes")

		} else {
			fmt.Fprint(w, "No")

		}
		fmt.Fprintln(w)
	}
}

func do(n, k int, a []int) bool {
	b := make([]int, len(a))
	for i := range a {
		b[i] = i
	}
	sort.Slice(b, func(i, j int) bool { return a[b[i]] < a[b[j]] })
	r := 0
	for i := range b {
		if i == 0 || b[i]-b[i-1] != 1 {
			r++
		}
	}
	return r <= k
}
