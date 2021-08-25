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
	var n int
	var m int64
	fmt.Fscan(r, &n, &m)
	w1 := do(n, m)
	fmt.Fprintln(w, w1)
}

/*x/2 x/3      ... x/n/2 ... x/n
1 2,3 3,4,5    ... n/2   ... n
2 4,5 6,7,8    ... n
3 6,7 9,10,11  ...
4 8,9 12,13,14 ...
. ...
n+n/2+n/3+...+n/n-1=n(1+1/2+1/3+...)=nlogn
*/

func do(n int, m int64) int64 {
	a := make([]int64, n+2)
	a[n] = 1
	for i := n - 1; i >= 1; i-- {
		a[i] = a[i+1]
		for j := 2; j*i <= n; j++ {
			a[i] = (a[i] + a[j*i] - a[min(n+1, j*i+j)] + m) % m
		}
		a[i] = (a[i] + a[i+1]) % m
	}
	return (a[1] - a[2] + m) % m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
