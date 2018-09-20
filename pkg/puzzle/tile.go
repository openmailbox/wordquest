package puzzle

import (
	"fmt"
)

// Tile - A single {x,y} coordinate within a Puzzle
type Tile struct {
	Value string `json:"value"`
	X     int    `json:"x"`
	Y     int    `json:"y"`
}

func (t Tile) String() string {
	return fmt.Sprintf("{'%v' @ %v,%v}", t.Value, t.X, t.Y)
}
