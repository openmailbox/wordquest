package puzzle

// Word - A up->down, down->up, or left->right block of cells
type Word struct {
	Tiles []*Tile
	Value string
}
