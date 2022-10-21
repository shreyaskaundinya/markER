package lexer

import (
	"fmt"
	"strings"

	lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"
)


func LexBegin(l *Lexer) LexFn {
	fmt.Println("Starting LEXING...")
	// skip the whitespace
	l.SkipWhitespace()

	// the start token can either be a create token or a relation token
	if strings.HasPrefix(l.InputToEnd(), lexertoken.CREATE) {
		fmt.Println("STATE : CREATE")
		return LexCreate
		} else if strings.HasPrefix(l.InputToEnd(), lexertoken.RELN) {
		fmt.Println("STATE : RELN")
		return LexReln
	} else {
	return nil
	}
}