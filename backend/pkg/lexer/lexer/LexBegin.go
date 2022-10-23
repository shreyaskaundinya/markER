package lexer

import (
	"strings"

	lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"
)


func LexBegin(l *Lexer) LexFn {
	// skip the whitespace
	l.SkipWhitespace()

	// the start token can either be a create token or a relation token
	if strings.HasPrefix(l.InputToEnd(), lexertoken.CREATE) {
		return LexCreate
		} else if strings.HasPrefix(l.InputToEnd(), lexertoken.RELN) {
		return LexReln
	} else {
	return nil
	}
}