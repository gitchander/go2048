package main

import (
	"image"

	"github.com/nsf/termbox-go"

	"github.com/gitchander/go2048/core"
)

type PlayState struct {
	sm           *StateManager
	drawer       *Drawer
	inputManager core.InputManager
}

var _ State = &PlayState{}

func NewPlayState(sm *StateManager) *PlayState {

	//size := image.Point{4, 4}
	//size := image.Point{8, 8}

	cellSize := image.Point{7, 3}
	//cellSize := image.Point{4, 1}
	drawer := NewDrawer(cellSize)

	gm := core.NewGameManager(sm.storage, drawer)

	var inputManager core.InputManager = gm

	return &PlayState{
		sm:           sm,
		drawer:       drawer,
		inputManager: inputManager,
	}
}

func (ps *PlayState) MoveLeft() {
	if ps.drawer.animation.done() {
		ps.inputManager.Move(core.Left)
	}
}

func (ps *PlayState) MoveRight() {
	if ps.drawer.animation.done() {
		ps.inputManager.Move(core.Right)
	}
}

func (ps *PlayState) MoveDown() {
	if ps.drawer.animation.done() {
		ps.inputManager.Move(core.Down)
	}
}

func (ps *PlayState) Restart() {
	if ps.drawer.animation.done() {
		ps.inputManager.Restart()
	}
}

func (ps *PlayState) MoveUp() {
	if ps.drawer.animation.done() {
		ps.inputManager.Move(core.Up)
	}
}

func (ps *PlayState) UndoMove() {
	if ps.drawer.animation.done() {
		ps.inputManager.UndoMove()
	}
}

func (ps *PlayState) KeepPlaying() {
	if ps.drawer.animation.done() {
		ps.inputManager.KeepPlaying()
	}
}

func (ps *PlayState) EventKey(Key termbox.Key) {
	switch Key {
	case termbox.KeyEsc:
		ps.sm.breakContinue()
		return
	case termbox.KeyArrowLeft:
		ps.MoveLeft()
	case termbox.KeyArrowRight:
		ps.MoveRight()
	case termbox.KeyArrowDown:
		ps.MoveDown()
	case termbox.KeyArrowUp:
		ps.MoveUp()
	}
}

func (ps *PlayState) EventCharacter(Ch rune) {
	switch Ch {
	case 'w':
		ps.MoveUp()
	case 'a':
		ps.MoveLeft()
	case 's':
		ps.MoveDown()
	case 'd':
		ps.MoveRight()
	case 'r':
		ps.Restart()
	case 'u':
		ps.UndoMove()
	case 'k':
		ps.KeepPlaying()
	}
}

func (ps *PlayState) EventResize(size image.Point) {
	ps.drawer.SetScreenSize(size)
}

func (ps *PlayState) EventMouse(pos image.Point) {

}

func (ps *PlayState) UpdateTime(t float64) {
	ps.drawer.UpdateTime(t)
}

func (ps *PlayState) Render() {
	ps.drawer.Draw()
}
