# Tetris Simulation in Go

This is a Tetris simulation program written in Go that handles block placements and line clearing.

## Block Types

- **Q**: Cube (2x2)
- **Z**: Left facing Z shape (3x2)
- **S**: Right facing S shape (3x2)
- **T**: Downward facing T shape (3x2)
- **J**: Left facing L shape (3x2)
- **L**: Right facing L shape (3x2)
- **I**: Horizontal 4-long line (4x1)

## Block Placement Format

Each block placement is specified as a string in the format `[BlockType][XPosition]`, where:
- `BlockType` is one of the block types above (Q, Z, S, T, J, L, I)
- `XPosition` is the x-coordinate where the leftmost element of the block will be placed

Examples:
- `Q0`: Place a cube at x=0
- `S4`: Place an S block at x=4
- `I0`: Place a horizontal line at x=0

## How It Works

1. The program takes a list of block placements
2. For each placement, it finds the lowest valid position for the block
3. Blocks stack on top of each other when appropriate
4. Full lines are automatically cleared (like in Tetris)
5. The program returns the Y position of the highest block

## Usage

```go
placements := []string{"Q0", "S4", "Z1", "T3", "I0"}
width, height := 10, 20
highestY := SimulateTetris(placements, width, height)
fmt.Printf("Highest Y position: %d\n", highestY)
```

## Running the Program

```bash
go run main.go
```

The program includes example usage and visual debugging output to show how blocks are placed and lines are cleared.

## Key Features

- **Gravity**: Blocks fall to the lowest possible position
- **Line Clearing**: Complete horizontal lines are removed
- **Collision Detection**: Blocks cannot overlap or go out of bounds
- **Visual Debugging**: Optional grid printing for debugging
- **Flexible Grid Size**: Configurable width and height 