package internal

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/openmailbox/wordquest/pkg/puzzle"
)

type game struct {
	players []player
}

type player struct {
	connection *websocket.Conn
}

const localAddress = ":8082"

var currentGame game
var currentPuzzle puzzle.Puzzle
var upgrader = websocket.Upgrader{}

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
}

func handleUpdates(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		// TODO: Tell browser we can't upgrade so fallback to polling
        log.Printf("Unable to establish websocket connection: %v\n", err)
        return
	}

	currentGame.newPlayer(conn)
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v from %v\n", r.Method, r.URL, r.RemoteAddr)
		handler.ServeHTTP(w, r)
		log.Printf("Completed %v %v\n", http.StatusOK, http.StatusText(http.StatusOK))
	})
}

func (g *game) newPlayer(newConn *websocket.Conn) {
	log.Printf("New connection from %v\n", newConn.RemoteAddr())

	var newPlayer = player{connection: newConn}
	currentGame.players = append(currentGame.players, newPlayer)

	log.Printf("%v total players", len(currentGame.players))

	// defer server.closeConnection(newPlayer)

	go func() {
		for {
			_, message, err := newPlayer.connection.ReadMessage()
			if err != nil {
				log.Println(err)
				break
			}

			log.Printf("%v\n", string(message))
		}
	}()
}

// StartServer - Starts the web server
func StartServer(newPuzzle puzzle.Puzzle) {
	currentPuzzle = newPuzzle
	currentGame = game{}

	log.Printf("Listening on %v...\n", localAddress)
	http.Handle("/", http.FileServer(http.Dir("../../web/static")))
	http.HandleFunc("/puzzle", handlePuzzle)
	http.HandleFunc("/updates", handleUpdates)
	log.Fatal(http.ListenAndServe(localAddress, logRequest(http.DefaultServeMux)))
}
