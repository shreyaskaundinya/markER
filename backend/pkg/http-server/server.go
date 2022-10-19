package server

import (
	"go-marker-backend/pkg/http-server/handlers"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()

	r.GET("/ping", handlers.ExampleHandler)
	
	r.Run() // listen and serve on 0.0.0.0:8080
}
