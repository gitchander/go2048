package core

import (
	"math/rand"

	"github.com/gitchander/go2048/utils"
)

type gameState struct {
	Score       int
	Over        bool
	Won         bool
	KeepPlaying bool
	Size        Point
	Tiles       []*Tile
}

type GameManager struct {
	score          int
	over           bool
	won            bool
	keepPlaying    bool
	grid           *grid
	mapTraversals  map[Direction]*traversals
	storageManager *StorageManager
	handler        Handler
	startTiles     int
	randSource     *rand.Rand
}

func NewGameManager(storage Storage, handler Handler) *GameManager {

	gm := &GameManager{
		//grid:        newGrid(defaultSize),
		//mapTraversals:  buildMapTraversals(defaultSize),
		storageManager: NewStorageManager(storage),
		handler:        handler,
		startTiles:     2,
		randSource:     utils.NewRandNow(),
	}

	gm.setup()

	return gm
}

func (gm *GameManager) setup() {

	size := DefaultSize()

	var previousState, _ = gm.storageManager.getGameState()

	// Reload the game from a previous game if present
	if previousState != nil {

		size = previousState.Size
		grid := newGrid(size)
		for _, t := range previousState.Tiles {
			grid.insertTile(newTile(t.Position, t.Value))
		}

		gm.grid = grid // Reload grid
		gm.mapTraversals = makeMapTraversals(size)

		gm.score = previousState.Score
		gm.over = previousState.Over
		gm.won = previousState.Won
		gm.keepPlaying = previousState.KeepPlaying

	} else {
		gm.grid = newGrid(size)
		gm.mapTraversals = makeMapTraversals(size)

		gm.score = 0
		gm.over = false
		gm.won = false
		gm.keepPlaying = false

		// Add the initial tiles
		gm.addStartTiles()
	}

	gm.handler.Init(size)

	// Update the actuator
	gm.actuate()
}

// Set up the initial tiles to start the game with
func (gm *GameManager) addStartTiles() {
	for i := 0; i < gm.startTiles; i++ {
		gm.grid.addRandomTile(gm.randSource)
	}
}

// Restart the game
func (gm *GameManager) Restart() {
	gm.storageManager.clearGameState()
	gm.handler.Message(MK_CLEAR)
	gm.setup()
}

// Keep playing after winning (allows going over 2048)
func (gm *GameManager) KeepPlaying() {
	gm.keepPlaying = true
	gm.handler.Message(MK_CLEAR)
}

// Sends the updated grid to the actuator
func (gm *GameManager) actuate() {

	if gm.storageManager.getBestScore() < gm.score {
		gm.storageManager.setBestScore(gm.score)
	}

	var tiles []*Tile
	gm.grid.rangeTiles(
		func(t *Tile) bool {
			tiles = append(tiles, t)
			return true
		},
	)

	// Clear the state when the game is over (game over only, not win)
	if gm.over {
		gm.storageManager.clearGameState()
	} else {

		state := gameState{
			Score:       gm.score,
			Over:        gm.over,
			Won:         gm.won,
			KeepPlaying: gm.keepPlaying,
			Size:        gm.grid.Size(),
			Tiles:       tiles,
		}

		gm.storageManager.setGameState(&state)
	}

	terminated := gm.isGameTerminated()

	//-----------------------------------
	//	gm.actuator.actuate(
	//		tiles,
	//		gm.score,
	//		gm.storageManager.getBestScore(),
	//		gm.over,
	//		gm.won,
	//		terminated,
	//	)
	//-----------------------------------
	// gm.handler

	bestScore := gm.storageManager.getBestScore()

	gm.handler.AnimationRequest(tiles)

	gm.handler.UpdateScore(gm.score)
	gm.handler.UpdateBestScore(bestScore)

	if terminated {
		if gm.over {
			gm.handler.Message(MK_YOU_LOSE) // You lose
		} else if gm.won {
			gm.handler.Message(MK_YOU_WIN) // You win!
		}
	}
}

// Return true if the game is lost, or has won and the user hasn't kept playing
func (gm *GameManager) isGameTerminated() bool {
	return gm.over || (gm.won && not(gm.keepPlaying))
}

// Move tiles on the grid in the specified direction
func (gm *GameManager) Move(d Direction) {

	vector := d.getVector()

	if gm.isGameTerminated() {
		// Don't do anything if the game's over
		return
	}

	var moved bool

	// Save the current tile positions and remove merger information
	gm.grid.resetTiles()

	rf := func(cell Point) bool {
		current := gm.grid.cellContent(cell)
		if current != nil {
			var positions = gm.grid.findFarthestPosition(cell, vector)
			var next = gm.grid.cellContent(positions.next)

			// Only one merger per row traversal?
			if (next != nil) && (next.Value == current.Value) && (next.MergedFrom == nil) {

				var merged = mergeTiles(positions.next, current, next)

				gm.grid.insertTile(merged)
				gm.grid.removeTile(current)

				// Converge the two tiles' positions
				current.updatePosition(positions.next)

				// Update the score
				gm.score += merged.Value

				// The mighty 2048 tile
				if merged.Value == 2048 {
					gm.won = true
				}

			} else {
				gm.grid.moveTile(current, positions.farthest)
			}

			if not(cell.Eq(current.Position)) {
				moved = true // The tile moved from its original cell!
			}
		}
		return true
	}

	// Traverse the grid in the right direction and move tiles
	t := gm.mapTraversals[d]
	t.Range(rf)

	if moved {
		gm.grid.addRandomTile(gm.randSource)
		if not(gm.grid.movesAvailable()) {
			gm.over = true // Game over!
		}
		gm.actuate()
	}
}

func (gm *GameManager) UndoMove() {

	var undoOK bool

	newGrid := newGrid(gm.grid.Size())
	gm.grid.rangeTiles(
		func(t *Tile) bool {
			if t.PreviousPosition != nil {
				undoOK = true
				newGrid.insertTile(newTile(*t.PreviousPosition, t.Value))
			} else if t.MergedFrom != nil {
				undoOK = true
				for _, merged := range t.MergedFrom {
					prevPos := merged.Position
					if merged.PreviousPosition != nil {
						prevPos = *(merged.PreviousPosition)
					}
					newGrid.insertTile(newTile(prevPos, merged.Value))
				}
			}
			return true
		},
	)

	if undoOK {
		gm.grid = newGrid
		gm.actuate()
	}
}

func (gm *GameManager) Draw() {
	gm.actuate()
}

func (gm *GameManager) PrintableGrid() []byte {
	return encodePrintableTest(gm.grid)
}
