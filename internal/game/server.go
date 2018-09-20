package game

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/openmailbox/wordquest/pkg/puzzle"
)

const localAddress = ":8082" // TODO: parameterize this

type server struct {
	currentGame *Game
	upgrader    websocket.Upgrader
}

func (s *server) handlePuzzle(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(s.currentGame.currentPuzzle)
	log.Printf("Completed %v %v\n", http.StatusOK, http.StatusText(http.StatusOK))
}

func (s *server) handleSubmit(w http.ResponseWriter, r *http.Request) {
	var status int
	var submission puzzle.Word

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&submission)
	if err != nil {
		log.Printf("Invalid submission: %v\n", err)
		status = http.StatusUnprocessableEntity
	} else {
		log.Printf("Submission: %v", submission)

		if s.currentGame.currentPuzzle.SubmitAnswer(submission) {
			status = http.StatusCreated
		} else {
			status = http.StatusNotFound
		}
	}

	w.WriteHeader(status)
	log.Printf("Completed %v %v\n", status, http.StatusText(status))
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
	})
}

// StartServer - Starts the web server
func (s *server) Start(game *Game) {
	s.currentGame = game

	log.Printf("Listening on %v...\n", localAddress)

	http.Handle("/", http.FileServer(http.Dir("../../web/static/dist")))
	http.HandleFunc("/puzzle", s.handlePuzzle)
	http.HandleFunc("/updates", s.handleUpdates)
	http.HandleFunc("/submit", s.handleSubmit)

	log.Fatal(http.ListenAndServe(localAddress, s.logRequest(http.DefaultServeMux)))
}
