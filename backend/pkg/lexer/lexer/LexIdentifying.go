package lexer

import (
	"strings"

	lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"
)

func LexIdentifying(l *Lexer) LexFn {
	l.SkipWhitespace()

	if(strings.HasPrefix(l.InputToEnd(), lexertoken.IDENTIFYING)) {
		l.Pos += len(lexertoken.IDENTIFYING)
		l.Emit(lexertoken.TOKEN_IDENTIFYING, true)
	}

	return LexTable1
}




