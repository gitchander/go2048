package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gitchander/go2048/core"
)

func main() {
	ms := core.NewMapStorage()
	gm := core.NewGameManager(ms, core.DummyHandler{})
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
			gm.Move(core.Left)
		case "d", "D":
			gm.Move(core.Right)
		case "w", "W":
			gm.Move(core.Up)
		case "s", "S":
			gm.Move(core.Down)
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

func printGrid(gm *core.GameManager) {
	fmt.Print(string(gm.PrintableGrid()))
}
