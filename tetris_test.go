package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_SimulateTetris(t *testing.T) {
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
			expected:   0,
		},
		{
			name:       "Single Q block",
			placements: []string{"Q0"},
			width:      10,
			height:     20,
			expected:   2,
		},
		{
			name:       "Single I block",
			placements: []string{"I0"},
			width:      10,
			height:     10,
			expected:   1,
		},
		{
			name:       "Single Z block",
			placements: []string{"Z0"},
			width:      10,
			height:     10,
			expected:   2,
		},
		{
			name:       "Single J block",
			placements: []string{"J0"},
			width:      10,
			height:     10,
			expected:   3,
		},
		{
			name:       "Single L block",
			placements: []string{"L0"},
			width:      10,
			height:     10,
			expected:   3,
		},
		{
			name:       "Single T block",
			placements: []string{"T0"},
			width:      10,
			height:     10,
			expected:   2,
		},
		{
			name:       "Single S block",
			placements: []string{"S0"},
			width:      10,
			height:     10,
			expected:   2,
		},
		{
			name:       "Stacked blocks",
			placements: []string{"Q0", "Q0"},
			width:      10,
			height:     20,
			expected:   4, // Two cubes stacked, height from bottom is 4
		},
		{
			name:       "Line clearing - 2 lines",
			placements: []string{"I0", "I4", "I0", "I4", "Q8", "I4", "I5"},
			width:      10,
			height:     20,
			expected:   2, // Lines should be cleared, leaving only the last 2 lines
		},
		{
			name:       "Line clearing - 1 line",
			placements: []string{"I0", "I4", "Q8"},
			width:      10,
			height:     20,
			expected:   1, // Line should be cleared, leaving only last line
		},
		{
			name:       "Example 2",
			placements: []string{"T1", "Z3", "I4"},
			width:      10,
			height:     100,
			expected:   4,
		},
		{
			name:       "Example 3",
			placements: []string{"Q0", "I2", "I6", "I0", "I6", "I6", "Q2", "Q4"},
			width:      10,
			height:     100,
			expected:   3,
		},
		{
			name:       "T block with line clearing between Q blocks",
			placements: []string{"Q0", "Q3", "T1"},
			width:      5,
			height:     10,
			expected:   2,
		},
		{
			name:       "T block with line clearing between L, J blocks",
			placements: []string{"L0", "J3", "T1"},
			width:      5,
			height:     10,
			expected:   1,
		},
		{
			name:       "Multiple lines cleared",
			placements: []string{"I0", "I8", "I0", "I8", "I0", "I8", "I4", "I4", "I4"},
			width:      12,
			height:     20,
			expected:   0,
		},
		{
			name:       "Fill height with I blocks",
			placements: []string{"I0", "I0", "I0", "I0", "I0"},
			width:      10,
			height:     5,
			expected:   5,
		},
		{
			name:       "Out of bounds - height exceeded",
			placements: []string{"L0", "L0", "L0", "L0"},
			width:      10,
			height:     10,
			expected:   -1,
		},
		{
			name:       "Out of bounds - width exceeded",
			placements: []string{"I4"},
			width:      5,
			height:     10,
			expected:   -1,
		},
		{
			name:       "Invalid block type",
			placements: []string{"K4"},
			width:      5,
			height:     10,
			expected:   -1,
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

func Test_clearLines(t *testing.T) {
	game := NewGame(4, 6)

	// Fill the bottom line completely with an "I" block
	y := game.findPlacement(blocks["I"], 0)
	game.placeBlock(blocks["I"], 0, y)

	// Add a block in the line above
	y = game.findPlacement(blocks["Q"], 0)
	game.placeBlock(blocks["Q"], 0, y)

	lineAbove := make([]bool, game.width)
	copy(lineAbove, game.grid[game.height-2])

	fmt.Println("Before line clearing:")
	game.PrintGrid()

	game.clearLines()

	fmt.Println("After line clearing:")
	game.PrintGrid()

	// assert that the line above has been moved down
	if !reflect.DeepEqual(game.grid[game.height-1], lineAbove) {
		t.Errorf("Line above has not been moved down correctly")
	}

	// Check height from the bottom
	heightFromBottom := game.getHighestY()
	if heightFromBottom != 2 {
		t.Errorf("Expected height from bottom to be 1, got %d", heightFromBottom)
	}
}
