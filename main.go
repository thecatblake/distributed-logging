package main

import (
	"github.com/thecatblake/distributed-logging/internal/server"
	"fmt"
)

func main() {
	log := &server.Log{}
	log.Append(server.Record{Value: []byte("Hello")})
	log.Append(server.Record{Value: []byte("World")})
	fmt.Println(log.Read(0))
	fmt.Println(log.Read(1))
	fmt.Println(log.Read(2))
}