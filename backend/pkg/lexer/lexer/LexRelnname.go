package lexer

import (
	"fmt"
	"strings"

	lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"
)

func LexRelnname(l *Lexer) LexFn {
	l.SkipWhitespace()
	if (strings.HasPrefix(l.InputToEnd(), "'")){
		l.Pos += len("'")
		
		// skip the quote
		l.Start += len("'")

		var ch rune

		for {
			ch = rune(l.Input[l.Pos])
			// found a single quote
			if (ch == 39) {
				l.Emit(lexertoken.TOKEN_RELN_NAME, true)
				l.Inc()
				return LexCardinality2
			}

			if l.IsEOF() {
				return nil
			}
			l.Inc()
		}
		
	} else{
		fmt.Println("no single quotes")
		return nil
	}
	
}