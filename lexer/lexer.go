package lexer

import (
	"interpreter/token"
	"strings"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.ignorewhitespace()
	switch l.ch {
	case '=':
		tok = newtoken(token.ASSIGN, l.ch)
	case ';':
		tok = newtoken(token.SEMICOLON, l.ch)
	case '(':
		tok = newtoken(token.LPAREN, l.ch)
	case ')':
		tok = newtoken(token.RPAREN, l.ch)
	case ',':
		tok = newtoken(token.COMMA, l.ch)
	case '+':
		tok = newtoken(token.PLUS, l.ch)
	case '{':
		tok = newtoken(token.LBRACE, l.ch)
	case '}':
		tok = newtoken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newtoken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readNumber() string {
	pos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) ignorewhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '_'
}

func (l *Lexer) readIdentifier() string {
	pos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (lex *Lexer) readIdent() string {
	ret := strings.Builder{}
	for {
		lex.readChar()
		if isLetter(lex.ch) {
			ret.WriteByte(lex.ch)
		} else {
			break
		}
	}
	return ret.String()
}

func newtoken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{tokenType, string(char)}
}
