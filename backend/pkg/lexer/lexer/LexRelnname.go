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
		var ch rune

		for {
			ch = rune(l.Input[l.Pos])
			if (ch == 39) {
				l.Emit(lexertoken.TOKEN_RELN_NAME, true)
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