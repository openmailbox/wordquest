package puzzle

import (
	"errors"
)

// Describes how to traverse a puzzle in a particular direction
type direction func(puzzle Puzzle, currentTile *Tile) (nextTile *Tile, e error)

func downToUp(puzzle Puzzle, currentTile *Tile) (nextTile *Tile, e error) {
	for _, tile := range puzzle.Tiles {
		if tile.Y == currentTile.Y-1 && tile.X == currentTile.X {
			return tile, nil
		}
	}

	return &Tile{}, errors.New("invalid tile location")
}

func leftToRight(puzzle Puzzle, currentTile *Tile) (nextTile *Tile, e error) {
	for _, tile := range puzzle.Tiles {
		if tile.X == currentTile.X+1 && tile.Y == currentTile.Y {
			return tile, nil
		}
	}

	return &Tile{}, errors.New("invalid tile location")
}

func upToDown(puzzle Puzzle, currentTile *Tile) (nextTile *Tile, e error) {
	for _, tile := range puzzle.Tiles {
		if tile.Y == currentTile.Y+1 && tile.X == currentTile.X {
			return tile, nil
		}
	}

	return &Tile{}, errors.New("invalid tile location")
}
