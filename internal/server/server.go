package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Addr string
	Db   *sql.DB
}

func New(addr string) *Server {
	return &Server{
		Addr: addr,
		Db:   nil, // Temp
	}
}

func (s *Server) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome the the highscore leaderboard"))
	}) // Temp

	log.Printf("Serving on %s\n", s.Addr)
	log.Fatal(http.ListenAndServe(s.Addr, router))
}
