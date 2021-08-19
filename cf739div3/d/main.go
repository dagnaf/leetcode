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
	var n int32
	fmt.Fscan(r, &t)
	for ti := 0; ti < t; ti++ {
		fmt.Fscan(r, &n)
		w1 := do(n)
		fmt.Fprint(w, w1)
		fmt.Fprintln(w)
	}
}

func do(n int32) int {
	s := fmt.Sprintf("%d", n)
	r := -1
	for _, t := range a {
		d := dist(t, s)
		if r == -1 || r > d {
			r = d
		}
	}
	return r
}

func dist(a, b string) int {
	i := 0
	j := 0
	r := 0
	for {
		for j < len(b) && a[i] != b[j] {
			j++
			r++
		}
		if j == len(b) {
			r += len(a) - i
			break
		}
		i++
		j++
		if i == len(a) {
			r += len(b) - j
			break
		}
	}
	return r
}

var a []string

func init() {
	i := uint64(1)
	for k := 1; k <= 64; k++ {
		a = append(a, fmt.Sprintf("%d", i))
		i *= 2
	}
}
