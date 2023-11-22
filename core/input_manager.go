package core

type InputManager interface {
	Restart()
	KeepPlaying()
	Move(d Direction)
	UndoMove()
	Draw()
}
