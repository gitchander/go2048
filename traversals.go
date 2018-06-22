package go2048

import "image"

// traverse

type traversals struct {
	xs, ys []int
}

// Build a list of positions to traverse in the right order
func newTraversals(size image.Point, d Direction) *traversals {

	t := &traversals{
		xs: serialInts(size.X),
		ys: serialInts(size.Y),
	}

	vector := d.getVector()

	// Always traverse from the farthest cell in the chosen direction
	if vector.X == 1 {
		reverseInts(t.xs)
	}
	if vector.Y == 1 {
		reverseInts(t.ys)
	}

	return t
}

func (t *traversals) Range(f func(image.Point)) {
	for _, x := range t.xs {
		for _, y := range t.ys {
			f(image.Pt(x, y))
		}
	}
}

func makeMapTraversals(size image.Point) map[Direction]*traversals {
	m := make(map[Direction]*traversals)
	for _, d := range directions {
		m[d] = newTraversals(size, d)
	}
	return m
}
