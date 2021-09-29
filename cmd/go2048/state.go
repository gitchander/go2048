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

type dummyState struct{}

func (dummyState) EventResize(size image.Point) {}
func (dummyState) EventKey(Key termbox.Key)     {}
func (dummyState) EventCharacter(Ch rune)       {}
func (dummyState) EventMouse(pos image.Point)   {}
func (dummyState) UpdateTime(t float64)         {}
func (dummyState) Render()                      {}

var _ State = dummyState{}
