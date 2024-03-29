package core

import (
	"github.com/gitchander/go2048/utils"
)

// traverse

type traversals struct {
	xs, ys []int
}

// Build a list of positions to traverse in the right order
func newTraversals(size Point, d Direction) *traversals {

	t := &traversals{
		xs: serialInts(size.X),
		ys: serialInts(size.Y),
	}

	vector := d.getVector()

	// Always traverse from the farthest cell in the chosen direction
	if vector.X == 1 {
		utils.Reverse(utils.IntSlice(t.xs))
	}
	if vector.Y == 1 {
		utils.Reverse(utils.IntSlice(t.ys))
	}

	return t
}

func (t *traversals) Range(f func(Point) bool) {
	for _, x := range t.xs {
		for _, y := range t.ys {
			if not(f(MakePoint(x, y))) {
				return
			}
		}
	}
}

func makeMapTraversals(size Point) map[Direction]*traversals {
	m := make(map[Direction]*traversals)
	for _, d := range directions {
		m[d] = newTraversals(size, d)
	}
	return m
}
