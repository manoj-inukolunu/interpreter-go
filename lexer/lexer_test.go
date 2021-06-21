package lexer

import (
	"interpreter/token"
	"testing"
)

func TestNextToken1(t *testing.T) {
	input := `var five = 5;
		var ten = 10;
		var add = fn(x,y){
			x+y;
		};
		var result = add(five,ten);
		`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.VAR, "var"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lex := New(input)

	for idx, val := range tests {
		tok := lex.NextToken()
		if tok.Type != val.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q,got%q", idx, val.expectedType, tok.Type)
		}
		if tok.Literal != val.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong.expected=%q,got=%q", idx, val.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	lex := New(input)

	for idx, val := range tests {
		tok := lex.NextToken()
		if tok.Type != val.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q,got%q", idx, val.expectedType, tok.Type)
		}
		if tok.Literal != val.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong.expected=%q,got=%q", idx, val.expectedLiteral, tok.Literal)
		}
	}
}
