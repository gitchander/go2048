package core

type Tile struct {
	Position Point
	Value    int

	PreviousPosition *Point  `json:"-"`
	MergedFrom       []*Tile `json:"-"`
}

func newTile(position Point, value int) *Tile {
	return &Tile{
		Position: position,
		Value:    value,
	}
}

func (t *Tile) reset() {
	t.PreviousPosition = nil
	t.MergedFrom = nil
}

func (t *Tile) updatePosition(position Point) {
	var prevPos = t.Position
	t.PreviousPosition = &prevPos
	t.Position = position
}

func mergeTiles(pos Point, ts ...*Tile) *Tile {
	var sum int
	for _, t := range ts {
		sum += t.Value
	}
	return &Tile{
		Position:         pos,
		Value:            sum,
		PreviousPosition: nil,
		MergedFrom:       ts,
	}
}
