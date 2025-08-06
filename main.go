package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// SimulateTetris simulates a game of Tetris with the given block placements
func SimulateTetris(placements []string, width, height int) int {
	game := NewGame(width, height)

	for _, placement := range placements {
		blockType := string(placement[0])
		x := int(placement[1] - '0')

		block, exists := blocks[blockType]
		if !exists {
			return -1 // Invalid block type
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

	return game.getHighestElement()
}

// PlayTetris starts an interactive Tetris game
func PlayTetris(width, height int) {
	game := NewGame(width, height)
	scanner := bufio.NewScanner(os.Stdin)

	// Get available block types
	var blockTypes []string
	for blockType := range blocks {
		blockTypes = append(blockTypes, blockType)
	}

	for {
		// Pick random block type
		blockType := blockTypes[rand.Intn(len(blockTypes))]
		block := blocks[blockType]

		// Print current board
		fmt.Printf("\nCurrent board:\n")
		game.printGrid()

		// Ask player for x position
		fmt.Printf("Block type: %s\n", blockType)
		game.printBlock(block.shape)
		fmt.Printf("Enter x position (0-%d): ", width-1)

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "quit" || input == "q" {
			break
		}

		x := int(input[0] - '0')
		if x < 0 || x >= width {
			fmt.Println("Invalid x position!")
			continue
		}

		// Find placement and place block
		y := game.findPlacement(block, x)
		if y == -1 {
			fmt.Println("Game Over! Cannot place block.")
			break
		}

		game.placeBlock(block, x, y)
		game.clearLines()

		// Check if game is over (block placed at top)
		if y <= 0 {
			fmt.Println("Game Over!")
			game.printGrid()
			break
		}
	}

	fmt.Printf("Final score: %d\n", game.getHighestElement())
}

func main() {
	width, height := 10, 15

	PlayTetris(width, height)

}
