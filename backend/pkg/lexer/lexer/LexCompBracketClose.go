package lexer

import (
	"strings"

	lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"
)

func LexCompBracketClose(l *Lexer) LexFn {
	l.SkipWhitespace()

	if (strings.HasPrefix(l.InputToEnd(), lexertoken.RIGHT_BRACKET)){
		l.Pos++
		l.Emit(lexertoken.TOKEN_RIGHT_BRACKET, true)
		return LexAttributeComma
	}
	
	// throw error
	return nil
}