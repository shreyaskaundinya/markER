package lexer

import (
	lexertoken "go-marker-backend/pkg/lexer/lexertoken"
)

func BeginLexing(name string, input string) *Lexer {
	return &Lexer{
		Name:   name,
		Input:  input,
		State:  LexBegin,
		Tokens: make(chan lexertoken.Token, 3),
	}
}

