package server

import (
	"github.com/shreyaskaundinya/markER/backend/pkg/http-server/handlers"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()

	r.GET("/ping", handlers.ExampleHandler)
	r.POST("/lex", handlers.HandleLexer)
	r.POST("/parse", handlers.HandleParser)
	
	r.Run("localhost:5050") 
}
