package lexer

import (
	"fmt"

	"github.com/quillee/monkey/src/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func (l *Lexer) readChar() {
	if l.readPosition < len(l.input) {
		l.ch = l.input[l.readPosition]
	} else {
        // 0 is NUL or EOF in ASCII
		l.ch = 0
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) readIdent() string {
    pos := l.position
    for isLetter(l.ch) {
        l.readChar()
    }

    return l.input[pos:l.position]
}

// @util
func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) NextToken() token.Token {
    var tok token.Token
    fmt.Println(string(l.ch))
    switch (l.ch) {
    case '=':
        tok = newToken(token.ASSIGN, l.ch)
    case '{':
        tok = newToken(token.LSQUIG, l.ch)
    case '}':
        tok = newToken(token.RSQUIG, l.ch)
    case '(':
        tok = newToken(token.LPAREN, l.ch)
    case ')':
        tok = newToken(token.RPAREN, l.ch)
    case ',':
        tok = newToken(token.COMMA, l.ch)
    case '+':
        tok = newToken(token.PLUS, l.ch)
    case ';':
        tok = newToken(token.SEMICOLON, l.ch)
    case 0:
        tok = newToken(token.EOF, 0)
        tok.Literal = ""
    default:
        if isLetter(l.ch) {
            tok.Type = token.IDENT
            tok.Literal = l.readIdent()
            return tok
        } else {
            tok = newToken(token.ILLEGAL, l.ch)
        }
    }

    l.readChar()
    return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token { Type: tokenType, Literal: string(ch) }
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
    l.readChar()
	return l
}

