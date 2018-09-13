package main

import (
	"fmt"

	"github.com/openmailbox/wordquest/pkg/puzzle"
)

func main() {
	newPuzzle := puzzle.GeneratePuzzle()
	fmt.Println(newPuzzle)
}
