package lexer

import (
	"strings"

	lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"
)

func LexAttributeComma(l *Lexer) LexFn {
	l.SkipWhitespace()

	if strings.HasPrefix(l.InputToEnd(), lexertoken.COMMA) {
		l.Pos += len(lexertoken.COMMA)
		l.Emit(lexertoken.TOKEN_COMMA, true)
		return LexAttributeName
	}

	// if no comma, end of attributes
	return LexTableBracketClose
}