package main

import "highscore/internal/server"

func main() {
	s := server.New(":8080")
	s.Run()
}
