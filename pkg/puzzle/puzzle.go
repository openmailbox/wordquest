package puzzle

import (
	"errors"
	"fmt"
	"strings"
)

// Puzzle - Describes the structure of the word search puzzle
type Puzzle struct {
	Length int
	Width  int
	Words  []Word
	Tiles  []Tile
}

// Add - Add a new word to the puzzle
func (p *Puzzle) Add(word Word) {
	directions := []direction{leftToRight, upToDown, downToUp}

	fmt.Println(len(p.Tiles))
	for _, tile := range p.Tiles {
		for _, direction := range directions {
			var tiles []*Tile

			slot, err := p.checkWordFit(word, &tile, 0, direction, tiles)

			if err != nil {
				fmt.Printf("Unable to add word: %v\n", err)
				continue
			}

			for i, changingTile := range slot {
				changingTile.Value = string(word.Value[i])
				fmt.Printf("{%v,%v} = %v", changingTile.X, changingTile.Y, string(word.Value[i]))
			}
		}
	}
}

// Attempts to fit a word into the puzzle from a starting tile.
// Returns traversed tiles and an error if the word does not fit.
func (p Puzzle) checkWordFit(word Word, currentTile *Tile, wordIndex int, traverse direction, visitedTiles []*Tile) ([]*Tile, error) {
	if len(currentTile.Value) == 0 || strings.Compare(currentTile.Value, string(word.Value[wordIndex])) == 0 {
		if len(word.Value)-1 > wordIndex {
			nextTile, err := traverse(p, currentTile)
			if err != nil {
				return visitedTiles, err // reached edge of puzzle
			}

			visitedTiles = append(visitedTiles, currentTile)
			p.checkWordFit(word, nextTile, wordIndex+1, traverse, visitedTiles)
		} else {
			return visitedTiles, nil
		}
	}

	message := fmt.Sprintf("tile '%v' does not match '%v'", currentTile.Value, string(word.Value[wordIndex]))
	return visitedTiles, errors.New(message)
}

// Initialize - Init the puzzle with blank tiles based on the provided Length and Width
func (p *Puzzle) Initialize() {
	for i := 0; i < p.Length; i++ {
		for j := 0; j < p.Width; j++ {
			p.Tiles = append(p.Tiles, Tile{"", i, j})
		}
	}
}
