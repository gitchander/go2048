package main

import (
	"flag"
	"log"
	"time"

	"github.com/nsf/termbox-go"
)

func main() {

	var storageName string
	flag.StringVar(&storageName, "storage", "game.db", "database storage filename")
	flag.Parse()

	err := termbox.Init()
	checkError(err)
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	storage, err := NewBoltStorage(storageName)
	checkError(err)
	defer storage.Close()

	sm := NewStateManager(storage)

	width, height := termbox.Size()
	sm.EventResize(width, height)

	err = sm.AppendState("play", NewPlayState(sm))
	checkError(err)

	sm.AppendState("menu",
		NewMenuState(sm,
			[]MenuItem{
				MenuItem{
					Title: "Quit",
					Action: func(sm *StateManager) {
						sm.breakContinue()
					},
				},
				MenuItem{
					Title: "Play",
					Action: func(sm *StateManager) {
						sm.SetState("game")
					},
				},
			},
		),
	)

	sm.SetState("play")

	for sm.Continue() {
		select {
		case event := <-eventQueue:
			switch event.Type {
			case termbox.EventKey:
				sm.EventKeyChar(event.Key, event.Ch)
			case termbox.EventResize:
				sm.EventResize(event.Width, event.Height)
			case termbox.EventMouse:
				sm.EventMouse(event.MouseX, event.MouseY)
			}
		default:
		}
		sm.UpdateTime(getTime())
		sm.Render()
		time.Sleep(10 * time.Millisecond)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
