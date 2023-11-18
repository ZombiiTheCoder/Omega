package src

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

type Lexer struct {
	Content []byte
	File    string
	ptr 	int
}

func CreateLexer(File string) *Lexer {
	data, err := os.ReadFile(File)
	if err != nil {
		log.Fatalf("Could Not Read File %v, May Not Exist\n", File)
	}
	return &Lexer{Content: data, File: File, ptr: 0}
}

func isNum(char byte) bool {
	return char >= '0' && char <= '9'
}

func isAlpha(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_' || char == '$'
}

func isAlnum(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_' || char == '$' || char >= '0' && char <= '9'
}

func (l *Lexer) notEof() bool {
	return l.ptr <= (len(l.Content)-1) 
}

func (l *Lexer) at() byte {
	if l.notEof() {
		return l.Content[l.ptr]
	}; return ' '
}

func (l *Lexer) peek(p int) byte {
	if l.ptr + p < (len(l.Content)-1) {
		return l.Content[l.ptr+p]
	}; return ' '
}

func (l *Lexer) Tokenize() []Token {
	Tokens := []Token{}
	for l.notEof() {
		switch l.at() {
		case '\n', '\r', ' ', '\t', '\b': l.ptr++
		
			case '-':
				switch l.peek(1) {
					case '=': Tokens = append(Tokens, Token{Literal: "-=", Type: MINUS_ASSIGN}); l.ptr+=2
					case '>': Tokens = append(Tokens, Token{Literal: "->", Type: ARROW}); l.ptr+=2
					default: Tokens = append(Tokens, Token{Literal: "-", Type: MINUS}); l.ptr++
				}
			case '+':
				switch l.peek(1) {
					case '=': Tokens = append(Tokens, Token{Literal: "+=", Type: PLUS_ASSIGN}); l.ptr+=2
					default: Tokens = append(Tokens, Token{Literal: "+", Type: PLUS}); l.ptr++
				}
			case '*':
				switch l.peek(1) {
					case '=': Tokens = append(Tokens, Token{Literal: "*=", Type: MULTIPLY_ASSIGN}); l.ptr+=2
					default: Tokens = append(Tokens, Token{Literal: "*", Type: MULTIPLY}); l.ptr++
				}
			case '/':
				switch l.peek(1) {
					case '=': Tokens = append(Tokens, Token{Literal: "/=", Type: DIVIDE_ASSIGN}); l.ptr+=2
					case '/': {
						l.ptr+=2
						for l.at() != '\n' && l.notEof() { l.ptr++ }
					}
					default: Tokens = append(Tokens, Token{Literal: "/", Type: DIVIDE}); l.ptr++
				}
			case '!':
				switch l.peek(1) {
					case '=': Tokens = append(Tokens, Token{Literal: "!=", Type: NOT_EQUALS}); l.ptr+=2
					default: Tokens = append(Tokens, Token{Literal: "!", Type: NOT}); l.ptr++
				}
			case '=':
				switch l.peek(1) {
					case '=': Tokens = append(Tokens, Token{Literal: "==", Type: EQUALS}); l.ptr+=2
					default: Tokens = append(Tokens, Token{Literal: "=", Type: BASIC_ASSIGN}); l.ptr++
				}
				
			case '.': Tokens = append(Tokens, Token{Literal: ".", Type: DOT}); l.ptr++
			case '?': Tokens = append(Tokens, Token{Literal: "?", Type: QUESTION}); l.ptr++
			case ':': Tokens = append(Tokens, Token{Literal: ":", Type: COLON}); l.ptr++
			case ';': Tokens = append(Tokens, Token{Literal: ";", Type: SEMI_COLON}); l.ptr++
			case ',': Tokens = append(Tokens, Token{Literal: ",", Type: COMMA}); l.ptr++

			case '{': Tokens = append(Tokens, Token{Literal: "{", Type: OPEN_BRACE}); l.ptr++
			case '}': Tokens = append(Tokens, Token{Literal: "}", Type: CLOSE_BRACE}); l.ptr++
			case '[': Tokens = append(Tokens, Token{Literal: "[", Type: OPEN_BRACKET}); l.ptr++
			case ']': Tokens = append(Tokens, Token{Literal: "]", Type: CLOSE_BRACKET}); l.ptr++
			case '(': Tokens = append(Tokens, Token{Literal: "(", Type: OPEN_PAREN}); l.ptr++
			case ')': Tokens = append(Tokens, Token{Literal: ")", Type: CLOSE_PAREN}); l.ptr++

		}

		if isAlpha(l.at()) {
			var value bytes.Buffer
			for isAlnum(l.at()) && l.notEof() { value.WriteByte(l.at()); l.ptr++ }
			Tokens = append(Tokens, Token{Literal: value.String(), Type: ToKeyword(value.String())})
		}

		if l.at() == '"' {
			l.ptr++
			var value bytes.Buffer
			for l.at() != '"' && l.notEof() { value.WriteByte(l.at()); l.ptr++ }; l.ptr++
			Tokens = append(Tokens, Token{Literal: value.String(), Type: Literal_String})
		}

		if isNum(l.at()) {
			Type := Literal_Int
			var value bytes.Buffer
			for isNum(l.at()) && l.notEof() { value.WriteByte(l.at()); l.ptr++ }
			if l.at() == '.' {
				Type = Literal_Float
				value.WriteByte(l.at()); l.ptr++
				for isNum(l.at()) && l.notEof() {
					value.WriteByte(l.at()); l.ptr++
				}
			}
			Tokens = append(Tokens, Token{Literal: value.String(), Type: Type})
		}

	}
	fmt.Println(string(l.at()))
	Tokens = append(Tokens, Token{Literal: "End Of File", Type: EOF})
	return Tokens
}