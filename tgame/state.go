package main

import (
	"image"

	"github.com/nsf/termbox-go"
)

type State interface {
	EventResize(size image.Point)
	EventKey(Key termbox.Key)
	EventCharacter(Ch rune)
	EventMouse(pos image.Point)
	UpdateTime(t float64)
	Render()
}

type fakeState struct{}

func (fakeState) EventResize(size image.Point) {}
func (fakeState) EventKey(Key termbox.Key)     {}
func (fakeState) EventCharacter(Ch rune)       {}
func (fakeState) EventMouse(pos image.Point)   {}
func (fakeState) UpdateTime(t float64)         {}
func (fakeState) Render()                      {}

var _ State = fakeState{}
