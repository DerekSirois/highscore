package server

import (
	"database/sql"
	"highscore/internal/services/game"
	"highscore/internal/services/score"
	"highscore/internal/services/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Addr string
	Db   *sql.DB
}

func New(addr string, db *sql.DB) *Server {
	return &Server{
		Addr: addr,
		Db:   db,
	}
}

func (s *Server) Run() {
	router := mux.NewRouter()

	userStore := user.NewStore(s.Db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(router)

	gameStore := game.NewStore(s.Db)
	gameHandler := game.NewHandler(gameStore, userStore)
	gameHandler.RegisterRoutes(router)

	scoreStore := score.NewStore(s.Db)
	scoreHandler := score.NewHandler(scoreStore, userStore)
	scoreHandler.RegisterRoutes(router)

	log.Printf("Serving on %s\n", s.Addr)
	log.Fatal(http.ListenAndServe(s.Addr, router))
}
