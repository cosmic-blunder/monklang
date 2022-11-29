package parser

import (
	"testing"

	"github.com/cosmic-blunder/monklang/ast"
	"github.com/cosmic-blunder/monklang/lexer"
)

func TestLetStatements(t *testing.T) {

	input := `
	   let x = 5;
	   let y= 10;
	   let foobar = 838383;
	`

	l := lexer.New((input))

	p := New(l)

	program := p.ParseProgram()

	if program == nil {

		t.Fatal("ParseProgram() return nil")
	}

	if len(program.Statements) != 3 {
		t.Fatal("Program.Statements  does not contain 3 statements got= %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{

		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {

		stmt := program.Statements[i]

		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}

}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {

	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLitral not 'let'. got =%q", s.TokenLiteral())
		return false
	}

	//letStmt,ok:=s.(*ast.LetStatement)

	return true
}