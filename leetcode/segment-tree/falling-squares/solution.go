package falling

import "sort"

func fallingSquares(positions [][]int) []int {
	p := positions
	mpos(p)
	n := &node{l: 0, r: len(ms) - 1}
	a := []int{}
	for _, v := range p {
		l, r := ms[v[0]], ms[v[0]+v[1]]
		h := get(n, l, r-1)
		set(n, l, r-1, h+v[1])
		a = append(a, n.v)
	}
	return a
}

func get(n *node, l, r int) int {
	if n.l >= l && n.r <= r {
		return n.v
	}
	create(n)
	push(n)
	m := (n.l + n.r) / 2
	if r <= m {
		return get(n.ln, l, r)
	} else if l > m {
		return get(n.rn, l, r)
	} else {
		return max(get(n.ln, l, m), get(n.rn, m+1, r))
	}
}

func push(n *node) {
	if n.lazy > 0 {
		n.ln.lazy = n.lazy
		n.ln.v = max(n.ln.v, n.lazy)
		n.rn.lazy = n.lazy
		n.rn.v = max(n.rn.v, n.lazy)
		n.lazy = 0
	}
}

func create(n *node) {
	if n.ln == nil || n.rn == nil {
		m := (n.l + n.r) / 2
		n.ln = &node{l: n.l, r: m}
		n.rn = &node{l: m + 1, r: n.r}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func set(n *node, l, r, v int) {
	if n.l >= l && n.r <= r {
		n.lazy = v
		n.v = max(n.v, v)
		return
	}
	create(n)
	push(n)
	m := (n.l + n.r) / 2
	if r <= m {
		set(n.ln, l, r, v)
	} else if l > m {
		set(n.rn, l, r, v)
	} else {
		set(n.ln, l, m, v)
		set(n.rn, m+1, r, v)
	}
	n.v = max(n.ln.v, n.rn.v)
}

var (
	ms map[int]int
	ks []int
)

func mpos(p [][]int) {
	ms = map[int]int{}
	for _, v := range p {
		ms[v[0]] = 0
		ms[v[0]+v[1]] = 0
	}
	keys := make([]int, 0, len(ms))
	for k := range ms {
		keys = append(keys, k)
	}
	ks = keys
	sort.Ints(ks)
	for i, v := range ks {
		ms[v] = i
	}
}

type node struct {
	l, r   int
	ln, rn *node
	lazy   int
	v      int
}
