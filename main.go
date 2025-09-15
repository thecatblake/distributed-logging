package main

import (
	"github.com/thecatblake/distributed-logging/internal/server"
	"log"
)

func main() {
	srv := server.NewHTTPServer(":8000")
	log.Fatal(srv.ListenAndServe())
}