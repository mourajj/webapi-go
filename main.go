package main

import (
	"webapi-go/database"
	"webapi-go/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}
