package src

import (
	"log"
	"strconv"
)

func (p *Parser) parseExpr() Expr {
	return p.parseAssignmentExpr()
}

func (p *Parser) parseAssignmentExpr() Expr {
	left := p.parseBitwiseOr()
	if p.at().Type == BASIC_ASSIGN || p.at().Type == DIVIDE_ASSIGN ||
		p.at().Type == MINUS_ASSIGN || p.at().Type == MULTIPLY_ASSIGN || p.at().Type == PLUS_ASSIGN {
		return &AssignmentExpr{
			Op:    p.next(),
			Left:  left,
			Right: p.parseAssignmentExpr(),
		}
	}
	return left
}

func (p *Parser) parseLogicalOr() Expr {
	left := p.parseLogicalAnd()
	for p.at().Type == LOGIC_OR {
		left = &BinaryExpr{
			Op:    p.next(),
			Left:  left,
			Right: p.parseLogicalAnd(),
		}
	}
	return left
}

func (p *Parser) parseLogicalAnd() Expr {
	left := p.parseBitwiseOr()
	for p.at().Type == LOGIC_AND {
		left = &BinaryExpr{
			Op:    p.next(),
			Left:  left,
			Right: p.parseBitwiseOr(),
		}
	}
	return left
}

func (p *Parser) parseBitwiseOr() Expr {
	left := p.parseBitwiseAnd()
	for p.at().Type == PIPE {
		left = &BinaryExpr{
			Op:    p.next(),
			Left:  left,
			Right: p.parseBitwiseAnd(),
		}
	}
	return left
}

func (p *Parser) parseBitwiseAnd() Expr {
	left := p.parseBitwiseXOR()
	for p.at().Type == AMPERSAN {
		left = &BinaryExpr{
			Op:    p.next(),
			Left:  left,
			Right: p.parseBitwiseXOR(),
		}
	}
	return left
}

func (p *Parser) parseBitwiseXOR() Expr {
	left := p.parseAdditiveExpr()
	for p.at().Type == XOR {
		left = &BinaryExpr{
			Op:    p.next(),
			Left:  left,
			Right: p.parseAdditiveExpr(),
		}
	}
	return left
}

func (p *Parser) parseAdditiveExpr() Expr {
	left := p.parseMultiplicativeExpr()
	for p.at().Type == PLUS || p.at().Type == MINUS {
		left = &BinaryExpr{
			Op:    p.next(),
			Left:  left,
			Right: p.parseAdditiveExpr(),
		}
	}
	return left
}

func (p *Parser) parseMultiplicativeExpr() Expr {
	left := p.parsePrimaryExpr()
	for p.at().Type == MULTIPLY || p.at().Type == DIVIDE || p.at().Type == MODULO {
		left = &BinaryExpr{
			Op:    p.next(),
			Left:  left,
			Right: p.parsePrimaryExpr(),
		}
	}
	return left
}

func (p *Parser) parsePrimaryExpr() Expr {
	switch p.at().Type {
	case Identifer:
		return &IdentiferLiteral{Value: p.next().Literal}
	case Literal_Int:
		value, _ := strconv.ParseInt(p.next().Literal, 10, 64)
		return &IntLiteral{Value: value}
	case Literal_Float:
		value, _ := strconv.ParseFloat(p.next().Literal, 64)
		return &FloatLiteral{Value: value}
	case Literal_String:
		return &StringLiteral{Value: p.next().Literal}
	case Literal_Boolean:
		value, _ := strconv.ParseBool(p.next().Literal)
		return &BooleanLiteral{Value: value}
	case OPEN_PAREN:
		p.expect(OPEN_PAREN)
		ex := p.parseExpr()
		p.expect(CLOSE_PAREN)
		return ex
	default: log.Fatalf("Unexpected Token %v", p.at()); return nil
	}
}