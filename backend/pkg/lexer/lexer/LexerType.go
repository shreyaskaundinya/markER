package lexer

import lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"

type Lexer struct {
	Name   string
	Input  string
	Tokens chan lexertoken.Token
	State  LexFn

	Start int
	Pos   int
	Width int
}