package parser

import (
	"github.com/cosmic-blunder/monklang/ast"
	"github.com/cosmic-blunder/monklang/lexer"
	"github.com/cosmic-blunder/monklang/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	//for setting current token and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {

	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {

	return nil
}
