package core

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

func (d Direction) getVector() Point {
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
	vector Point
}{
	Left: {
		name:   "left",
		vector: MakePoint(-1, 0),
	},
	Right: {
		name:   "right",
		vector: MakePoint(1, 0),
	},
	Up: {
		name:   "up",
		vector: MakePoint(0, -1),
	},
	Down: {
		name:   "down",
		vector: MakePoint(0, 1),
	},
}

//var directionVector = map[Direction]Point{
//	Left:  MakePoint(-1, 0),
//	Right: MakePoint(1, 0),
//	Up:    MakePoint(0, -1),
//	Down:  MakePoint(0, 1),
//}
