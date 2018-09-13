package puzzle

// Tile - A single {x,y} coordinate within a Puzzle
type Tile struct {
	Value string `json:"value"`
	X     int    `json:"x"`
	Y     int    `json:"y"`
}
