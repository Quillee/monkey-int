package token

type TokenType string

type Token struct {
    Type TokenType
    Literal string
}

var keywords = map[string]TokenType {
    "fn": FUNCTION,
    "let": LET,
    "true": TRUE,
    "false": FALSE,
    "if": IF,
    "else": ELSE,
    "return": RETURN,
}

func CheckKeyword(key string) (TokenType, bool) {
    if val, ok := keywords[key]; ok {
        return val, ok
    }
    
    return IDENT, false
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
    MINUS = "-"
    DIVIDE = "/"
    STAR = "*"
    BANG  = "!"
    LT = "<"
    LTE = "<="
    GT = ">"
    GTE = ">="
    EQ = "=="
    NOT_EQ = "!="


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
    RETURN = "RETURN"
    IF = "IF"
    ELSE = "ELSE"
    TRUE = "TRUE"
    FALSE = "FALSE"
)

