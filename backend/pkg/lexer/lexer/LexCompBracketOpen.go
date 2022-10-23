package lexer

import (
	"strings"

	lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"
)

func LexCompBracketOpen(l *Lexer) LexFn {
	l.SkipWhitespace()

	if (strings.HasPrefix(l.InputToEnd(), lexertoken.LEFT_BRACKET)){
		l.Pos++
		l.Emit(lexertoken.TOKEN_LEFT_BRACKET, true)
		return LexCompAttributeName
	}
	
	// throw error
	return nil
}