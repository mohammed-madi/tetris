package main

import (
	"fmt"
	"testing"
)

func TestSimulateTetris(t *testing.T) {
	tests := []struct {
		name       string
		placements []string
		width      int
		height     int
		expected   int
	}{
		{
			name:       "Empty placements",
			placements: []string{},
			width:      10,
			height:     20,
			expected:   0, // No blocks placed, so height from bottom is 0
		},
		{
			name:       "Single cube",
			placements: []string{"Q0"},
			width:      10,
			height:     20,
			expected:   2, // Cube placed at bottom, so height from bottom is 2
		},
		{
			name:       "Stacked blocks",
			placements: []string{"Q0", "Q0"},
			width:      10,
			height:     20,
			expected:   4, // Two cubes stacked, height from bottom is 4
		},
		{
			name:       "Line clearing",
			placements: []string{"I0", "I4", "I0", "I4", "I4", "I5"},
			width:      10,
			height:     20,
			expected:   1, // Lines should be cleared, leaving only the last incomplete line
		},
		{
			name:       "Mixed blocks",
			placements: []string{"Q0", "S4", "Z1", "T3"},
			width:      10,
			height:     20,
			expected:   4,
		},
		{
			name:       "High placement",
			placements: []string{"I0"},
			width:      10,
			height:     10,
			expected:   1, // Single line at bottom, height from bottom is 1
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SimulateTetris(tt.placements, tt.width, tt.height)
			if result != tt.expected {
				t.Errorf("SimulateTetris() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestLineClearing(t *testing.T) {
	game := NewGame(4, 6)

	// Fill the bottom line completely
	game.grid[5][0] = true
	game.grid[5][1] = true
	game.grid[5][2] = true
	game.grid[5][3] = true

	// Add a block in the line above
	game.grid[4][0] = true

	fmt.Println("Before line clearing:")
	game.PrintGrid()

	linesCleared := game.clearLines()

	fmt.Println("After line clearing:")
	game.PrintGrid()

	if linesCleared != 1 {
		t.Errorf("Expected 1 line cleared, got %d", linesCleared)
	}

	// The block from line 4 should now be at line 5
	if !game.grid[5][0] {
		t.Errorf("Expected block to fall down after line clearing")
	}

	// Check the height from bottom
	heightFromBottom := game.getHighestY()
	if heightFromBottom != 1 {
		t.Errorf("Expected height from bottom to be 1, got %d", heightFromBottom)
	}
}

func TestCoordinateSystem(t *testing.T) {
	game := NewGame(4, 4)

	// Place a block at the very bottom
	game.grid[3][0] = true
	height := game.getHighestY()
	if height != 1 {
		t.Errorf("Block at bottom should have height 1 from bottom, got %d", height)
	}

	// Place a block one row up
	game.grid[2][0] = true
	height = game.getHighestY()
	if height != 2 {
		t.Errorf("Block one row up should have height 2 from bottom, got %d", height)
	}

	// Place a block at the very top
	game.grid[0][0] = true
	height = game.getHighestY()
	if height != 4 {
		t.Errorf("Block at top should have height 4 from bottom, got %d", height)
	}
}

// Example usage function for testing
func ExampleSimulateTetris() {
	// Test 1: Basic functionality
	fmt.Println("Test 1: Basic block placement")
	placements1 := []string{"Q0", "S4"}
	result1 := SimulateTetris(placements1, 10, 10)
	fmt.Printf("Result: %d\n", result1)

	// Test 2: Line clearing
	fmt.Println("Test 2: Line clearing")
	placements2 := []string{"I0", "I1", "I2", "I3", "I4", "I5"}
	result2 := SimulateTetris(placements2, 10, 10)
	fmt.Printf("Result: %d\n", result2)

	// Test 3: Complex stacking
	fmt.Println("Test 3: Complex stacking")
	placements3 := []string{"Q0", "Q2", "Q4", "Q6", "Q8", "Q0", "Q2", "Q4", "Q6", "Q8"}
	result3 := SimulateTetris(placements3, 10, 10)
	fmt.Printf("Result: %d\n", result3)
}
