package handlers

import (
	"fmt"
	"net/http"

	"github.com/shreyaskaundinya/markER/backend/pkg/lexer/lexer"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Code string
}

func HandleLexer(c *gin.Context) {
	var response Body
	c.Header("Content-Type", "application/json")
	if err := c.BindJSON(&response); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	fmt.Println(response.Code)
	lex := lexer.BeginLexing("Lexer", response.Code)
	
	state := lexer.LexBegin

	for state != nil {
		state = state(lex)
	}

	close(lex.Tokens)
	
	c.JSON(http.StatusAccepted, gin.H{"message": "hello world"})
	
	for tok := range lex.Tokens {
		fmt.Println(tok)
	}
}