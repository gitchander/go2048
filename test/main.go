package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	pot "github.com/gitchander/go2048"
)

func main() {
	var fs = pot.NewFakeStorage()

	gm := pot.NewGameManager(fs, pot.DummyHandler{})
	printGrid(gm)

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		line, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		line = strings.TrimSuffix(line, "\n")

		switch line {
		case "":
			// do nothing
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
			os.Exit(0)
		default:
			fmt.Println("bad command:", line)
		}

		printGrid(gm)
	}
}

func printGrid(gm *pot.GameManager) {
	fmt.Print(string(gm.PrintableGrid()))
}
