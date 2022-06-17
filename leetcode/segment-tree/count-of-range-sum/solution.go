package count

import "sort"

func countRangeSum(nums []int, lower int, upper int) int {
	l, r := lower, upper
	p := make([]int64, len(nums)+1)
	for i := 1; i < len(p); i++ {
		p[i] += p[i-1] + int64(nums[i-1])
	}
	mpos(p, l, r)
	n := &node{l: 0, r: len(ms) - 1}
	a := 0
	for _, v := range p {
		a += get(n, ms[v-int64(r)], ms[v-int64(l)])
		add(n, ms[v])
	}
	return a
}

func add(n *node, k int) {
	if n.l == n.r {
		n.v++
		return
	}
	create(n)
	m := (n.l + n.r) / 2
	if k <= m {
		add(n.ln, k)
	} else {
		add(n.rn, k)
	}
	n.v = n.ln.v + n.rn.v
}

func get(n *node, l, r int) int {
	if n.l >= l && n.r <= r {
		return n.v
	}
	create(n)
	m := (n.l + n.r) / 2
	if r <= m {
		return get(n.ln, l, r)
	} else if l > m {
		return get(n.rn, l, r)
	} else {
		return get(n.ln, l, m) + get(n.rn, m+1, r)
	}
}

func create(n *node) {
	if n.ln == nil || n.rn == nil {
		m := (n.l + n.r) / 2
		n.ln = &node{l: n.l, r: m}
		n.rn = &node{l: m + 1, r: n.r}
	}
}

var (
	ms map[int64]int
	ks []int64
)

func mpos(p []int64, l, r int) {
	ms = map[int64]int{}
	for _, v := range p {
		ms[v] = 0
		ms[v-int64(l)] = 0
		ms[v-int64(r)] = 0
	}
	keys := make([]int64, 0, len(ms))
	for k := range ms {
		keys = append(keys, k)
	}
	ks = keys
	sort.Slice(ks, func(i, j int) bool { return ks[i] < ks[j] })
	for i, v := range ks {
		ms[v] = i
	}
}

type node struct {
	l, r   int
	ln, rn *node
	v      int
}
