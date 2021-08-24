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
	var a = make([]int, 1e3+1)
	fmt.Fscan(r, &t)
	for ti := 0; ti < t; ti++ {
		fmt.Fscan(r, &n)
		for i := 1; i <= n; i++ {
			fmt.Fscan(r, &a[i])
		}
		w1 := do(n, a)
		fmt.Fprintln(w, w1)
	}
}

func sorted(n int, a []int) bool {
	for i := 1; i <= n; i++ {
		if a[i] != i {
			return false
		}
	}
	return true
}

func do(n int, a []int) int {
	r := 0
	for !sorted(n, a) {
		r++
		k1, k2 := 1, 2
		if r%2 == 0 {
			k1, k2 = 2, 1
		}
		for i := k1; i <= n-k2; i += 2 {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
	}
	return r
}
