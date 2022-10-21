package lexer

import (
	lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"
)
func BeginLexing(name string, input string) *Lexer {
	return &Lexer{
		Name:   name,
		Input:  input,
		State:  LexBegin,
		Tokens: make(chan lexertoken.Token, 3),
	}
}

