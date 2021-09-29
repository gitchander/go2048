package go2048

import (
	"image"
	"math/rand"
)

func DefaultSize() image.Point {
	return image.Point{4, 4}
}

type grid struct {
	size image.Point
	sst  [][]*Tile
}

func newGrid(size image.Point) *grid {
	sst := make([][]*Tile, size.X)
	for x := range sst {
		sst[x] = make([]*Tile, size.Y)
	}
	return &grid{
		size: size,
		sst:  sst,
	}
}

func (g *grid) set(cell image.Point, t *Tile) {
	g.sst[cell.X][cell.Y] = t
}

func (g *grid) get(cell image.Point) *Tile {
	return g.sst[cell.X][cell.Y]
}

func (g *grid) Size() image.Point {
	return g.size
}

// Inserts a tile at its position
func (g *grid) insertTile(t *Tile) {
	g.set(t.Position, t)
}

func (g *grid) removeTile(t *Tile) {
	g.set(t.Position, nil)
}

// Move a tile and its representation
func (g *grid) moveTile(t *Tile, cell image.Point) {
	g.set(t.Position, nil)
	g.set(cell, t)
	t.updatePosition(cell)
}

func (g *grid) rangeCells(f func(cell image.Point, t *Tile) bool) {
	for x, st := range g.sst {
		for y, t := range st {
			cell := image.Point{x, y}
			if not(f(cell, t)) {
				return
			}
		}
	}
}

func (g *grid) rangeTiles(f func(*Tile) bool) {
	for _, st := range g.sst {
		for _, t := range st {
			if t != nil {
				if not(f(t)) {
					return
				}
			}
		}
	}
}

// Check if there are any cells available
func (g *grid) hasAvailableCells() (ok bool) {
	g.rangeCells(
		func(cell image.Point, t *Tile) bool {
			if t == nil {
				ok = true
				return false
			}
			return true
		})
	return
}

func (g *grid) availableCells() (cells []image.Point) {
	g.rangeCells(
		func(cell image.Point, t *Tile) bool {
			if t == nil {
				cells = append(cells, cell)
			}
			return true
		})
	return
}

// Check if the specified cell is taken
func (g *grid) cellAvailable(cell image.Point) bool {
	return not(g.cellOccupied(cell))
}

func (g *grid) cellOccupied(cell image.Point) bool {
	return g.cellContent(cell) != nil
}

func (g *grid) cellContent(cell image.Point) *Tile {
	if g.withinBounds(cell) {
		return g.get(cell)
	}
	return nil
}

// Use for encodePrintable
func (g *grid) CellValue(cell image.Point) (val int, ok bool) {
	if v := g.cellContent(cell); v != nil {
		return v.Value, true
	}
	return 0, false
}

func (g *grid) withinBounds(cell image.Point) bool {

	if (cell.X < 0) || (g.size.X <= cell.X) {
		return false
	}

	if (cell.Y < 0) || (g.size.Y <= cell.Y) {
		return false
	}

	return true
}

// Adds a tile in a random position
func (g *grid) addRandomTile(r *rand.Rand) {

	cells := g.availableCells()
	if (cells == nil) || (len(cells) == 0) {
		return
	}

	var value int
	if r.Float64() < 0.9 {
		value = 2
	} else {
		value = 4
	}

	cell := cells[r.Intn(len(cells))]
	g.insertTile(newTile(cell, value))
}

// Save the current tile positions and remove merger information
func (g *grid) resetTiles() {
	g.rangeTiles(
		func(t *Tile) bool {
			t.reset()
			return true
		},
	)
}

// Check for available matches between tiles (more expensive check)
func (g *grid) tileMatchesAvailable() (ok bool) {
	g.rangeCells(
		func(cell image.Point, t *Tile) bool {
			if t != nil {
				for _, d := range directions {
					vector := d.getVector()
					other := g.cellContent(cell.Add(vector))
					if (other != nil) && (other.Value == t.Value) {
						ok = true // These two tiles can be merged
						return false
					}
				}
			}
			return true
		})
	return
}

func (g *grid) movesAvailable() bool {
	return g.hasAvailableCells() || g.tileMatchesAvailable()
}

type positions struct {
	farthest, next image.Point
}

func (g *grid) findFarthestPosition(cell, vector image.Point) positions {

	var previous image.Point

	// Progress towards the vector direction until an obstacle is found
	for {
		previous = cell
		cell = previous.Add(vector)

		if not(g.withinBounds(cell) && g.cellAvailable(cell)) {
			break
		}
	}

	return positions{
		farthest: previous,
		next:     cell, // Used to check if a merge is required
	}
}
