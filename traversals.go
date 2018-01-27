package go2048

import "image"

type traversals struct {
	xs, ys []int
}

func (t traversals) forEach(fn func(image.Point)) {
	for _, x := range t.xs {
		for _, y := range t.ys {
			fn(image.Point{x, y})
		}
	}
}

// Build a list of positions to traverse in the right order
func buildTraversals(size, vector image.Point) traversals {

	var t = traversals{
		xs: serialIntSlice(size.X),
		ys: serialIntSlice(size.Y),
	}

	// Always traverse from the farthest cell in the chosen direction
	if vector.X == 1 {
		reverseIntSlice(t.xs)
	}
	if vector.Y == 1 {
		reverseIntSlice(t.ys)
	}

	return t
}

func buildMapTraversals(size image.Point) map[Direction]traversals {
	m := make(map[Direction]traversals)
	for _, d := range directions {
		vector := directionVector[d]
		m[d] = buildTraversals(size, vector)
	}
	return m
}
