package src

type NodeType int

const (
	Node_ExprStmt NodeType = iota
	Node_Program
	Node_FunctionDeclaration
	Node_VariableDeclaration
	Node_ReturnStmt

	Node_AssignmentExpr
	Node_BinaryExpr
	Node_UnaryExpr

	Node_Identifer
	Node_Int
	Node_Float
	Node_Bool
	Node_String
	Node_Null

	Node_ObjType
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

type TypedParameter struct {
	Type  *ObjType
	Param string
}

type FunctionDeclaration struct {
	Stmt
	Name       string
	Params     []*TypedParameter
	Body       []Stmt
	ReturnType *ObjType
}

func (s *FunctionDeclaration) Type() NodeType { return Node_FunctionDeclaration }

type ReturnStmt struct {
	Stmt
	Value Expr
}

func (s *ReturnStmt) Type() NodeType { return Node_ReturnStmt }

type AssignmentExpr struct {
	Expr
	Op    Token
	Left  Expr
	Right Expr
}

func (s *AssignmentExpr) Type() NodeType { return Node_AssignmentExpr }

type BinaryExpr struct {
	Expr
	Op    Token
	Left  Expr
	Right Expr
}

func (s *BinaryExpr) Type() NodeType { return Node_BinaryExpr }

type UnaryExpr struct {
	Expr
	Op    Token
	Right Expr
}

func (s *UnaryExpr) Type() NodeType { return Node_UnaryExpr }

type IdentiferLiteral struct {
	Expr
	Value string
}

func (s *IdentiferLiteral) Type() NodeType { return Node_Identifer }

type IntLiteral struct {
	Expr
	Value int64
}

func (s *IntLiteral) Type() NodeType { return Node_Int }

type FloatLiteral struct {
	Expr
	Value float64
}

func (s *FloatLiteral) Type() NodeType { return Node_Float }

type BooleanLiteral struct {
	Expr
	Value bool
}

func (s *BooleanLiteral) Type() NodeType { return Node_Bool }

type StringLiteral struct {
	Expr
	Value string
}

func (s *StringLiteral) Type() NodeType { return Node_String }

type NullLiteral struct {
	Expr
	Value interface{}
}

func (s *NullLiteral) Type() NodeType { return Node_Null }

type ObjType struct {
	Expr
	NodeType TokenType
}

func (s *ObjType) Type() NodeType { return Node_ObjType }