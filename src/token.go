package src

import "fmt"

type Token struct {
	Type    TokenType
	Literal string
}

func (t *Token) String() string {
	return fmt.Sprintf("Literal: %v, Type: %v", t.Literal, t.Type)
}

type TokenType int

const (
	EOF TokenType = iota

	MINUS
	PLUS
	MULTIPLY
	DIVIDE

	BASIC_ASSIGN
	MINUS_ASSIGN
	PLUS_ASSIGN
	MULTIPLY_ASSIGN
	DIVIDE_ASSIGN

	OPEN_BRACE
	CLOSE_BRACE
	OPEN_BRACKET
	CLOSE_BRACKET
	OPEN_PAREN
	CLOSE_PAREN
	DOT
	QUESTION
	COLON
	SEMI_COLON
	COMMA
	NOT
	ARROW

	PIPE
	LOGIC_OR
	AMPERSAN
	LOGIC_AND
	XOR
	MODULO

	NOT_EQUALS
	EQUALS

	// Types
	Type_String
	Type_Int
	Type_Float
	Type_Boolean
	Type_Function

	// Literals
	Identifer
	Literal_String
	Literal_Int
	Literal_Float
	Literal_Boolean
	
	// Keyword
	Keyword_Function
	Keyword_If
	Keyword_Else
	Keyword_Return
)

func ToKeyword(value string) TokenType {
	keys := map[string]TokenType{
		"int":     Type_Int,
		"float":   Type_Float,
		"bool": Type_Boolean,
		"str":  Type_String,
		"Func": Type_Function,

		"fn": Keyword_Function,
		"if": Keyword_If,
		"else": Keyword_Else,
		"return": Keyword_Return,
	}
	if val, ok := keys[value]; ok {
		return val
	}
	return Identifer
}