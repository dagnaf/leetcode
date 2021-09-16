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

type edge struct {
	a, b int
}

func rw(r io.Reader, w io.Writer) {
	var n, m1, m2 int
	fmt.Fscan(r, &n, &m1, &m2)
	e1 := make([]edge, m1)
	for i := 0; i < m1; i++ {
		fmt.Fscan(r, &e1[i].a, &e1[i].b)
	}
	e2 := make([]edge, m2)
	for i := 0; i < m2; i++ {
		fmt.Fscan(r, &e2[i].a, &e2[i].b)
	}
	w1 := do(n, e1[:m1], e2[:m2])
	fmt.Fprintln(w, len(w1))
	for _, v := range w1 {
		fmt.Fprintln(w, v.a, v.b)
	}
}

func getF(f []int, x int) int {
	if f[x] != x {
		f[x] = getF(f, f[x])
	}
	return f[x]
}

func newF(n int) []int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = i
	}
	return f
}

func newInitF(n int, e []edge) []int {
	f := newF(n)
	for _, v := range e {
		union(f, v.a, v.b)
	}
	return f
}

func union(f []int, i, j int) {
	fi, fj := getF(f, i), getF(f, j)
	f[fi] = fj
}

func same(f []int, i, j int) bool {
	return getF(f, i) == getF(f, j)
}

func do(n int, e1, e2 []edge) []edge {
	if len(e1) > len(e2) {
		e1, e2 = e2, e1
	}
	f1 := newInitF(n, e1)
	f2 := newInitF(n, e2)
	var r []edge
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if !same(f1, i, j) && !same(f2, i, j) {
				r = append(r, edge{i, j})
				union(f1, i, j)
				union(f2, i, j)
			}
		}
	}
	return r
}
