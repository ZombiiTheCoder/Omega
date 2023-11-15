package ast

type NodeType int

const (
	Node_ExprStmt NodeType = iota
	Node_Program
	Node_ExitStmt
)

type Node interface {
	Type() NodeType
}
type Stmt interface{ Node }
type Expr interface{ Node }

type ExprStmt struct {
	Stmt
	Expr Expr
}

func (s *ExprStmt) Type() NodeType { return Node_ExprStmt }

type Program struct {
	Stmt
	PackageName string
	Body        []Stmt
	File        string
}

func (s *Program) Type() NodeType { return Node_Program }

type ExitStmt struct {
	Stmt
	StatusCode int
}

func (s *ExitStmt) Type() NodeType { return Node_ExitStmt }
