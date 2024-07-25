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

func (l *Lexer) readUserLiteral(determinator func(byte) bool) string {
    pos := l.position
    for determinator(l.ch) {
        l.readChar()
    }

    return l.input[pos:l.position]
}

func (l *Lexer) peekChar() byte {
    if (len(l.input) > l.readPosition) {
        return l.input[l.readPosition]
    }

    return 0
}

// @util
func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// @util
func isDigit(ch byte) bool {
    return ch >= '0' && ch <= '9'
}

func (l *Lexer) NextToken() token.Token {
    var tok token.Token

    l.skipWhitespace()
    fmt.Printf(string(l.ch))
    fmt.Println()

    switch (l.ch) {
    case '=':
        tok = newToken(token.ASSIGN, l.ch)
        if l.peekChar() == '=' {
            tok.Type = token.EQ
            l.readChar()
            tok.Literal += string(l.ch)
        }
    case ';':
        tok = newToken(token.SEMICOLON, l.ch)
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
    case '-':
        tok = newToken(token.MINUS, l.ch)
    case '/':
        tok = newToken(token.DIVIDE, l.ch)
    case '!':
        tok = newToken(token.BANG, l.ch)
        if l.peekChar() == '=' {
            tok.Type = token.NOT_EQ
            l.readChar()
            tok.Literal += string(l.ch)
        }
    case '*':
        tok = newToken(token.STAR, l.ch)
    case '<':
        tok = newToken(token.LT, l.ch)
        if l.peekChar() == '=' {
            tok.Type = token.LTE
            l.readChar()
            tok.Literal += string(l.ch)
        }
    case '>':
        tok = newToken(token.GT, l.ch)
        if l.peekChar() == '=' {
            tok.Type = token.GTE
            l.readChar()
            tok.Literal += string(l.ch)
        }
    case 0:
        tok = newToken(token.EOF, 0)
        tok.Literal = ""
    default:
        // if acceptable character, create custom identifier
        if isLetter(l.ch) {
            tok.Type = token.IDENT
            tok.Literal = l.readUserLiteral(isLetter)
            if val, ok := token.CheckKeyword(tok.Literal); ok {
                tok.Type = val
            }
            // early exit is required, cus we l.readIdent will read chars until next non "letter"
            return tok
        } else if isDigit(l.ch) {
            tok.Type = token.INT
            tok.Literal = l.readUserLiteral(isDigit)

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

func (l *Lexer) skipWhitespace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        l.readChar()
    }
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
    l.readChar()
	return l
}

