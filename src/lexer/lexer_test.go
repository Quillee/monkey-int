package lexer

import (
	"testing"

	"github.com/quillee/monkey/src/token"
)

type ExpectationMap struct {
    expectedType    token.TokenType
    expectedLiteral string
}

type FormatterFunc func (format string, args ...any) 

func test_util_map_token_to_literal(tt token.TokenType) ExpectationMap {
    switch (tt) {
    case token.ASSIGN:
        return ExpectationMap{ token.ASSIGN, "=" }
    case token.PLUS:
        return ExpectationMap{token.PLUS, "+"}
    case token.LPAREN:
        return ExpectationMap{token.LPAREN, "("}
    case token.RPAREN:
        return ExpectationMap{token.RPAREN, ")"}
    case token.LSQUIG:
        return ExpectationMap{token.LSQUIG, "{"}
    case token.RSQUIG:
        return ExpectationMap{token.RSQUIG, "}"}
    case token.COMMA:
        return ExpectationMap{token.COMMA, ","}
    case token.SEMICOLON:
        return ExpectationMap{token.SEMICOLON, ";"}
    case token.LET:
        return ExpectationMap{ token.LET, "let" }
    case token.EOF:
        return ExpectationMap{token.EOF, ""}
    }
    return ExpectationMap{}
}

func test_util_check_expected_results(lexer *Lexer, tests []ExpectationMap, error_pipe FormatterFunc) {
	for i, tt := range tests {
		tok := lexer.NextToken()

		if tok.Type != tt.expectedType {
			error_pipe("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			error_pipe("tests[%d] - literal wrong. expected =%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken(t *testing.T) {
	input := `=+(){},;`


	tests := []ExpectationMap {
        test_util_map_token_to_literal(token.ASSIGN),
        test_util_map_token_to_literal(token.PLUS),
        test_util_map_token_to_literal(token.LPAREN),
        test_util_map_token_to_literal(token.RPAREN),
        test_util_map_token_to_literal(token.LSQUIG),
        test_util_map_token_to_literal(token.RSQUIG),
        test_util_map_token_to_literal(token.COMMA),
        test_util_map_token_to_literal(token.SEMICOLON),
        test_util_map_token_to_literal(token.EOF),
	}

	lexer := New(input)

    test_util_check_expected_results(lexer, tests, t.Fatalf)

    input = `let five = 5;
    let ten = 10;

    let add = fn(x, y) {
        x + y;
    };

    let result = add(five, ten);`

	lexer = New(input)

    tests = []ExpectationMap{
        test_util_map_token_to_literal(token.LET),
        {token.IDENT, "five"},
        test_util_map_token_to_literal(token.ASSIGN),
        { token.INT, "5" },
        test_util_map_token_to_literal(token.SEMICOLON),
        test_util_map_token_to_literal(token.LET),
        { token.IDENT, "ten" },
        test_util_map_token_to_literal(token.ASSIGN),
        { token.INT, "10" },
        test_util_map_token_to_literal(token.SEMICOLON),
        test_util_map_token_to_literal(token.LET),
        {token.IDENT, "add"},
        test_util_map_token_to_literal(token.ASSIGN),
        {token.FUNCTION, "fn"},
        test_util_map_token_to_literal(token.LPAREN),
        {token.IDENT, "x"},
        test_util_map_token_to_literal(token.COMMA),
        {token.IDENT, "y"},
        test_util_map_token_to_literal(token.RPAREN),
        test_util_map_token_to_literal(token.LSQUIG),
        {token.IDENT, "x"},
        test_util_map_token_to_literal(token.PLUS),
        {token.IDENT, "y"},
        test_util_map_token_to_literal(token.SEMICOLON),

        test_util_map_token_to_literal(token.RSQUIG),
        test_util_map_token_to_literal(token.SEMICOLON),

        test_util_map_token_to_literal(token.LET),
        {token.IDENT, "result"},
        test_util_map_token_to_literal(token.ASSIGN),
        {token.IDENT, "add"},
        test_util_map_token_to_literal(token.LPAREN),
        {token.IDENT, "five"},
        test_util_map_token_to_literal(token.COMMA),
        {token.IDENT, "ten"},
        test_util_map_token_to_literal(token.RPAREN),
        test_util_map_token_to_literal(token.SEMICOLON),
    }

    test_util_check_expected_results(lexer, tests, t.Fatalf)
}
