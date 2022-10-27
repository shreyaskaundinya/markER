package lexer

import (
	"strings"

	lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"
)

func LexAttributeType(l *Lexer) LexFn {
	l.SkipWhitespace()

	inp := l.InputToEnd()

	var token lexertoken.TokenType

	switch {
		case strings.HasPrefix(inp, lexertoken.DTYPE_BOOL):
			l.Pos += len(lexertoken.DTYPE_BOOL)
			token = lexertoken.TOKEN_DTYPE_BOOL
			
		case strings.HasPrefix(inp, lexertoken.DTYPE_DATE):
			l.Pos += len(lexertoken.DTYPE_DATE)
			token = lexertoken.TOKEN_DTYPE_DATE
			
		case strings.HasPrefix(inp, lexertoken.DTYPE_DATETIME):
			l.Pos += len(lexertoken.DTYPE_DATETIME)
			token = lexertoken.TOKEN_DTYPE_DATETIME
			
		case strings.HasPrefix(inp, lexertoken.DTYPE_FLOAT):
			l.Pos += len(lexertoken.DTYPE_FLOAT)
			token = lexertoken.TOKEN_DTYPE_FLOAT
			
		case strings.HasPrefix(inp, lexertoken.DTYPE_INT):
			l.Pos += len(lexertoken.DTYPE_INT)
			token = lexertoken.TOKEN_DTYPE_INT
			
		case strings.HasPrefix(inp, lexertoken.DTYPE_TIMESTAMP):
			l.Pos += len(lexertoken.DTYPE_TIMESTAMP)
			token = lexertoken.TOKEN_DTYPE_TIMESTAMP
			
		case strings.HasPrefix(inp, lexertoken.DTYPE_VARCHAR):
			l.Pos += len(lexertoken.DTYPE_VARCHAR)
			token = lexertoken.TOKEN_DTYPE_VARCHAR
		
		case strings.HasPrefix(inp, lexertoken.DTYPE_COMPOSITE):
			l.Pos += len(lexertoken.DTYPE_COMPOSITE)
			token = lexertoken.TOKEN_DTYPE_COMPOSITE
			l.Emit(token, true)
			// return composite left bracket
			return LexCompBracketOpen
		
		default:
			token = -1
	}

	if token > 0 {
		l.Emit(token, true)
		// TODO : LexAttributeProperties
		//return LexAttributeComma
		return LexAttributeProps
	} else {
		// throw error
		return nil
	}
}