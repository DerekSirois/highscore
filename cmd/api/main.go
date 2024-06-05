package main

import (
	"highscore/internal/db"
	"highscore/internal/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db := db.InitDb()

	s := server.New(":8080", db)
	s.Run()
}
