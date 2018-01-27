package go2048

import "strconv"

type MessageKind int

const (
	MK_CLEAR    MessageKind = iota
	MK_YOU_LOSE             // You lose
	MK_YOU_WIN              // You win!
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

type AnimationRequester interface {
	AnimationRequest(tiles []Tile)
	UpdateScore(score int)
	UpdateBestScore(bestScore int)
	Message(MessageKind)
}

type Actuator struct {
	ar AnimationRequester
}

// Clear the game won/lost message
func (a *Actuator) continueGame() {
	a.ar.Message(MK_CLEAR)
}

func (a *Actuator) actuate(g *grid, score int, bestScore int, over bool, won bool, terminated bool) {

	var tiles []Tile
	g.forEach(
		func(t *Tile) {
			tiles = append(tiles, *t)
		},
	)

	a.ar.AnimationRequest(tiles)

	a.ar.UpdateScore(score)
	a.ar.UpdateBestScore(bestScore)

	if terminated {
		if over {
			a.ar.Message(MK_YOU_LOSE) // You lose
		} else if won {
			a.ar.Message(MK_YOU_WIN) // You win!
		}
	}
}