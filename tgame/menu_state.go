package main

import (
	"image"

	"github.com/nsf/termbox-go"
)

type MenuItem struct {
	Title  string
	Action func(*StateManager)
}

type MenuState struct {
	sm          *StateManager
	items       []MenuItem
	activeIndex int
}

func NewMenuState(sm *StateManager, items []MenuItem) *MenuState {
	activeIndex := -1
	if len(items) > 0 {
		activeIndex = 0
	}
	return &MenuState{
		sm:          sm,
		items:       items,
		activeIndex: activeIndex,
	}
}

var _ State = &MenuState{}

func (ms *MenuState) UpdateTime(t float64) {

}

func (ms *MenuState) Render() {

}

func (ms *MenuState) EventResize(size image.Point) {

}

func (ms *MenuState) EventKey(Key termbox.Key) {
	switch Key {
	case termbox.KeyEsc:
	case termbox.KeyArrowLeft:
	case termbox.KeyArrowRight:
	case termbox.KeyArrowUp:
	case termbox.KeyArrowDown:
	}
}

func (ms *MenuState) EventCharacter(Ch rune) {

}

func (ms *MenuState) EventMouse(pos image.Point) {

}
