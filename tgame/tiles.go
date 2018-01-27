package main

import (
	"image"

	"github.com/nsf/termbox-go"

	game "github.com/gitchander/go2048"
)

type tilesDrawer struct {
	screenSize image.Point
	gridSize   image.Point
	cellSize   image.Point
	fg, bg     termbox.Attribute
	valueRunes map[int][]rune
	tiles      []game.Tile
}

func (td *tilesDrawer) SetScreenSize(size image.Point) {
	td.screenSize = size
}

func (td *tilesDrawer) Draw(t float64) {

	const maxT = 1.0 // 0.8

	for _, tile := range td.tiles {
		if t < maxT {
			if tile.PreviousPosition != nil {
				td.drawMoveTile(tile.Value, *(tile.PreviousPosition), tile.Position, t/maxT)
			} else if tile.MergedFrom != nil {
				mergedTiles := tile.MergedFrom
				for _, merged := range mergedTiles {
					prevPos := merged.Position
					if merged.PreviousPosition != nil {
						prevPos = *(merged.PreviousPosition)
					}
					td.drawMoveTile(merged.Value, prevPos, tile.Position, t/maxT)
				}
			}
		} else {
			var (
				dx = (td.cellSize.X + 1)
				dy = (td.cellSize.Y + 1)
			)

			pos := image.Point{
				X: tile.Position.X * dx,
				Y: tile.Position.Y * dy,
			}

			td.drawTile(tile.Value, pos)
		}
	}

	/*
		{
			if tile.PreviousPosition != nil {
				if t < maxT {
					drawMoveTile(vd, tile.Value, *(tile.PreviousPosition), tile.Position, t/maxT)
				} else {
					drawTile(vd, tile.Value, tile.Position)
				}
			} else if tile.MergedFrom != nil {
				if t < maxT {
					mergedTiles := tile.MergedFrom
					for _, merged := range mergedTiles {
						prevPos := merged.Position
						if merged.PreviousPosition != nil {
							prevPos = *(merged.PreviousPosition)
						}
						drawMoveTile(vd, merged.Value, prevPos, tile.Position, t/maxT)
					}
				} else {
					drawTile(vd, tile.Value, tile.Position)
				}
			} else {
				if t > maxT {
					if (int(math.Floor(t*20)) % 2) == 0 {
						drawTile(vd, tile.Value, tile.Position)
					}
				}
			}
		}
	*/
}

func (td *tilesDrawer) drawMoveTile(value int, begin, end image.Point, t float64) {
	var (
		dx = (td.cellSize.X + 1)
		dy = (td.cellSize.Y + 1)
	)

	position := image.Point{
		X: lerpInt(begin.X*dx, end.X*dx, t),
		Y: lerpInt(begin.Y*dy, end.Y*dy, t),
	}
	td.drawTile(value, position)
}

func (td *tilesDrawer) drawTile(value int, position image.Point) {

	rs, ok := td.valueRunes[value]
	if !ok {
		rs = valueToRunes(value, td.cellSize.X)
		td.valueRunes[value] = rs
	}

	gridPos := image.Point{
		X: (td.screenSize.X - ((td.cellSize.X+1)*td.gridSize.X + 1)) / 2,
		Y: (td.screenSize.Y - ((td.cellSize.Y+1)*td.gridSize.Y + 1)) / 2,
	}

	var (
		tilePosX = gridPos.X + 1 + position.X // *(td.cellSize.X+1)
		tilePosY = gridPos.Y + 1 + position.Y // *(td.cellSize.Y+1)
	)

	y := tilePosY
	for i := 0; i < td.cellSize.Y; i++ {
		x := tilePosX
		if i == td.cellSize.Y/2 {
			for _, r := range rs {
				termbox.SetCell(x, y, r, td.fg, td.bg)
				x++
			}
		} else {
			for range rs {
				termbox.SetCell(x, y, ' ', td.fg, td.bg)
				x++
			}
		}
		y++
	}
}

func valueToRunes(v int, size int) []rune {

	rs := make([]rune, size)

	var negative bool
	if v < 0 {
		v = -v
		negative = true
	}

	pos := len(rs) - 1

	var digit int
	for i := len(rs); i > 0; i-- {
		v, digit = quoRem(v, 10)
		if digit != 0 {
			pos = i - 1
		}
		rs[i-1] = rune('0' + digit)
	}

	if negative && (pos > 0) {
		pos--
		rs[pos] = '-'
	}

	newPos, rem := quoRem(pos, 2)
	if rem > 0 {
		newPos++
	}

	for i := 0; i < newPos; i++ {
		rs[i] = ' '
	}

	copy(rs[newPos:], rs[pos:])

	for i := newPos + len(rs) - pos; i < len(rs); i++ {
		rs[i] = ' '
	}

	return rs
}
