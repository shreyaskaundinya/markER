package lexer

import (
	lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"
)

func LexTable2(l *Lexer) LexFn {
	l.SkipWhitespace()

	var ch rune

	for {
		ch = rune(l.Input[l.Pos])
		// read till semi colon
		if ch == 59 {
			l.Emit(lexertoken.TOKEN_TABLE_NAME, true)
			return LexSemicolon
		}

		if l.IsEOF() {
			return nil
		}
		l.Inc()
	}
}