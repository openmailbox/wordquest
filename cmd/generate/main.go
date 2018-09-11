package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"

	"github.com/openmailbox/wordquest/pkg/puzzle"
)

func main() {
	secs := time.Now().Unix()
	fmt.Printf("Random seed: %v", secs)
	rand.Seed(secs)

	dat, err := ioutil.ReadFile("../../data/words-computer.txt")
	if err != nil {
		panic(err)
	}

	words := strings.Split(string(dat), "\n")

	var newPuzzle puzzle.Puzzle
	newPuzzle.Length = 10
	newPuzzle.Width = 10

	fmt.Printf("Initializing a %vx%v puzzle...\n", newPuzzle.Length, newPuzzle.Width)

	newPuzzle.Initialize()

	fmt.Println("Filling puzzle with words:")

	for _, word := range words {
		var newWord puzzle.Word
		newWord.Value = word

		newPuzzle.Add(newWord)
	}

	newPuzzle.Fill()

	fmt.Println(newPuzzle)
}
