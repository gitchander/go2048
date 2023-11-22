package core

import (
	"fmt"
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
	return fmt.Sprintf("%s(%d)", "MessageKind", int(mk))
}
