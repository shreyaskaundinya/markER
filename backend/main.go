package main

import (
	"fmt"

	server "github.com/shreyaskaundinya/markER/backend/pkg/http-server"
)

func main() {
	fmt.Println("Hello World")
	server.RunServer()
}