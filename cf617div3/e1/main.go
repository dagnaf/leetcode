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
	var n int
	var s string
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &s)
	w1, w2 := do(n, []byte(s))
	if w2 {
		fmt.Fprintln(w, "YES")
		fmt.Fprintln(w, string(w1))
	} else {
		fmt.Fprintln(w, "NO")
	}
}

func do(n int, s []byte) ([]byte, bool) {
	st := getSt(n, s)
	re := getRe(n, st)
	g := map[int][]int{}
	for _, e := range re {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	vis := map[int]int{}
	for i := 0; i < n; i++ {
		if _, ok := vis[i]; !ok {
			if !color(i, g, vis, 0) {
				return nil, false
			}
		}
	}
	c := make([]byte, n)
	for i, v := range vis {
		c[i] = byte('0' + v)
	}
	return c, true
}

func color(cur int, g map[int][]int, vis map[int]int, co int) bool {
	v, ok := vis[cur]
	if ok {
		return v == co
	}
	vis[cur] = co
	for _, next := range g[cur] {
		if !color(next, g, vis, 1-co) {
			return false
		}
	}
	return true
}

func getRe(n int, st st) (re [][2]int) {
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if st.i[j] > st.i[i] {
				re = append(re, [2]int{st.i[i], st.i[j]})
			}
		}
	}
	return
}

func getSt(n int, s []byte) st {
	st := st{s: s, i: make([]int, n)}
	for i := range s {
		st.i[i] = i
	}
	sort.Stable(st)
	return st
}

type st struct {
	s []byte
	i []int
}

func (st st) Len() int {
	return len(st.s)
}

func (st st) Less(i, j int) bool {
	if st.s[i] == st.s[j] {
		return st.i[i] < st.i[j]
	}
	return st.s[i] < st.s[j]
}

func (st st) Swap(i, j int) {
	st.s[i], st.s[j] = st.s[j], st.s[i]
	st.i[i], st.i[j] = st.i[j], st.i[i]
}
