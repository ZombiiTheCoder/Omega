package main

import (
	"fmt"
	"omega/src"
	"os"
)

func main() {
	file := os.Args[1]
	StartPoint := os.Args[2]
	fmt.Println(StartPoint, file)
	// Cmd Format
	// omega example/test.om test.Main

	lex := src.CreateLexer(file)
	par := src.CreateParser(lex)
	fmt.Println(par.ProduceAst())
	// cmp := src.CreateCompiler(par, src.Bytecode)
	// bvm := src.CreateVirtualMachine(cmp)

	// acp := src.CreateCompiler(par, src.Assembly)
}