package lexer

import (
	"strings"

	lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"
)

func LexTableBracketClose(l *Lexer) LexFn {
	l.Emit(lexertoken.TOKEN_COMMA, true)
	l.SkipWhitespace()

	if (strings.HasPrefix(l.InputToEnd(), lexertoken.RIGHT_BRACKET)){
		l.Pos++
		l.Emit(lexertoken.TOKEN_RIGHT_BRACKET, true)
		return LexSemicolon
	}

	// throw error
	return nil
}