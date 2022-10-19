package lexer

import lexertoken "go-marker-backend/pkg/lexer/lexertoken"

type Lexer struct {
	Name   string
	Input  string
	Tokens chan lexertoken.Token
	State  LexFn

	Start int
	Pos   int
	Width int
}