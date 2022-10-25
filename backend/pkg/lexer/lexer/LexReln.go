package lexer

import (
	"fmt"
	lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"
)

/* reln token */

func LexReln(l *Lexer) LexFn {
	l.pos += len(lexertoken.RELN)
	l.Emit(lexertoken.TOKEN_RELN, true)

	fmt.Println("TAG: RELN")

	l.SkipWhitespace()
	return RelnTable
}