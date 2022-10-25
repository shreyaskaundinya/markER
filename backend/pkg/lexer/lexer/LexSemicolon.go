package lexer

import (
	"fmt"
	"strings"

	lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"
)

func LexTable1(l *Lexer) LexFn {
	l.SkipWhitespace()

	if(strings.HasPrefix(l.InputToEnd(), lexertoken.SEMICOLON)) {
		l.Pos += len(lexertoken.SEMICOLON)
		l.Emit(lexertoken.TOKEN_SEMICOLON, true)
		fmt.Println("TAG: DELIMITED")
	}

	return LexBegin
}