package go2048

import "image"

type Tile struct {
	Position image.Point
	Value    int

	PreviousPosition *image.Point `json:"-"`
	MergedFrom       []*Tile      `json:"-"`
}

func newTile(position image.Point, value int) *Tile {
	return &Tile{
		Position: position,
		Value:    value,
	}
}

func (t *Tile) resetPrevious() {
	t.PreviousPosition = nil
	t.MergedFrom = nil
}

func (t *Tile) updatePosition(position image.Point) {
	var prevPos = t.Position
	t.PreviousPosition = &prevPos
	t.Position = position
}

func mergeTiles(pos image.Point, ts ...*Tile) *Tile {
	if len(ts) != 2 {
		panic("merged not two tiles")
	}
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
