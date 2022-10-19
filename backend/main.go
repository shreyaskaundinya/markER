package main

import (
	"fmt"

	server "go-marker-backend/pkg/http-server"
)

func main() {
	fmt.Println("Hello World")
	server.RunServer()
}