package lexer

import (
	"unicode"
	"unicode/utf8"

	lexertoken "github.com/shreyaskaundinya/markER/backend/pkg/lexer/token"
)

type LexFn func(*Lexer) LexFn

/*
	Add the token to the channel
*/
func (l *Lexer) Emit(tokenType lexertoken.TokenType, move bool) {
	// add token to channel
	l.Tokens <- lexertoken.Token{Type: tokenType, Value: l.Input[l.Start:l.Pos]}
	
	if move {
		// move the pointer to after the token
		l.Start = l.Pos
	}
}

func (l *Lexer) Inc() {
	l.Pos++
	if l.Pos > utf8.RuneCountInString(l.Input) {
		l.Emit(lexertoken.TOKEN_EOF, true)
	}
}

func (l *Lexer) Dec() {
	l.Pos--
}

func (l *Lexer) Next() rune {
	r := rune(l.Input[l.Pos])
	l.Inc()
	return r
}

func (l *Lexer) InputToEnd() string {
	return l.Input[l.Pos:]
}

func (l *Lexer) SkipWhitespace() {
	for {
		ch := l.Next()
		if !unicode.IsSpace(ch) {
			l.Dec()
			break
		}
		if ch == lexertoken.EOF {
			l.Emit(lexertoken.TOKEN_EOF, true)
		}
	}
}

func (l *Lexer) IsEOF() bool {
	if ch := l.Next(); ch == lexertoken.EOF {
		l.Emit(lexertoken.TOKEN_EOF, true)
		return true
	} else {
		return false;
	}
}