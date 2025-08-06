# Tetris Simulation in Go

This is a Tetris simulation program written in Go that handles block placements and line clearing.

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

## How It Works

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

### Option 1: Run the program using ```go run```

```
go run . < input.txt >  output.txt
```

### Option 2: Build an executable file using ```go build``` and run the executable file

```bash
go build -o tetris
```

This command will build the source code and create an executable named ```tetris```. Running this will ensure that the executable will be compatible with your OS.

```bash
./tetris < input.txt >  output.txt
```

The input will be read line by line from ```input.txt``` and the output will be written - line by line - in ```output.txt```.

The main difference between the 2 methods is that ```go run``` will compile the program on each run whereas with the second method, it will be compiled using ```go build``` and the executable can be ran repeatedly without compiling the code again.

## Time Complexity

The time complexity of the ```SimulateTetris``` function is ```O(P * H * W)``` where P is the number of placements, H is the grid height and W is the grid width.

The loops in ```placeBlock``` and ```canPlace``` are not affected by input size and can be considered constant time.


## Testing / debugging
A ```PrintGrid``` function was used for debugging purposes but is not used in the main function.

```bash
go test -v
```

This command will run the tests in ```tetris_test.go```, printing out the test names, their status and any debug logging.


