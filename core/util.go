package core

// serialInts returns []int{ 0, 1, 2, 3, ... n-1 }
func serialInts(n int) []int {
	as := make([]int, n)
	for i := range as {
		as[i] = i
	}
	return as
}

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a % b
	return
}

func mod(a, b int) int {
	d := a % b
	if d < 0 {
		d += b
	}
	return d
}

func not(a bool) bool {
	return !a
}
