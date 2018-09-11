package internal

import (
    "fmt"
	"log"
	"net/http"
)

const localAddress = ":8082"

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v from %v\n", r.Method, r.URL, r.RemoteAddr)
		handler.ServeHTTP(w, r)
	})
}

// StartServer - Starts the web server
func StartServer() {
    fmt.Printf("Listening on %v...\n", localAddress)
	http.Handle("/", http.FileServer(http.Dir("../../web/static")))
	log.Fatal(http.ListenAndServe(localAddress, logRequest(http.DefaultServeMux)))
}
