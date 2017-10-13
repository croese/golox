package lexer

import "github.com/croese/golox/token"

type Lexer struct {
	input         string
	position      int // current position; on current char
	readPosition  int // current reading position; after current char
	ch            byte
	currentLine   int
	currentColumn int
}

func New(input string) *Lexer {
	l := &Lexer{input: input, currentLine: 1, currentColumn: 0}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '(':
		tok = l.newToken(token.LEFT_PAREN)
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.isAtEnd() {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.advance()
}

func (l *Lexer) advance() {
	l.position = l.readPosition
	l.readPosition += 1
	l.currentColumn++
}

func (l *Lexer) peekChar() byte {
	if l.isAtEnd() {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) isAtEnd() bool {
	return l.readPosition >= len(l.input)
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.currentLine++
			l.currentColumn = 0
		}
		l.readChar()
	}
}

func (l *Lexer) newToken(t token.TokenType) token.Token {

}
