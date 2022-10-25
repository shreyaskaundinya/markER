package lexer

import (
	"unicode"

	lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"
)

func LexTable2(l *Lexer) LexFn {
	l.SkipWhitespace()

	var ch rune

	for {
		ch = rune(l.Input[l.Pos])
		if unicode.IsSpace(ch) {
			l.Emit(lexertoken.TOKEN_TABLE_NAME, true)
			return LexSemicolon
		}

		if l.IsEOF() {
			return nil
		}
		l.Inc()
	}
}