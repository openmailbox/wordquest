package wordpuzzle

// Puzzle - The representation of a word puzzle
type Puzzle struct {
	Length int
	Width  int
	Tiles  []Tile
}

// Tile - A single {x,y} coordinate within a Puzzle
type Tile struct {
	X     int
	Y     int
	Value string
}

// Slot - A up->down, down->up, or left->right block of cells
type Slot struct {
	Tiles []Tile
}

// MakePuzzle - Generate a new puzzle with the specified length and width
func MakePuzzle(length int, width int) Puzzle {
	puzzle := Puzzle{Length: length, Width: width}

	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			puzzle.Tiles = append(puzzle.Tiles, Tile{X: i, Y: j})
		}
	}

	return puzzle
}