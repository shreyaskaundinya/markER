package lexer

import (
	lexertoken "go-marker-backend/pkg/lexer/lexertoken"
	"unicode"
	"unicode/utf8"
)

type LexFn func(*Lexer) LexFn

/*
	Add the token to the channel
*/
func (l *Lexer) Emit(tokenType lexertoken.TokenType) {
	// add token to channel
	l.Tokens <- lexertoken.Token{Type: tokenType, Value: l.Input[l.Start:l.Pos]}
	// move the pointer to after the token
	l.Start = l.Pos
}

func (l *Lexer) Inc() {
	l.Pos++
	if l.Pos > utf8.RuneCountInString(l.Input) {
		l.Emit(lexertoken.TOKEN_EOF)
	}
}

func (l *Lexer) Dec() {
	l.Pos--
}

func (l *Lexer) Next() rune {
	l.Inc()
	r := rune(l.Input[l.Pos])
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
			l.Emit(lexertoken.TOKEN_EOF)
		}
	}
}

