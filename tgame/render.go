package main

import (
	"fmt"
	"image"

	"github.com/nsf/termbox-go"

	game "github.com/gitchander/go2048"
)

type animationInfo struct {
	inProgress bool
	start      float64
	duration   float64 // seconds
	current    float64
}

func (ai *animationInfo) done() bool {
	return !(ai.inProgress)
}

type Drawer struct {
	screenSize image.Point
	gridSize   image.Point
	cellSize   image.Point

	animation animationInfo

	score, bestScore int
	mk               game.MessageKind
	gd               *gridDrawer
	td               *tilesDrawer
}

//func NewDrawer(gridSize, cellSize image.Point) *Drawer {
func NewDrawer(cellSize image.Point) *Drawer {

	return &Drawer{
		cellSize: cellSize,
		animation: animationInfo{
			duration: 0.1,
		},
	}

	//	return &Drawer{
	//		gridSize: gridSize,
	//		cellSize: cellSize,
	//		animation: animationInfo{
	//			duration: 0.1,
	//		},
	//		gd: &gridDrawer{
	//			gridSize: gridSize,
	//			cellSize: cellSize,
	//			fg:       termbox.ColorCyan | termbox.AttrBold,
	//			bg:       termbox.ColorCyan,
	//			table:    tableRune1,
	//		},
	//		td: &tilesDrawer{
	//			gridSize:   gridSize,
	//			cellSize:   cellSize,
	//			fg:         termbox.ColorYellow | termbox.AttrBold,
	//			bg:         termbox.ColorBlue,
	//			valueRunes: make(map[int][]rune),
	//		},
	//	}
}

func (d *Drawer) Init(gridSize image.Point) {

	d.gridSize = gridSize
	//cellSize: cellSize,
	//		animation: animationInfo{
	//			duration: 0.1,
	//		},

	d.gd = &gridDrawer{
		gridSize: gridSize,
		cellSize: d.cellSize,
		fg:       termbox.ColorCyan | termbox.AttrBold,
		bg:       termbox.ColorCyan,
		table:    game.BorderTable(5),
	}
	d.td = &tilesDrawer{
		gridSize:   gridSize,
		cellSize:   d.cellSize,
		fg:         termbox.ColorYellow | termbox.AttrBold,
		bg:         termbox.ColorBlue,
		valueRunes: make(map[int][]rune),
	}

	d.td.SetScreenSize(d.screenSize)
	d.gd.SetScreenSize(d.screenSize)
}

func (a *Drawer) AnimationRequest(tiles []*game.Tile) {

	a.animation.start = getTime()
	a.animation.inProgress = true

	a.td.tiles = tiles
}

func (a *Drawer) UpdateTime(t float64) {
	a.animation.current = t
}

func (a *Drawer) UpdateScore(score int) {
	a.score = score
}

func (a *Drawer) UpdateBestScore(bestScore int) {
	a.bestScore = bestScore
}

func (a *Drawer) Message(mk game.MessageKind) {
	a.mk = mk
}

func (a *Drawer) Draw() {
	if a.animation.inProgress {
		if dcurr := (a.animation.current - a.animation.start); dcurr < a.animation.duration {
			t := dcurr / a.animation.duration
			a.draw(a.score, a.bestScore, t)
			return
		}
		a.draw(a.score, a.bestScore, 1)
		a.animation.inProgress = false
	} else {
		a.draw(a.score, a.bestScore, 1)
	}
}

func (a *Drawer) SetScreenSize(size image.Point) {
	a.screenSize = size
	a.td.SetScreenSize(size)
	a.gd.SetScreenSize(size)
}

func (a *Drawer) draw(score, bestScore int, t float64) {

	renderScore(a.screenSize, a.gridSize, a.cellSize, score, bestScore, a.mk)

	a.gd.Draw()
	a.td.Draw(t)
}

var controls = []string{
	"Controls:",
	"↑ or w - move up",
	"← or a - move left",
	"↓ or s - move down",
	"→ or d - move right",
	"     u - undo move",
	"     r - restart",
	"     k - keep playing",
	"   esc - quit",
}

func renderScore(tbSize, gridSize, cellSize image.Point, score, bestScore int, mk game.MessageKind) {

	var dx, dy = gridSize.X, gridSize.Y

	gDx := (cellSize.X*dx + dx + 1)
	gDy := (cellSize.Y*dy + dy + 1)

	x0 := (tbSize.X - gDx) / 2
	y0 := (tbSize.Y - gDy) / 2

	var d imageDrawer

	d.Fg = termbox.ColorYellow | termbox.AttrBold
	d.Bg = termbox.ColorBlue

	d.Box(image.Rect(x0, y0-6, x0+gDx, y0-1))

	textBestScore := fmt.Sprintf("Best  : %5d", bestScore)
	textScore := fmt.Sprintf("Score : %5d", score)

	d.Text(image.Point{x0 + 3, y0 - 4}, "Go 2048")
	d.Text(image.Point{x0 + gDx - len(textBestScore) - 2, y0 - 5}, textBestScore)
	d.Text(image.Point{x0 + gDx - len(textScore) - 2, y0 - 3}, textScore)

	// render Footer
	d.Fg = termbox.ColorCyan // | termbox.AttrBold
	d.Border(image.Rect(x0, y0+gDy+1, x0+gDx, y0+gDy+5+len(controls)))

	switch mk {
	case game.MK_YOU_LOSE:
		d.Fg = termbox.ColorRed | termbox.AttrBold
		d.Text(image.Point{x0 + 2, y0 + gDy + 2}, mk.String())
	case game.MK_YOU_WIN:
		d.Fg = termbox.ColorGreen | termbox.AttrBold
		d.Text(image.Point{x0 + 2, y0 + gDy + 2}, mk.String())
	}

	d.Fg = termbox.ColorWhite // | termbox.AttrBold
	d.Lines(image.Point{x0 + 2, y0 + gDy + 4}, controls)
}
