package puzzle

import (
	"bytes"
	"fmt"
)

// Word - A up->down, down->up, or left->right block of cells
type Word struct {
	Tiles []*Tile `json:"tiles"`
	Value string  `json:"value"`
}

// Compare - If the value and tiles of two words are the same, return true
func (w Word) Compare(other Word) bool {
	if w.String() == other.String() {
		return true
	}

	return false
}

func (w Word) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("'%v' composed from %v tiles: ", w.Value, len(w.Tiles)))

	for _, tile := range w.Tiles {
		buffer.WriteString(fmt.Sprintf("%v ", tile))
	}

	return buffer.String()
}
