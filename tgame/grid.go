package main

import (
	"image"

	"github.com/nsf/termbox-go"
)

type gridDrawer struct {
	screenSize image.Point
	gridSize   image.Point
	cellSize   image.Point
	fg, bg     termbox.Attribute
	table      [][]rune
}

func (gd *gridDrawer) SetScreenSize(size image.Point) {
	gd.screenSize = size
}

func (gd *gridDrawer) Draw() {

	var dx, dy = gd.gridSize.X, gd.gridSize.Y

	x0 := (gd.screenSize.X - (gd.cellSize.X*dx + dx + 1)) / 2
	y0 := (gd.screenSize.Y - (gd.cellSize.Y*dy + dy + 1)) / 2

	y := y0

	gd.drawLine(x0, y, 0) // draw top border
	y++

	for iy := 0; iy < dy; iy++ {

		if iy > 0 {
			gd.drawLine(x0, y, 2) // draw middle border
			y++
		}

		for i := 0; i < gd.cellSize.Y; i++ {
			gd.drawLineSpaces(x0, y, 1, iy) // draw clear line
			y++
		}
	}

	gd.drawLine(x0, y, 3) // draw bottom border
	y++
}

func (gd *gridDrawer) drawLine(x, y int, ri int) {

	fg, bg := gd.fg, gd.bg

	dx := gd.gridSize.X

	rs := gd.table[ri]

	termbox.SetCell(x, y, rs[0], fg, bg)
	x++
	for ix := 0; ix < dx; ix++ {
		if ix > 0 {
			termbox.SetCell(x, y, rs[2], fg, bg)
			x++
		}
		for i := 0; i < gd.cellSize.X; i++ {
			termbox.SetCell(x, y, rs[1], fg, bg)
			x++
		}
	}
	termbox.SetCell(x, y, rs[3], fg, bg)
}

func (gd *gridDrawer) drawLineSpaces(x, y int, ri int, iy int) {

	fg, bg := gd.fg, gd.bg

	dx := gd.gridSize.X

	rs := gd.table[ri]

	termbox.SetCell(x, y, rs[0], fg, bg)
	x++
	for ix := 0; ix < dx; ix++ {
		if ix > 0 {
			termbox.SetCell(x, y, rs[2], fg, bg)
			x++
		}
		for i := 0; i < gd.cellSize.X; i++ {
			termbox.SetCell(x, y, ' ', fg, bg)
			x++
		}
	}
	termbox.SetCell(x, y, rs[3], fg, bg)
}

/*
func (ld *lineDrawer) drawLineValues(x, y int, ri int, iy int) {

	fg, bg := ld.fg, ld.bg

	dx, _ := ld.m.Size()

	rs := ld.table[ri]

	termbox.SetCell(x, y, rs[0], fg, bg)
	x++
	for ix := 0; ix < dx; ix++ {
		if ix > 0 {
			termbox.SetCell(x, y, rs[2], fg, bg)
			x++
		}
		if n, _ := ld.m.Get(ix, iy); n != nil {

			vrs, ok := ld.valueRunes[n.val]
			if !ok {
				vrs = valueToRunes(n.val, ld.cellSizeX)
				ld.valueRunes[n.val] = vrs
			}

			for _, r := range vrs {
				termbox.SetCell(x, y, r, n.fg, n.bg)
				x++
			}
		} else {
			for i := 0; i < ld.cellSizeX; i++ {
				termbox.SetCell(x, y, ' ', fg, bg)
				x++
			}
		}
	}
	termbox.SetCell(x, y, rs[3], fg, bg)
}
*/
