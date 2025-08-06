# Tetris Simulation in Go

This is a Tetris command line game in Golang that is 90% AI generated. My approach to writing this game was to split the requirements into smaller milestones I could generate with GitHub Copilot and then verify manually and with automated tests.

I use the same approach when prototyping anything with AI.
Here are the milestones:
1. Stack blocks in a grid of predefined width and height, clear full lines and end game when top line is reached
2. Add command line input, randomly pick a block and prompt the user for an X coordinate to place the block
3. Add timer that simulates a block dropping, calculate based on grid height. If user misses the timer, a random position is chosen
4. Add rotation (?)

## Block Types

- **Q** : Cube (2x2)
- **Z** : Left facing Z shape (3x2)
- **S** : Right facing S shape (3x2)
- **T** : Downward facing T shape (3x2)
- **J** : Left facing L shape (3x2)
- **L** : Right facing L shape (3x2)
- **I** : Horizontal 4-long line (4x1)

## Block Placement Format

Each block placement is specified as a string in the format `[BlockType][XPosition]`, where:
- `BlockType` is one of the block types above (Q, Z, S, T, J, L, I)
- `XPosition` is the x-coordinate where the leftmost element of the block will be placed

Examples:
- `Q0`: Place a cube at x=0
- `S4`: Place an S block at x=4
- `I0`: Place a horizontal line at x=0

## How SimulateTetris Works

1. The program takes a list of comma separated block placements.
2. For each placement, it finds the lowest valid position for the block.
3. Blocks stack on top of each other when appropriate.
4. Full lines are automatically cleared on every block placement.
5. The program prints the Y position of the highest block element. If an invalid entry is given, the program prints -1.

## SimulateTetris example usage

```go
placements := []string{"Q0", "S4", "Z1", "T3", "I0"}
width, height := 10, 20
highestY := SimulateTetris(placements, width, height)
fmt.Printf("Highest Y position: %d\n", highestY)
```

## Running the Program (requires Go >1.21)


```
go run .
```



## Time Complexity

The time complexity of the ```SimulateTetris``` function is ```O(P * H * W)``` where P is the number of placements, H is the grid height and W is the grid width.

The loops in ```placeBlock``` and ```canPlace``` are not affected by input size and can be considered constant time.


## Testing / debugging
A ```PrintGrid``` function was used for debugging purposes but is not used in the main function.

```bash
go test -v
```

This command will run the tests in ```tetris_test.go```, printing out the test names, their status and any debug logging.


