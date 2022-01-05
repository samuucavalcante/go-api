package main

import (
	"github.com/samuucavalcante/go-api/database"
	"github.com/samuucavalcante/go-api/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()

	server.Run()
}
