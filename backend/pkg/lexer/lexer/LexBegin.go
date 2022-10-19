package lexer

import (
	lexertoken "go-marker-backend/pkg/lexer/lexertoken"
	"strings"
)


func LexBegin(l *Lexer) LexFn {
	l.SkipWhitespace()

	if strings.HasPrefix(l.InputToEnd(), lexertoken.CREATE) {
		return LexCreate
	} else if strings.HasPrefix(l.InputToEnd(), lexertoken.RELN) {
		return LexReln
	} else {
		return nil
	}
}