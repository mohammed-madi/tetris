package main

// Block represents a Tetris block with its shape and position
type Block struct {
	shape  [][]bool
	width  int
	height int
}

// Game represents the Tetris game state
type Game struct {
	grid   [][]bool
	width  int
	height int
}

// Block definitions
var blocks = map[string]Block{
	"Q": { // Cube
		shape: [][]bool{
			{true, true},
			{true, true},
		},
		width:  2,
		height: 2,
	},
	"Z": { // Left facing Z
		shape: [][]bool{
			{true, true, false},
			{false, true, true},
		},
		width:  3,
		height: 2,
	},
	"S": { // Right facing S
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

// canPlace checks if a block can be placed at the given position
func (g *Game) canPlace(block Block, x, y int) bool {
	if x < 0 || y < 0 || x+block.width > g.width || y+block.height > g.height {
		return false
	}

	for dy := 0; dy < block.height; dy++ {
		for dx := 0; dx < block.width; dx++ {
			if block.shape[dy][dx] && g.grid[y+dy][x+dx] {
				return false
			}
		}
	}
	return true
}

// placeBlock places a block at the given position
func (g *Game) placeBlock(block Block, x, y int) bool {
	if !g.canPlace(block, x, y) {
		return false
	}

	for dy := 0; dy < block.height; dy++ {
		for dx := 0; dx < block.width; dx++ {
			if block.shape[dy][dx] {
				g.grid[y+dy][x+dx] = true
			}
		}
	}
	return true
}

// findPlacement finds the lowest valid position for a block
func (g *Game) findPlacement(block Block, x int) int {
	// Start from the top and move down
	for y := 0; y <= g.height-block.height; y++ {
		if g.canPlace(block, x, y) {
			// Check if this is the lowest position (try one more down)
			if y+block.height >= g.height || !g.canPlace(block, x, y+1) {
				return y
			}
		}
	}
	return -1 // Cannot place
}

// clearLines removes full lines and returns the number of lines cleared
func (g *Game) clearLines() int {
	linesCleared := 0

	for y := g.height - 1; y >= 0; y-- {
		full := true
		for x := 0; x < g.width; x++ {
			if !g.grid[y][x] {
				full = false
				break
			}
		}

		if full {
			// Remove the line by shifting everything down
			for yy := y; yy > 0; yy-- {
				copy(g.grid[yy], g.grid[yy-1])
			}
			// Clear the top line
			for x := 0; x < g.width; x++ {
				g.grid[0][x] = false
			}
			linesCleared++
			y++ // Check the same line again since everything shifted down
		}
	}

	return linesCleared
}

// getHighestY returns the Y position of the highest block (bottom = 0)
func (g *Game) getHighestY() int {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			if g.grid[y][x] {
				return g.height - y // Convert to bottom=0 coordinate system
			}
		}
	}
	return 0 // No blocks placed, so height from bottom is 0
}
