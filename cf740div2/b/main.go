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
	var a, b int
	fmt.Fscan(r, &t)
	for ti := 0; ti < t; ti++ {
		fmt.Fscan(r, &a, &b)
		w1 := do(a, b)
		fmt.Fprintln(w, len(w1))
		for i, v := range w1 {
			if i > 0 {
				fmt.Fprint(w, " ")
			}
			fmt.Fprint(w, v)
			if i == len(w1)-1 {
				fmt.Fprintln(w)
			}
		}
	}
}

func do(a, b int) []int {
	m := map[int]struct{}{}
	m = g(a, b, m)
	m = g(b, a, m)
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func g(a, b int, m map[int]struct{}) map[int]struct{} {
	n1 := (a + b + 1) / 2
	n2 := a + b - n1
	for b1 := 0; b1 <= n2 && b1 <= a; b1++ {
		s1 := a - b1
		b2 := n1 - s1
		s2 := b - b2
		if s1 >= 0 && b2 >= 0 && s2 >= 0 {
			m[b1+b2] = struct{}{}
		}
	}
	return m
}
