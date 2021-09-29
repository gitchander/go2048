package utils

import (
	"sort"
)

func _() {
	var _ sort.Interface

	// type Interface interface {
	// 	// Len is the number of elements in the collection.
	// 	Len() int
	// 	Less(i, j int) bool
	// 	// Swap swaps the elements with indexes i and j.
	// 	Swap(i, j int)
	// }
}

type Swapper interface {
	// Len is the number of elements in the collection.
	Len() int

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

func Reverse(s Swapper) {
	i, j := 0, (s.Len() - 1)
	for i < j {
		s.Swap(i, j)
		i, j = i+1, j-1
	}
}

type IntSlice []int

var _ Swapper = IntSlice(nil)

func (p IntSlice) Len() int      { return len(p) }
func (p IntSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
