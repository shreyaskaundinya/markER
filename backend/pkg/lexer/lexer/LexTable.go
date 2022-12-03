package lexer

import (
	"strings"

	lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"
)

/*
	State for the Create Token
*/ 
func LexTable(l *Lexer) LexFn {
	l.SkipWhitespace()

	// if (l.IsEOF()) {}

	if (strings.HasPrefix(l.InputToEnd(), lexertoken.TBL_WEAK)) {
		// table type : WEAK
		l.Pos += len(lexertoken.TBL_WEAK)
		// fmt.Println("TABLE TYPE : ", l.Input[l.Start:l.Pos])
		l.Emit(lexertoken.TOKEN_TBL_WEAK, true)
		
	} else if (strings.HasPrefix(l.InputToEnd(), lexertoken.TBL_STRONG)) {
		// table type : STRONG
		l.Pos += len(lexertoken.TBL_STRONG)
		// fmt.Println("TABLE TYPE : ", l.Input[l.Start:l.Pos])
		l.Emit(lexertoken.TOKEN_TBL_STRONG, true)
	} else {
		// default table type : STRONG
		l.Emit(lexertoken.TOKEN_TBL_STRONG, false)
	}

	l.SkipWhitespace()

	if (strings.HasPrefix(l.InputToEnd(), lexertoken.TABLE)) {
		l.Pos += len(lexertoken.TABLE)
		l.Emit(lexertoken.TOKEN_TABLE, true)
	}

	return LexTableName
}
