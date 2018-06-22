package go2048

// []int{ 0, 1, 2, 3, ... n-1 }
func serialInts(n int) (d []int) {
	if n > 0 {
		d = make([]int, n)
		for i := range d {
			d[i] = i
		}
	}
	return d
}

func reverseInts(d []int) {
	i, j := 0, len(d)-1
	for i < j {
		d[i], d[j] = d[j], d[i]
		i, j = i+1, j-1
	}
}

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a - quo*b
	return
}

func mod(a, b int) int {
	d := a % b
	if d < 0 {
		d += b
	}
	return d
}
