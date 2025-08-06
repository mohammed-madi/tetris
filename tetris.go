package main

import "fmt"

// Block represents a Tetris block with its shape and position
type Block struct {
	shape  [][]bool
	width  int
	height int
}

// Block definitions
var blocks = map[string]Block{
	"Q": { // Square
		shape: [][]bool{
			{true, true},
			{true, true},
		},
		width:  2,
		height: 2,
	},
	"Z": { // Z shaped block
		shape: [][]bool{
			{true, true, false},
			{false, true, true},
		},
		width:  3,
		height: 2,
	},
	"S": { // S shaped block
		shape: [][]bool{
			{false, true, true},
			{true, true, false},
		},
		width:  3,
		height: 2,
	},
	"T": { // Downward facing T
		shape: [][]bool{
			{true, true, true},
			{false, true, false},
		},
		width:  3,
		height: 2,
	},
	"J": { // Left facing L shape
		shape: [][]bool{
			{false, true},
			{false, true},
			{true, true},
		},
		width:  2,
		height: 3,
	},
	"L": { // Right facing L shape
		shape: [][]bool{
			{true, false},
			{true, false},
			{true, true},
		},
		width:  2,
		height: 3,
	},
	"I": { // Horizontal 4-long line
		shape: [][]bool{
			{true, true, true, true},
		},
		width:  4,
		height: 1,
	},
}

// Game represents the Tetris game state
type Game struct {
	grid   [][]bool // grid is inverted: grid[0] is the top row
	width  int
	height int
}

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

// canPlace checks if a block can be placed at the given position
func (g *Game) canPlace(block Block, x, y int) bool {
	if x < 0 || y < 0 || x+block.width > g.width || y+block.height > g.height {
		return false
	}

	for dy := 0; dy < block.height; dy++ {
		for dx := 0; dx < block.width; dx++ {
			// check all block segments are not colliding with existing blocks
			if block.shape[dy][dx] && g.grid[y+dy][x+dx] {
				return false
			}
		}
	}
	return true
}

// placeBlock places a block at the given position
func (g *Game) placeBlock(block Block, x, y int) {
	for dy := 0; dy < block.height; dy++ {
		for dx := 0; dx < block.width; dx++ {
			if block.shape[dy][dx] {
				g.grid[y+dy][x+dx] = true
			}
		}
	}
}

// findPlacement finds the lowest valid position for a block
func (g *Game) findPlacement(block Block, x int) int {
	// Start from the top and move down
	for y := 0; y <= g.height-block.height; y++ {
		if g.canPlace(block, x, y) && !g.canPlace(block, x, y+1) {
			return y
		}
	}
	return -1 // Cannot place
}

// clearLines removes full lines and shifts down all lines above
func (g *Game) clearLines() int {
	lines := 0
	for y := g.height - 1; y >= 0; y-- {
		full := true
		for x := 0; x < g.width; x++ {
			if !g.grid[y][x] {
				full = false
				break
			}
		}

		if full {
			lines++
			// Remove the line by shifting everything down
			for yy := y; yy > 0; yy-- {
				copy(g.grid[yy], g.grid[yy-1])
			}
			// Clear the top line
			clear(g.grid[0])
			// Check the same line again since everything shifted down
			y++
		}
	}
	return lines
}

// getHighestElement returns the Y position (or height) of the highest block element
func (g *Game) getHighestElement() int {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			if g.grid[y][x] {
				return g.height - y // Convert to bottom=0 coordinate system
			}
		}
	}
	return 0 // No blocks placed, so height from bottom is 0
}

// printGrid prints the current game grid for debugging purposes
func (g *Game) printGrid() {
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
		fmt.Printf("%2d: %s\n", g.height-y-1, line)
	}

	// Print x coordinates
	fmt.Print("    ")
	for x := 0; x < g.width; x++ {
		fmt.Printf("%d", x%10)
	}
	fmt.Println()
	fmt.Println()
}

// printBlock prints the shape of a block to show the user its structure
func (g *Game) printBlock(block [][]bool) {
	fmt.Println("Block shape:")
	for _, row := range block {
		for _, cell := range row {
			if cell {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
