package main

import (
	"fmt"
)

// NewGame creates a new Tetris game
func NewGame(width, height int) *Game {
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
	}
	return &Game{
		grid:   grid,
		width:  width,
		height: height,
	}
}

// SimulateTetris simulates a game of Tetris with the given block placements
func SimulateTetris(placements []string, width, height int) int {
	game := NewGame(width, height)

	for _, placement := range placements {

		blockType := string(placement[0])
		x := int(placement[1] - '0')

		block, exists := blocks[blockType]
		if !exists {
			continue
		}

		// Find the lowest valid position for this block
		y := game.findPlacement(block, x)
		if y == -1 {
			return -1 // Cannot place this block
		}

		// Place the block
		game.placeBlock(block, x, y)

		// Clear any full lines
		game.clearLines()
	}
	game.PrintGrid()
	return game.getHighestY()
}

// PrintGrid prints the current game grid for debugging
func (g *Game) PrintGrid() {
	fmt.Println("Grid:")
	for y := 0; y < g.height; y++ {
		line := ""
		for x := 0; x < g.width; x++ {
			if g.grid[y][x] {
				line += "█"
			} else {
				line += "."
			}
		}
		fmt.Printf("%2d: %s\n", y, line)
	}
	fmt.Println()
}

func main() {

	placements := []string{"I0", "I4", "I0", "I4", "I4", "I5"}
	width, height := 10, 20

	highestY := SimulateTetris(placements, width, height)
	fmt.Printf("Highest Y position: %d\n", highestY)

	// Demonstrate with a visual example
	fmt.Println("Example simulation:")
	game := NewGame(10, 10)

	examplePlacements := []string{"Q0", "S4", "Z1"}
	for i, placement := range examplePlacements {
		fmt.Printf("Step %d: Placing %s\n", i+1, placement)

		blockType := string(placement[0])
		x := int(placement[1] - '0')
		block := blocks[blockType]

		y := game.findPlacement(block, x)
		if y != -1 {
			game.placeBlock(block, x, y)
			game.clearLines()
		}

		game.PrintGrid()
	}

	fmt.Printf("Final highest Y position: %d\n", game.getHighestY())
}
