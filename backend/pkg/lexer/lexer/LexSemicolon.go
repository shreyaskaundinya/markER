package lexer

import (
	"strings"

	lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"
)

func LexSemicolon(l *Lexer) LexFn {
	l.SkipWhitespace()

	if(strings.HasPrefix(l.InputToEnd(), lexertoken.SEMICOLON)) {
		l.Pos += len(lexertoken.SEMICOLON)
		l.Emit(lexertoken.TOKEN_SEMICOLON, true)
	}

	return LexBegin
}