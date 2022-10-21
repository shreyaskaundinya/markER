package lexer

import (
	"fmt"
	"unicode"

	lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"
)

func LexTableName(l *Lexer) LexFn {
	l.SkipWhitespace()

	var ch rune

	for {
		ch = rune(l.Input[l.Pos])
		if unicode.IsSpace(ch) {
			l.Emit(lexertoken.TOKEN_TABLE_NAME, true)
			fmt.Println("SPACE FOUND")
			return nil
		}

		if l.IsEOF() {
			return nil
		}
		l.Inc()
	}
}