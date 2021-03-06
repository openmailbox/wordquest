package puzzle

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

// Puzzle - Describes the structure of the word search puzzle
type Puzzle struct {
	Length int
	Width  int
	Tiles  []*Tile
	Words  []Word // Set of all possible answers
	Solved []Word // Set of already-discovered answers
}

// Add - Add a new word to the puzzle
func (p *Puzzle) Add(word Word) {
	directions := []direction{leftToRight, upToDown, downToUp}
	placed := false

	// Randomize the iteration order of the tiles to have less deterministic word placement
	for _, tile := range p.Tiles {
		for _, direction := range directions {
			var tiles []*Tile

			tiles, err := p.checkWordFit(word, tile, 0, direction, tiles)

			if err != nil {
				//fmt.Printf("Unable to add word: %v\n", err)
				continue
				// TODO: If the first tial is the one that can't fit, don't bother checking other directions
			}

			for i, newTile := range tiles {
				tileInPuzzle, err := p.GetTial(newTile.X, newTile.Y)
				if err != nil {
					panic(err)
				}

				tileInPuzzle.Value = string(word.Value[i])
				word.Tiles = append(word.Tiles, tileInPuzzle)

				fmt.Printf("{%v,%v} = %v, ", tileInPuzzle.X, tileInPuzzle.Y, string(word.Value[i]))
			}

			fmt.Println("")
			placed = true
			p.Words = append(p.Words, word)
			break
		}

		if placed {
			break
		}
	}
}

// Attempts to fit a word into the puzzle from a starting tile.
// Returns traversed tiles and an error if the word does not fit.
func (p Puzzle) checkWordFit(word Word, currentTile *Tile, wordIndex int, traverse direction, visitedTiles []*Tile) ([]*Tile, error) {
	if len(currentTile.Value) == 0 || strings.Compare(currentTile.Value, string(word.Value[wordIndex])) == 0 {
		visitedTiles = append(visitedTiles, currentTile)

		if len(word.Value) > wordIndex+1 {
			nextTile, err := traverse(p, currentTile)
			if err != nil {
				return visitedTiles, err // reached edge of puzzle
			}

			return p.checkWordFit(word, nextTile, wordIndex+1, traverse, visitedTiles)
		}

		return visitedTiles, nil
	}

	message := fmt.Sprintf("tile '%v' does not match '%v'", currentTile.Value, string(word.Value[wordIndex]))
	return visitedTiles, errors.New(message)
}

// Fill - Fill in all remaining blank tiles with random letters
func (p *Puzzle) Fill() {
	var words []string

	for _, word := range p.Words {
		words = append(words, word.Value)
	}

	letterList := strings.Join(words, "")
	length := len(letterList)

	for _, tile := range p.Tiles {
		if len(tile.Value) == 0 {
			tile.Value = string(letterList[rand.Intn(length)])
		}
	}
}

// GetTial - Find the tial in the current puzzle by coordinates
func (p *Puzzle) GetTial(x int, y int) (foundTile *Tile, e error) {
	for _, tile := range p.Tiles {
		if tile.X == x && tile.Y == y {
			return tile, nil
		}
	}

	message := fmt.Sprintf("tile {%v,%v} not found in puzzle.", x, y)
	return &Tile{}, errors.New(message)
}

// SubmitAnswer - Accept a new answer if it matches a word in the answer list and it has not already been submitted
func (p *Puzzle) SubmitAnswer(submission Word) bool {
	for _, word := range p.Words {
		if word.Compare(submission) {
			for _, submittedWord := range p.Solved {
				if submittedWord.Compare(submission) {
					return false
				}
			}
			p.Solved = append(p.Solved, word)

			return true
		}
	}

	return false
}

// MarshalJSON - JSON representation of the puzzle suitable for clients (without solutions)
func (p Puzzle) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Length    int     `json:"length"`
		Width     int     `json:"width"`
		Tiles     []*Tile `json:"tiles"`
		Solutions []Word  `json:"solutions"`
	}{
		Length:    p.Length,
		Width:     p.Width,
		Tiles:     p.Tiles,
		Solutions: p.Solved,
	})
}

// Initialize - Init the puzzle with blank tiles based on the provided Length and Width
func (p *Puzzle) Initialize() {
	for i := 0; i < p.Length; i++ {
		for j := 0; j < p.Width; j++ {
			p.Tiles = append(p.Tiles, &Tile{"", j, i})
		}
	}
}

func (p Puzzle) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("A %vx%v puzzle containing %v words:\n", p.Length, p.Width, len(p.Words)))

	for _, tile := range p.Tiles {
		buffer.WriteString(fmt.Sprintf(" %v ", tile.Value))

		if tile.X >= p.Width-1 {
			buffer.WriteString(fmt.Sprintln())
		}
	}

	return buffer.String()
}
