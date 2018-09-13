package internal

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/openmailbox/wordquest/pkg/puzzle"
)

const localAddress = ":8082"

var currentPuzzle puzzle.Puzzle

func handlePuzzle(w http.ResponseWriter, r *http.Request) {
	var tmpPuzzle struct {
		Length int            `json:"length"`
		Width  int            `json:"width"`
		Tiles  []*puzzle.Tile `json:"tiles"`
	}

	tmpPuzzle.Length = currentPuzzle.Length
	tmpPuzzle.Width = currentPuzzle.Width
	tmpPuzzle.Tiles = currentPuzzle.Tiles

	json.NewEncoder(w).Encode(tmpPuzzle)
	log.Printf("Completed %v %v\n", http.StatusOK, http.StatusText(http.StatusOK))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v from %v\n", r.Method, r.URL, r.RemoteAddr)
		handler.ServeHTTP(w, r)
	})
}

// StartServer - Starts the web server
func StartServer(newPuzzle puzzle.Puzzle) {
	currentPuzzle = newPuzzle

	log.Printf("Listening on %v...\n", localAddress)
	http.Handle("/", http.FileServer(http.Dir("../../web/static")))
	http.HandleFunc("/puzzle", handlePuzzle)
	log.Fatal(http.ListenAndServe(localAddress, logRequest(http.DefaultServeMux)))
}
