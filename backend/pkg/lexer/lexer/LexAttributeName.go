package lexer

import (
	"unicode"

	lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"
)

func LexAttributeName(l *Lexer) LexFn {
	l.SkipWhitespace()

	var ch rune

	for {
		ch = rune(l.Input[l.Pos])

		if unicode.IsSpace(ch) {
			l.Emit(lexertoken.TOKEN_ATTR_NAME, true)
			return LexAttributeType
		}

		if l.IsEOF() {
			// throw error
			return nil
		}

		l.Inc()
	}
}