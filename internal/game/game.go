package game

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/openmailbox/wordquest/pkg/puzzle"
)

// Game - A type representing a running word search puzzle game
type Game struct {
	currentPuzzle puzzle.Puzzle
	players       []player
	webServer     server
}

func (g *Game) newPlayer(newConn *websocket.Conn) {
	log.Printf("New connection from %v\n", newConn.RemoteAddr())

	var newPlayer = player{connection: newConn}
	g.players = append(g.players, newPlayer)

	log.Printf("%v total players", len(g.players))

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

// Start - Primary interface for starting a new game.
func (g *Game) Start() {
	log.Println("Starting new game...")

	g.currentPuzzle = puzzle.GeneratePuzzle()

	g.webServer = server{}
	g.webServer.Start(g)
}
