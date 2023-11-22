package main

import (
	"image"

	"github.com/nsf/termbox-go"

	"github.com/gitchander/go2048/core"
)

type imageDrawer struct {
	Fg termbox.Attribute // Foreground
	Bg termbox.Attribute // Background
}

func (d *imageDrawer) Box(bounds image.Rectangle) {
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			termbox.SetCell(x, y, ' ', d.Fg, d.Bg)
		}
	}
}

func (d *imageDrawer) Border(bounds image.Rectangle) {

	drawLine := func(y int, rs []rune) {
		x := bounds.Min.X
		if x < bounds.Max.X {
			termbox.SetCell(x, y, rs[0], d.Fg, d.Bg)
			x++
		}
		for x+1 < bounds.Max.X {
			termbox.SetCell(x, y, rs[1], d.Fg, d.Bg)
			x++
		}
		if x < bounds.Max.X {
			termbox.SetCell(x, y, rs[3], d.Fg, d.Bg)
			x++
		}
	}

	table := core.BorderTable(1)

	y := bounds.Min.Y
	if y < bounds.Max.Y {
		drawLine(y, table[0])
		y++
	}
	for y+1 < bounds.Max.Y {
		drawLine(y, table[1])
		y++
	}
	if y < bounds.Max.Y {
		drawLine(y, table[3])
		y++
	}
}

func (d *imageDrawer) Text(pos image.Point, text string) {
	for _, r := range text {
		termbox.SetCell(pos.X, pos.Y, r, d.Fg, d.Bg)
		pos.X++
	}
}

func (d *imageDrawer) Lines(pos image.Point, lines []string) {
	for _, s := range lines {
		d.Text(pos, s)
		pos.Y++
	}
}
