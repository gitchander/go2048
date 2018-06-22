package main

import (
	"fmt"
	"log"

	pot "github.com/gitchander/go2048"
)

func main() {
	Main()
}

func Main() {

	var fs = pot.NewFakeStorage()

	gm := pot.NewGameManager(fs, pot.DummyAnimationRequester{})

	printTable(gm)

	var line string

start:
	for {
		_, err := fmt.Scanf("%s", &line)
		if err != nil {
			log.Fatal(err)
		}

		switch line {
		case "a", "A":
			gm.Move(pot.Left)
		case "d", "D":
			gm.Move(pot.Right)
		case "w", "W":
			gm.Move(pot.Up)
		case "s", "S":
			gm.Move(pot.Down)
		case "q", "Q":
			fallthrough
		case "e", "E":
			break start
		}

		printTable(gm)
	}
}

func printTable(gm *pot.GameManager) {
	fmt.Println(string(gm.DataTable()))
}
