package main

import (
    "github.com/openmailbox/wordquest/internal"
    "github.com/openmailbox/wordquest/pkg/puzzle"
)

func main() {
    newPuzzle := puzzle.GeneratePuzzle()
    internal.StartServer(newPuzzle)
}