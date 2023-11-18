package src

import (
	"log"
	"strings"
)

type Parser struct {
	Tokens []Token
	File   string
	ptr    int
}

func CreateParser(lex *Lexer) *Parser {
	return &Parser{Tokens: lex.Tokenize(), File: lex.File, ptr: 0}
}

func (p *Parser) notEof() bool {
	return p.at().Type != EOF
}

func (p *Parser) at() Token { return p.Tokens[p.ptr] }

func (p *Parser) next() Token { p.ptr++; return p.Tokens[p.ptr-1] }

func (p *Parser) expect(Expected TokenType) Token {
	p.ptr++; l := p.Tokens[p.ptr-1]
	if l.Type != Expected {
		log.Fatalf("Expected %v Got %v", Expected, l)
	}
	return l
}

func (p *Parser) ProduceAst() Stmt {
	prog := &Program{File:p.File, Body: []Stmt{}, PackageName: strings.Split(p.File, ".")[0]}
	for p.notEof() {
		prog.Body = append(prog.Body, p.parseStmt())
	}
	return prog
}

func (p *Parser) parseStmt() Stmt {
	switch p.at().Type {
	case Keyword_Function: return p.parseFunctionDeclaration()
	case Keyword_Return: return &ReturnStmt{Value: p.parseExpr()}
	default: return p.parseExpr()
	}
}

func (p *Parser) parseBlockStmt() []Stmt {
	block := []Stmt{}
	for p.notEof() {
		block = append(block, p.parseStmt())
	}
	return block
}

func (p *Parser) parseFunctionDeclaration() Stmt {
	p.expect(Keyword_Function)
	name := p.next().Literal
	body := p.parseBlockStmt()
	return &FunctionDeclaration{
		Name: name,
		Body: body,
	}
}