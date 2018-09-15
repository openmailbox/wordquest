package puzzle

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// GeneratePuzzle - A factory for making new puzzles
func GeneratePuzzle() Puzzle {
	secs := time.Now().Unix()
	fmt.Printf("Random seed: %v", secs)
	rand.Seed(secs)

	dat, err := ioutil.ReadFile("../../test/test-wordlist-computer.txt")
	if err != nil {
		panic(err)
	}

	words := strings.Split(string(dat), "\r\n")

	var newPuzzle Puzzle
	newPuzzle.Length = 10
	newPuzzle.Width = 10

	fmt.Printf("Initializing a %vx%v puzzle...\n", newPuzzle.Length, newPuzzle.Width)

	newPuzzle.Initialize()

	fmt.Println("Filling puzzle with words:")

	for _, word := range words {
		var newWord Word
		newWord.Value = word

		newPuzzle.Add(newWord)
	}

	newPuzzle.Fill()

	return newPuzzle
}
