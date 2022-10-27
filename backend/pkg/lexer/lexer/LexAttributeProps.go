package lexer

import (
	"strings"
	"fmt"
	lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"
)

func LexAttributeProps(l *Lexer) LexFn {
	l.SkipWhitespace()

	inp := l.InputToEnd()

	if (strings.HasPrefix(inp, lexertoken.PROP_DERIVED)){
		l.Pos += len(lexertoken.PROP_DERIVED)
		l.Emit(lexertoken.TOKEN_PROP_DERIVED)
		fmt.Println("TAG: DERIVED KEY")

	if (strings.HasPrefix(inp, lexertoken.PROP_FK)){
		l.Pos += len(lexertoken.PROP_FK)
		l.Emit(lexertoken.TOKEN_PROP_FK)
		fmt.Println("TAG: FOREIGN KEY")

	}
	l.SkipWhitespace()	
	
	if (strings.HasPrefix(inp, lexertoken.PROP_NOTNULL)){
		l.Pos += len(lexertoken.PROP_NOTNULL)
		l.Emit(lexertoken.TOKEN_PROP_NOTNULL)
		fmt.Println("TAG: NOTNULL")
	}
	l.SkipWhitespace()

	if (strings.HasPrefix(inp, lexertoken.PROP_MULVAL)) {
		l.Pos += len(lexertoken.PROP_MULVAL)
		l.Emit(lexertoken.TOKEN_PROP_MULVAL)
		fmt.Println("TAG: MULTIVALUE")
	}
	l.SkipWhitespace()

	if (strings.HasPrefix(inp, lexertoken.PROP_PK)) {
		l.Pos += len(lexertoken.PROP_PK)
		l.Emit(lexertoken.TOKEN_PROP_PK)
		fmt.Println("TAG: PRIMARY KEY")
	}
	l.SkipWhitespace()
	
	if (strings.HasPrefix(inp, lexertoken.PROP_UNIQUE)){
		l.Pos += len(lexertoken.PROP_UNIQUE)
		l.Emit(lexertoken.TOKEN_PROP_UNIQUE)
		fmt.Println("TAG: UNIQUE")
	}

	return LexAttributeComma
}