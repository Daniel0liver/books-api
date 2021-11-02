package main

import "github.com/daniel0liver/books-api/server"

func main() {
	server := server.NewServer()

	server.Run()
}
