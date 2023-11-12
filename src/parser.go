package src

import (
	"log"
	. "omega/src/ast"
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
	prog := &Program{File:p.File, Body: []Stmt{}}
	p.expect(Keyword_Package)
	prog.PackageName = p.expect(Identifer).Literal
	p.expect(SEMI_COLON)
	for p.notEof() {
		prog.Body = append(prog.Body, p.parseBaseStmt())
	}
	return prog
}

func (p *Parser) parseBaseStmt() Stmt {
	switch p.at().Type {
	}
	return nil
}