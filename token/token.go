package token

type TokenType string

type Token struct {

           Type TokenType
	   Literal string
}

const (

	ILLEGAL="ILLEGAL"
	EOF    ="EOF"

	//Identifier+ literals
	IDENT = "IDENT" //ADD,FOOBAR ,X ,Y ...
	INT   ="INT"

	//Operators
	ASSIGN="="
	PLUS="PLUS"


	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE="{"
	RBRACE="}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      ="LET"
) 
