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
	var a = make([]int, 1e4+2)
	fmt.Fscan(r, &t)
	for ti := 0; ti < t; ti++ {
		fmt.Fscan(r, &n)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &a[i])
		}
		w1 := do(n, a[:n])
		if len(w1) == 0 {
			fmt.Fprint(w, -1)
		} else {
			for i, v := range w1 {
				if i > 0 {
					fmt.Fprint(w, " ")
				}
				fmt.Fprint(w, v)
			}
		}
		fmt.Fprintln(w)
	}
}

func arr(a, b int) []int {
	r := make([]int, 0, b-a+1)
	for i := a; i <= b; i++ {
		r = append(r, i)
	}
	return r
}

func do(n int, a []int) []int {
	if a[0] == 1 {
		return append([]int{n + 1}, arr(1, n)...)
	}
	if a[n-1] == 0 {
		return arr(1, n+1)
	}
	for i := 0; i < n-1; i++ {
		if a[i] == 0 && a[i+1] == 1 {
			return append(append(arr(1, i+1), n+1), arr(i+2, n)...)
		}
	}
	return nil
}
