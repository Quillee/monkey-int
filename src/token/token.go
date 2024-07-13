package token

type TokenType string

type Token struct {
    Type TokenType
    Literal string
}

var keywords = map[string]TokenType {
    "fn": FUNCTION,
    "let": LET,
}

const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"

    // Identifiers and literals
    IDENT = "IDENT"
    INT = "INT"

    // Operators
    ASSIGN = "="
    PLUS = "+"

    // Delimiters
    COMMA = ","
    SEMICOLON = ";"
    LPAREN = "("
    RPAREN = ")"
    LSQUIG = "{"
    RSQUIG = "}"

    // Keywords
    FUNCTION = "FUNCTION"
    LET = "LET"
)

