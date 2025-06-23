package main

import (
	"bufio"
	"fmt"
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

	return game.getHighestY()
}

func main() {
	width, height := 10, 100
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		placements := strings.Split(line, ",")
		highestY := SimulateTetris(placements, width, height)
		fmt.Println(highestY)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
