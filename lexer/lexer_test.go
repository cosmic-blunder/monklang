package lexer

import (
	  "testing"

	  "github.com/cosmic-blunder/monklang/token"
)



func TestNextToken(t *testing.T){

      input:=`=+(){},;`


      tests:=[]struct {
	      expectedType token.TokenType
              expectedLiteral string           
      }{

                 { token.ASSIGN,"=" },
		 { token.PLUS,"+" },
		 { token.LPAREN , "(" },
		 { token.RPAREN ,")"},
		 { token.LBRACE,"{"},
		 { token.RBRACE,"}"},
		 { token.COMMA,","},
		 { token.SEMICOLON,";"},
		 { token.EOF,""},
      }


      l:=New(input)


      for i, tt := range tests {
           tok:=l.NextToken()

	   if tok.Type!=tt.expectedType{

              t.Fatalf("test[%d] -  tokentype wrong.expected=%q,got=%q",i,tt.expected,token.Type)
	   }

	   if tok.Literal != tt.expectedLiteral{
               t.Fatalf("test[%d] - literal wrong. expected=%q , got=%q",i,ii.expectedLiteral,tok.Lietral);
	   }
      }

}


