package go2048

import "image"

var DefaultSize = image.Point{4, 4}

type grid struct {
	tiles [][]*Tile
}

func tilesFromSize(size image.Point) [][]*Tile {
	tiles := make([][]*Tile, size.X)
	for x := range tiles {
		tiles[x] = make([]*Tile, size.Y)
	}
	return tiles
}

func newGrid(size image.Point) *grid {
	return &grid{
		tiles: tilesFromSize(size),
	}
}

func (g *grid) set(cell image.Point, t *Tile) {
	g.tiles[cell.X][cell.Y] = t
}

func (g *grid) get(cell image.Point) *Tile {
	return g.tiles[cell.X][cell.Y]
}

func (g *grid) Size() image.Point {
	var size = image.Point{
		X: len(g.tiles),
		Y: 0,
	}
	for _, ts := range g.tiles {
		size.Y = max(size.Y, len(ts))
	}
	return size

	//	var (
	//		nX = len(g.tiles)
	//		nY int
	//	)
	//	for _, ts := range g.tiles {
	//		nY = max(nY, len(ts))
	//	}
	//	return image.Point{
	//		X: nX,
	//		Y: nY,
	//	}
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

func (g *grid) availableCells() (cells []image.Point) {
	for x, ts := range g.tiles {
		for y, t := range ts {
			if t == nil {
				cells = append(cells, image.Point{x, y})
			}
		}
	}
	return
}

func (g *grid) forEach(fn func(*Tile)) {
	for _, ts := range g.tiles {
		for _, t := range ts {
			if t != nil {
				fn(t)
			}
		}
	}
}

func (g *grid) forEachCell(fn func(image.Point, *Tile)) {
	for x, ts := range g.tiles {
		for y, t := range ts {
			var cell = image.Point{x, y}
			fn(cell, t)
		}
	}
}

// Check if there are any cells available
func (g *grid) cellsAvailable() bool {
	for _, ts := range g.tiles {
		for _, t := range ts {
			if t == nil {
				return true
			}
		}
	}
	return false
}

// Check if the specified cell is taken
func (g *grid) cellAvailable(cell image.Point) bool {
	return !g.cellOccupied(cell)
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

func (g *grid) withinBounds(cell image.Point) bool {

	nX := len(g.tiles)
	if (cell.X < 0) || (cell.X >= nX) {
		return false
	}

	nY := len(g.tiles[cell.X])
	if (cell.Y < 0) || (cell.Y >= nY) {
		return false
	}

	return true
}

// Adds a tile in a random position
func (g *grid) addRandomTile() {

	cells := g.availableCells()
	if (cells == nil) || (len(cells) == 0) {
		return
	}

	var value int
	if random() < 0.9 {
		value = 2
	} else {
		value = 4
	}

	cell := cells[randIntn(len(cells))]
	g.insertTile(newTile(cell, value))
}

// Save the current tile positions and remove merger information
func (g *grid) prepareTiles() {
	g.forEach(
		func(t *Tile) {
			t.savePosition()
		},
	)
}

// Check for available matches between tiles (more expensive check)
func (g *grid) tileMatchesAvailable() bool {
	for x, ts := range g.tiles {
		for y, t := range ts {
			if t != nil {
				var cell = image.Point{x, y}
				for _, d := range directions {
					vector := directionVector[d]
					other := g.cellContent(cell.Add(vector))
					if (other != nil) && (other.Value == t.Value) {
						return true // These two tiles can be merged
					}
				}
			}
		}
	}
	return false
}

func (g *grid) movesAvailable() bool {
	return g.cellsAvailable() || g.tileMatchesAvailable()
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

		if !(g.withinBounds(cell) && g.cellAvailable(cell)) {
			break
		}
	}

	return positions{
		farthest: previous,
		next:     cell, // Used to check if a merge is required
	}
}
