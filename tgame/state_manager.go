package main

import (
	"errors"
	"image"

	"github.com/nsf/termbox-go"

	game "github.com/gitchander/go2048"
)

type StateManager struct {
	storage    game.Storage
	screenSize image.Point
	f_quit     bool
	states     map[string]State
	active     State
}

func NewStateManager(storage game.Storage) *StateManager {
	return &StateManager{
		storage: storage,
		states:  make(map[string]State),
		active:  fakeState{},
	}
}

func (sm *StateManager) AppendState(id string, state State) error {
	if _, ok := sm.states[id]; ok {
		return errors.New("state id has already exist")
	}
	sm.states[id] = state
	return nil
}

func (sm *StateManager) UpdateTime(t float64) {
	sm.active.UpdateTime(t)
}

func (sm *StateManager) SetState(id string) error {
	state, ok := sm.states[id]
	if !ok {
		return errors.New("state id is missing")
	}
	state.EventResize(sm.screenSize)
	sm.active = state
	return nil
}

func (sm *StateManager) Render() {

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	//termbox.Clear(termbox.ColorDefault, termbox.ColorWhite)

	sm.active.Render()

	termbox.Flush()
}

func (sm *StateManager) EventKeyChar(Key termbox.Key, Ch rune) {
	if Ch == 0 {
		sm.active.EventKey(Key)
	} else {
		sm.active.EventCharacter(Ch)
	}
}

func (sm *StateManager) EventResize(width, height int) {
	size := image.Point{X: width, Y: height}
	sm.screenSize = size
	sm.active.EventResize(size)
}

func (sm *StateManager) EventMouse(posX, posY int) {
	sm.active.EventMouse(image.Point{X: posX, Y: posY})
}

func (sm *StateManager) Continue() bool {
	return !(sm.f_quit)
}

func (sm *StateManager) breakContinue() {
	sm.f_quit = true
}
