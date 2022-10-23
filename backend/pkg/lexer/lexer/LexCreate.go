package lexer

import (
	lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"
)

/*
	State for the Create Token
*/ 
func LexCreate(l *Lexer) LexFn {
	l.Pos += len(lexertoken.CREATE)
	l.Emit(lexertoken.TOKEN_CREATE, true)

	l.SkipWhitespace()
	return LexTable
}
