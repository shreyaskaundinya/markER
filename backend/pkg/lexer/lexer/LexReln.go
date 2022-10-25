package lexer

import (
	lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"
)

/* reln token */

func LexReln(l *Lexer) LexFn {
	l.Pos += len(lexertoken.RELN)
	l.Emit(lexertoken.TOKEN_RELN, true)

	l.SkipWhitespace()
	return LexIdentifying
}