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
	var n int
	fmt.Fscan(r, &t)
	for ti := 0; ti < t; ti++ {
		fmt.Fscan(r, &n)
		a := make([][]int, n)
		for i := 0; i < n; i++ {
			var k int
			fmt.Fscan(r, &k)
			a[i] = make([]int, k)
			for j := 0; j < k; j++ {
				fmt.Fscan(r, &a[i][j])
			}
		}
		w1 := do(a)
		fmt.Fprintln(w, w1)
	}
}

func minPower(a []cave) int {
	acc := make([]int, len(a))
	for i := 1; i < len(a); i++ {
		acc[i] = acc[i-1] + a[i-1].add
	}
	min, cur := 0, 0
	for i := range a {
		if cur < a[i].min {
			cur = a[i].min
			min = cur - acc[i]
		}
		cur += a[i].add
	}
	return min
}

type cave struct {
	min int
	add int
}

func do(a [][]int) int {
	b := make([]cave, len(a))
	for i := 0; i < len(a); i++ {
		c := make([]cave, len(a[i]))
		for j := range c {
			c[j].min = a[i][j] + 1
			c[j].add = 1
		}
		b[i].min = minPower(c)
		b[i].add = len(a[i])
	}
	sort.Slice(b, func(i, j int) bool {
		if b[i].min == b[j].min {
			return b[i].add > b[j].add
		}
		return b[i].min < b[j].min
	})
	return minPower(b)
}
