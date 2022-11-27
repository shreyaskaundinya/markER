package server

import (
	"time"

	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shreyaskaundinya/markER/backend/pkg/http-server/handlers"
)

func RunServer() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/ping", handlers.ExampleHandler)
	r.POST("/lex", handlers.HandleLexer)
	r.POST("/parse", handlers.HandleParser)
	
	r.Run("localhost:5050") 
}
