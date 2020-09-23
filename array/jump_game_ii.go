package array

func jump(nums []int) int {
	n := len(nums)
	d := make([]int, n)
	a := make([]int, n)
	k := 0
	a[0] = n - 1
	for i := n - 2; i >= 0; i-- {
		x := nums[i] + i
		if x > n - 1 {
			x = n - 1
		}
		l, r := 0, k
		for l < r {
			m := (l+r)/2
			if a[m] <= x {
				r = m
			} else {
				l = m + 1
			}
		}
		k = l+1
		a[k] = i
		d[i] = d[a[l]] + 1
	}
	return d[0]
}

// todo best o(n)