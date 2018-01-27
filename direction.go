package go2048

import "image"

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

var directions = []Direction{
	Left,
	Right,
	Up,
	Down,
}

var directionVector = map[Direction]image.Point{
	Left:  image.Point{-1, 0},
	Right: image.Point{1, 0},
	Up:    image.Point{0, -1},
	Down:  image.Point{0, 1},
}
