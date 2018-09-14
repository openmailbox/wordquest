package game

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const localAddress = ":8082" // TODO: parameterize this

type server struct {
	currentGame *Game
	upgrader    websocket.Upgrader
}

func (s *server) handlePuzzle(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(s.currentGame.currentPuzzle)
}

func (s *server) handleUpdates(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		// TODO: Tell browser we can't upgrade so fallback to polling
		log.Printf("Unable to establish websocket connection: %v\n", err)
		return
	}

	s.currentGame.newPlayer(conn)
}

func (s *server) logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v from %v\n", r.Method, r.URL, r.RemoteAddr)
		handler.ServeHTTP(w, r)
		log.Printf("Completed %v %v\n", http.StatusOK, http.StatusText(http.StatusOK))
	})
}

// StartServer - Starts the web server
func (s *server) Start(game *Game) {
	s.currentGame = game

	log.Printf("Listening on %v...\n", localAddress)

	http.Handle("/", http.FileServer(http.Dir("../../web/static/dist")))
	http.HandleFunc("/puzzle", s.handlePuzzle)
	http.HandleFunc("/updates", s.handleUpdates)

	log.Fatal(http.ListenAndServe(localAddress, s.logRequest(http.DefaultServeMux)))
}
