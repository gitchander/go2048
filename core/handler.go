package core

type Handler interface {
	Init(size Point)
	AnimationRequest(tiles []*Tile)
	UpdateScore(score int)
	UpdateBestScore(bestScore int)
	Message(MessageKind)
}

type DummyHandler struct{}

func (DummyHandler) Init(size Point)                {}
func (DummyHandler) AnimationRequest(tiles []*Tile) {}
func (DummyHandler) UpdateScore(score int)          {}
func (DummyHandler) UpdateBestScore(bestScore int)  {}
func (DummyHandler) Message(MessageKind)            {}
