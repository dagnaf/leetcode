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

func getP(f []int, n, j int) []int {
	var p []int
	for i := 1; i <= n; i++ {
		if !same(f, i, j) {
			p = append(p, i)
		}
	}
	return p
}

func do(n int, e1, e2 []edge) []edge {
	return do2(n, e1, e2)
}

var _, _ = do1, do2

func do1(n int, e1, e2 []edge) []edge {
	f1 := newInitF(n, e1)
	f2 := newInitF(n, e2)
	var r []edge
	for i := 2; i <= n; i++ {
		if !same(f1, 1, i) && !same(f2, 1, i) {
			r = append(r, edge{1, i})
			union(f1, 1, i)
			union(f2, 1, i)
		}
	}
	p1 := getP(f1, n, 1)
	p2 := getP(f2, n, 1)

	for len(p1) > 0 && len(p2) > 0 {
		if same(f1, p1[0], 1) && same(f2, p1[0], 1) {
			p1 = p1[1:]
			continue
		}
		if same(f1, p2[0], 1) && same(f2, p2[0], 1) {
			p2 = p2[1:]
			continue
		}
		r = append(r, edge{p1[0], p2[0]})
		union(f1, p1[0], p2[0])
		union(f2, p1[0], p2[0])
	}
	return r
}

func do2(n int, e1, e2 []edge) []edge {
	if len(e1) < len(e2) {
		e1, e2 = e2, e1
	}
	f1 := newInitF(n, e1)
	f2 := newInitF(n, e2)
	row := map[int]map[int]struct{}{}
	col := map[int]map[int]struct{}{}
	rep := map[int]map[int]int{}
	for i := 1; i <= n; i++ {
		x, y := getF(f1, i), getF(f2, i)
		if row[x] == nil {
			row[x] = map[int]struct{}{}
		}
		row[x][y] = struct{}{}
		if col[y] == nil {
			col[y] = map[int]struct{}{}
		}
		col[y][x] = struct{}{}
		if rep[x] == nil {
			rep[x] = map[int]int{}
		}
		rep[x][y] = i
	}

	x := make([]int, 0, len(row))
	for k := range row {
		x = append(x, k)
	}
	for i, xi := range x {
		if len(row[xi]) > 1 {
			x[0], x[i] = x[i], x[0]
			break
		}
	}

	var ans []edge
	x0 := x[0]
	for _, x1 := range x[1:] {
		var y0, y1 int
		for y0 = range row[x0] {
			break
		}
		for y1 = range row[x1] {
			if y0 != y1 {
				break
			}
		}

		ans = append(ans, edge{rep[x0][y0], rep[x1][y1]})

		if len(row[x0]) < len(row[x1]) {
			x0, x1 = x1, x0
		}

		if len(col[y0]) < len(col[y1]) {
			y0, y1 = y1, y0
		}

		mergeRow(row, col, x0, x1, rep)
		mergeCol(row, col, y0, y1, rep)
	}
	return ans
}

func mergeRow(row, col map[int]map[int]struct{}, x0, x1 int, rep map[int]map[int]int) {
	for y := range row[x1] {
		row[x0][y] = struct{}{}
		rep[x0][y] = rep[x1][y]
		delete(col[y], x1)
		col[y][x0] = struct{}{}
	}
}

func mergeCol(row, col map[int]map[int]struct{}, y0, y1 int, rep map[int]map[int]int) {
	for x := range col[y1] {
		col[y0][x] = struct{}{}
		rep[x][y0] = rep[x][y1]
		delete(row[x], y1)
		row[x][y0] = struct{}{}
	}
}
