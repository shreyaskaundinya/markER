package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shreyaskaundinya/markER/backend/pkg/parser/parser"

	"github.com/gin-gonic/gin"
)

type ParserBody struct {
	Code string `json:"code"`
}

func HandleParser(c *gin.Context) {
	var response ParserBody
	c.Header("Content-Type", "application/json")
	if err := c.BindJSON(&response); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	
	fmt.Println(response.Code)

	out := parser.StartParsing(response.Code)
	
	json, _ := json.Marshal(out)

	fmt.Println(string(json))
	
	c.JSON(http.StatusAccepted, gin.H{"message": "hello world"})
}