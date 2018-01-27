package go2048

type InputManager interface {
	Restart()
	KeepPlaying()
	Move(d Direction)
	UndoMove()
	Draw()
}
