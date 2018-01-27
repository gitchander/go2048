package go2048

func serialIntSlice(n int) (d []int) {
	if n > 0 {
		d = make([]int, n)
		for i := range d {
			d[i] = i
		}
	}
	return d
}

func reverseIntSlice(d []int) {
	i, j := 0, len(d)-1
	for i < j {
		d[i], d[j] = d[j], d[i]
		i, j = i+1, j-1
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
