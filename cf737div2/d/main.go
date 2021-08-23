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

type seg struct {
	l, r int
}

func rw(r io.Reader, w io.Writer) {
	var n, m int
	fmt.Fscan(r, &n, &m)
	if n <= 0 {
		return
	}
	a := make([][]seg, n+1)
	for i := 0; i < m; i++ {
		var b seg
		var bi int
		fmt.Fscan(r, &bi, &b.l, &b.r)
		a[bi] = append(a[bi], b)
	}
	w1 := do(n, a)
	fmt.Fprint(w, len(w1))
	fmt.Fprintln(w)
	for i, v := range w1 {
		if i != 0 {
			fmt.Fprint(w, " ")
		}
		fmt.Fprint(w, v)
	}
}

func do(n int, a [][]seg) []int {
	a, l := compress(a)
	t := newSegTree(seg{1, l})
	trace := make([]int, n+1)
	retVal := segVal{}
	for i := 1; i <= n; i++ {
		curVal := segVal{max: 1}
		for _, s := range a[i] {
			if segVal := t.get(s); segVal.max+1 > curVal.max {
				curVal = segVal
				curVal.max++
			}
		}
		trace[i] = curVal.i
		curVal.i = i
		for _, s := range a[i] {
			t.upd(s, curVal)
		}
		if curVal.max > retVal.max {
			retVal = curVal
		}
	}
	return newRet(n, trace, retVal)
}

func newRet(n int, trace []int, retVal segVal) []int {
	vis := make([]bool, n+1)
	for retVal.i != 0 {
		vis[retVal.i] = true
		retVal.i = trace[retVal.i]
	}
	ret := []int{}
	for i := 1; i <= n; i++ {
		if !vis[i] {
			ret = append(ret, i)
		}
	}
	return ret
}

func (node *segNode) upd(s seg, v segVal) {
	if s.l == node.seg.l && s.r == node.seg.r {
		node.v = v
		node.lazy = &node.v
		return
	}
	node.pushDown()
	m := (node.seg.l + node.seg.r) / 2
	if s.r <= m {
		node.l.upd(s, v)
	} else if s.l > m {
		node.r.upd(s, v)
	} else {
		node.l.upd(seg{s.l, m}, v)
		node.r.upd(seg{m + 1, s.r}, v)
	}
	node.pushUp()
}

func (node *segNode) pushDown() {
	if node.lazy == nil {
		return
	}
	for _, child := range []*segNode{node.l, node.r} {
		if child != nil {
			child.v = *node.lazy
			child.lazy = &child.v
		}
	}
	node.lazy = nil
}

func (node *segNode) pushUp() {
	for _, child := range []*segNode{node.l, node.r} {
		if child != nil && child.v.max > node.v.max {
			node.v = child.v
		}
	}
}

func (node *segNode) get(s seg) segVal {
	if s.l == node.seg.l && s.r == node.seg.r {
		return node.v
	}
	node.pushDown()
	var lv, rv segVal
	m := (node.seg.l + node.seg.r) / 2
	if s.r <= m {
		lv = node.l.get(s)
	} else if s.l > m {
		rv = node.r.get(s)
	} else {
		lv = node.l.get(seg{s.l, m})
		rv = node.r.get(seg{m + 1, s.r})
	}
	node.pushUp()
	if lv.max >= rv.max {
		return lv
	}
	return rv
}

type segVal struct {
	max, i int
}

type segNode struct {
	seg  seg
	l, r *segNode
	v    segVal
	lazy *segVal
}

func newSegTree(s seg) *segNode {
	node := &segNode{}
	node.seg = s
	if s.l == s.r {
		return node
	}
	m := (s.l + s.r) / 2
	node.l = newSegTree(seg{s.l, m})
	node.r = newSegTree(seg{m + 1, s.r})
	return node
}

func compress(a [][]seg) ([][]seg, int) {
	m := map[int]int{}
	for _, segs := range a {
		for i := range segs {
			m[segs[i].l] = 0
			m[segs[i].r] = 0
		}
	}
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for i, k := range keys {
		m[k] = i + 1
	}
	for _, segs := range a {
		for i := range segs {
			segs[i].l = m[segs[i].l]
			segs[i].r = m[segs[i].r]
		}
	}
	return a, len(keys)
}
