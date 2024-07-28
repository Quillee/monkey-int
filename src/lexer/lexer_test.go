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
        return ExpectationMap{ tt, "=" }
    case token.PLUS:
        return ExpectationMap{ tt, "+"}
    case token.LPAREN:
        return ExpectationMap{ tt, "("}
    case token.RPAREN:
        return ExpectationMap{ tt, ")"}
    case token.LSQUIG:
        return ExpectationMap{ tt, "{"}
    case token.RSQUIG:
        return ExpectationMap{ tt, "}"}
    case token.COMMA:
        return ExpectationMap{ tt, ","}
    case token.SEMICOLON:
        return ExpectationMap{ tt, ";"}
    case token.LET:
        return ExpectationMap{ tt, "let" }
    case token.STAR:
        return ExpectationMap{ tt, "*" } 
    case token.DIVIDE:
        return ExpectationMap{ tt, "/" }
    case token.MINUS:
        return ExpectationMap{ tt, "-" }
    case token.BANG:
        return ExpectationMap{ tt, "!" } 
    case token.EQ:
        return ExpectationMap{ tt, "=="}
    case token.NOT_EQ:
        return ExpectationMap{ tt, "!="}
    case token.LT:
        return ExpectationMap{ tt, "<"}
    case token.LTE:
        return ExpectationMap{ tt, "<="}
    case token.GT:
        return ExpectationMap{ tt, ">"}
    case token.GTE:
        return ExpectationMap{ tt, ">="}
    case token.ELSE:
        return ExpectationMap{ tt, "else" } 
    case token.IF:
        return ExpectationMap{ tt, "if" }
    case token.RETURN:
        return ExpectationMap{ tt, "return" }
    case token.TRUE:
        return ExpectationMap{ tt, "true" }
    case token.FALSE:
        return ExpectationMap{ tt, "false" }
    case token.EOF:
        return ExpectationMap{ tt, ""}
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
        let add = fn(x, y) { x + y; };
        let result = add(five, ten);

        !-/*5;
        5 < 10 > 5;

        if (5 < 10) { 
            return true;
        } else { 
            return false;
        }

        10 == 10;
        10 != 9;
    `


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
        test_util_map_token_to_literal(token.BANG),
        test_util_map_token_to_literal(token.MINUS),
        test_util_map_token_to_literal(token.DIVIDE),
        test_util_map_token_to_literal(token.STAR),
        {token.INT, "5"},
        test_util_map_token_to_literal(token.SEMICOLON),

        {token.INT, "5"},
        test_util_map_token_to_literal(token.LT),
        {token.INT, "10"},
        test_util_map_token_to_literal(token.GT),
        {token.INT, "5"},
        test_util_map_token_to_literal(token.SEMICOLON),

        test_util_map_token_to_literal(token.IF),
        test_util_map_token_to_literal(token.LPAREN),
        {token.INT, "5"},
        test_util_map_token_to_literal(token.LT),
        {token.INT, "10"},
        test_util_map_token_to_literal(token.RPAREN),
        test_util_map_token_to_literal(token.LSQUIG),
        test_util_map_token_to_literal(token.RETURN),
        test_util_map_token_to_literal(token.TRUE),
        test_util_map_token_to_literal(token.SEMICOLON),

        test_util_map_token_to_literal(token.RSQUIG),
        test_util_map_token_to_literal(token.ELSE),
        test_util_map_token_to_literal(token.LSQUIG),
        test_util_map_token_to_literal(token.RETURN),
        test_util_map_token_to_literal(token.FALSE),
        test_util_map_token_to_literal(token.SEMICOLON),
        test_util_map_token_to_literal(token.RSQUIG),
        {token.INT, "10"},
        test_util_map_token_to_literal(token.EQ),
        {token.INT, "10"},
        test_util_map_token_to_literal(token.SEMICOLON),

        {token.INT, "10"},
        test_util_map_token_to_literal(token.NOT_EQ),
        {token.INT, "9"},
        test_util_map_token_to_literal(token.SEMICOLON),
    }

    test_util_check_expected_results(lexer, tests, t.Fatalf)
}
