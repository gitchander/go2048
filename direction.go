package go2048

import "image"

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

func (d Direction) String() string {
	return dirs[d].name
}

func (d Direction) getVector() image.Point {
	return dirs[d].vector
}

var directions = [...]Direction{
	Left,
	Right,
	Up,
	Down,
}

var dirs = [...]struct {
	name   string
	vector image.Point
}{
	Left: {
		name:   "left",
		vector: image.Point{-1, 0},
	},
	Right: {
		name:   "right",
		vector: image.Point{1, 0},
	},
	Up: {
		name:   "up",
		vector: image.Point{0, -1},
	},
	Down: {
		name:   "down",
		vector: image.Point{0, 1},
	},
}

//var directionVector = map[Direction]image.Point{
//	Left:  image.Point{-1, 0},
//	Right: image.Point{1, 0},
//	Up:    image.Point{0, -1},
//	Down:  image.Point{0, 1},
//}
