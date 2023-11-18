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
	case Keyword_Return: {
		p.expect(Keyword_Return)
		value := p.parseExpr()
		p.expect(SEMI_COLON)
		return &ReturnStmt{Value: value}
	}
	default: v := &ExprStmt{
		Expr: p.parseExpr(),
	}
	p.expect(SEMI_COLON)
	return v
	}
}

func (p *Parser) parseBlockStmt() []Stmt {
	p.expect(OPEN_BRACE)
	block := []Stmt{}
	for p.notEof() && p.at().Type != CLOSE_BRACE {
		block = append(block, p.parseStmt())
	}
	p.expect(CLOSE_BRACE)
	return block
}

func (p *Parser) parseFunctionDeclaration() Stmt {
	var Type = &ObjType{
		NodeType: Type_Void,
	}
	p.expect(Keyword_Function)
	name := p.next().Literal
	params := p.parseParams(OPEN_PAREN, CLOSE_PAREN)
	if p.at().Type == ARROW {
		p.expect(ARROW)
		t := p.parsePrimaryExpr()
		if t.Type() != Node_ObjType {
			log.Fatalf("Expected A Type Not %v", t)
		}
		Type = t.(*ObjType)
	}
	body := p.parseBlockStmt()
	return &FunctionDeclaration{
		Name: name,
		Params: params,
		Body: body,
		ReturnType: Type,
	}
}

func (p *Parser) parseParams(open, close TokenType) []*TypedParameter {
	params := []*TypedParameter{}
	p.expect(open)
	if p.at().Type == close {
		p.expect(close)
		return params
	}
	params = append(params, &TypedParameter{
		Type: p.parsePrimaryExpr().(*ObjType),
		Param: p.parsePrimaryExpr().(*IdentiferLiteral).Value,
	})
	for p.at().Type == COMMA {
		p.expect(COMMA)
		params = append(params, &TypedParameter{
			Type: p.parsePrimaryExpr().(*ObjType),
			Param: p.parsePrimaryExpr().(*IdentiferLiteral).Value,
		})
	}
	p.expect(close)
	return params
}