package main

import "webapi-go/server"

func main() {
	server := server.NewServer()
	server.Run()
}
