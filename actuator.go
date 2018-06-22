package go2048

import (
	"image"
	"strconv"
)

type MessageKind int

const (
	MK_CLEAR    MessageKind = iota // Clear the game won/lost message
	MK_YOU_LOSE                    // You lose
	MK_YOU_WIN                     // You win!
)

var name_MessageKind = map[MessageKind]string{
	MK_CLEAR:    "Clear",
	MK_YOU_LOSE: "You lose",
	MK_YOU_WIN:  "You win!",
}

func (mk MessageKind) String() string {
	if name, ok := name_MessageKind[mk]; ok {
		return name
	}
	return strconv.Itoa(int(mk))
}

type Handler interface {
	Init(size image.Point)
	AnimationRequest(tiles []*Tile)
	UpdateScore(score int)
	UpdateBestScore(bestScore int)
	Message(MessageKind)
}

type DummyHandler struct{}

func (DummyHandler) Init(size image.Point)          {}
func (DummyHandler) AnimationRequest(tiles []*Tile) {}
func (DummyHandler) UpdateScore(score int)          {}
func (DummyHandler) UpdateBestScore(bestScore int)  {}
func (DummyHandler) Message(MessageKind)            {}
