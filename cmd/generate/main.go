package main

import (
	"fmt"

	"github.com/openmailbox/wordquest/pkg/wordpuzzle"
)

func main() {
	puzzle := wordpuzzle.MakePuzzle(10, 10)

	fmt.Printf("New %vx%v puzzle created.\n", puzzle.Length, puzzle.Width)
}
