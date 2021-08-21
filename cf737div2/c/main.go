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
	var n, k int
	fmt.Fscan(r, &t)
	for ti := 0; ti < t; ti++ {
		fmt.Fscan(r, &n, &k)
		w1 := do(n, k)
		fmt.Fprint(w, w1)
		fmt.Fprintln(w)
	}
}

func do(n, k int) int64 {
	if k == 0 {
		return 1
	}
	tk := getTk(n, k)
	gr, eq := getGrEq(n)
	r := gr + eq
	for i := 2; i <= k; i++ {
		r = (((eq * r) % mod) + ((gr * tk[i-1]) % mod)) % mod
	}
	return r
}

const mod = 1e9 + 7

func getTn(n int) int64 {
	var tn int64 = 1
	for i := 1; i <= n; i++ {
		tn = (tn * 2) % mod
	}
	return tn
}

func getTk(n, k int) []int64 {
	tk := make([]int64, k+1)
	tk[0] = 1
	tn := getTn(n)
	for i := 1; i <= k; i++ {
		tk[i] = (tk[i-1] * tn) % mod
	}
	return tk
}

// getGrEq
// c[n,k]=c[n-1,k-1]+c[n-1,k]
// c[n,0]+c[n,2]+...=c[n-1,0]+c[n-1,1]+c[n-1,2]+...=2^n...
func getGrEq(n int) (int64, int64) {
	if n == 1 {
		return 0, 2
	}
	if n == 2 {
		return 1, 1
	}
	tn := getTn(n - 1)
	var gr, eq int64 = 0, tn
	if n%2 == 0 {
		gr = 1
		eq = (eq + mod - 1) % mod
	} else {
		eq = (eq + 1) % mod
	}
	return gr, eq
}
