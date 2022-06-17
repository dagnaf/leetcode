package online

type MajorityChecker struct {
	ns []*node
}

const sz = 2 * 1e4

func Constructor(arr []int) MajorityChecker {
	c := MajorityChecker{}
	c.ns = make([]*node, len(arr)+1)
	c.ns[0] = &node{l: 1, r: sz}
	for i, v := range arr {
		c.ns[i+1] = add(c.ns[i], v)
	}
	return c
}

func (this *MajorityChecker) Query(left int, right int, threshold int) int {
	c := this
	return get(c.ns[left], c.ns[right+1], threshold)
}

type node struct {
	l, r   int
	ln, rn *node
	k      int
	v      int
}

func get(n *node, n2 *node, v int) int {
	if n.l == n.r {
		if n2.v-n.v >= v {
			return n.k
		}
		return -1
	}
	create(n)
	if n2.ln.v-n.ln.v >= v {
		return get(n.ln, n2.ln, v)
	} else if n2.rn.v-n.rn.v >= v {
		return get(n.rn, n2.rn, v)
	}
	return -1
}

func add(n *node, k int) *node {
	n2 := &node{}
	*n2 = *n
	if n.l == n.r {
		n.k = k
		n2.k = k
		n2.v++
		return n2
	}
	create(n)
	m := (n.l + n.r) / 2
	if k <= m {
		n2.ln = add(n.ln, k)
		n2.rn = n.rn
	} else {
		n2.ln = n.ln
		n2.rn = add(n.rn, k)
	}
	n2.v = n2.ln.v + n2.rn.v
	return n2
}

func create(n *node) {
	if n.ln == nil || n.rn == nil {
		m := (n.l + n.r) / 2
		n.ln = &node{l: n.l, r: m}
		n.rn = &node{l: m + 1, r: n.r}
	}
}

/**
 * Your MajorityChecker object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,threshold);
 */

func run(cmd []string, arg [][]int) []interface{} {
	c := MajorityChecker{}
	got := make([]interface{}, len(cmd))
	for i := range cmd {
		switch cmd[i] {
		case "MajorityChecker":
			c = Constructor(arg[i])
		case "query":
			got[i] = float64(c.Query(arg[i][0], arg[i][1], arg[i][2]))
		}
	}
	return got
}

/*
给定数组，在线查询区间中是否存在大于区间半数频数的数
可持久化线段树、主席树
数组前缀0-i建立线段树，维护频数直方图
查询l-r的频数即两颗树相减
查询频数大于半数，只需要查询一侧子树
*/
