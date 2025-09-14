package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	"fmt"
)

type Parser struct {
	lex *lexer.Lexer
	curToken token.Token
	peekToken token.Token
	errors []string
}

func New(lex *lexer.Lexer) *Parser {
	parser := &Parser{
		lex: lex,
		errors: []string{},
	}

	// Read two tokens to fill current and peek
	parser.nextToken()
	parser.nextToken()

	return parser
}

func (parser *Parser) Errors() []string {
	return parser.errors
}

func (parser *Parser) peekError(tokType token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", tokType, parser.peekToken.Type)
	parser.errors = append(parser.errors, msg)
}

func (parser *Parser) nextToken() {
	parser.curToken = parser.peekToken
	parser.peekToken = parser.lex.NextToken()
}

func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for parser.curToken.Type != token.EOF {
		stmt := parser.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		parser.nextToken()
	}
	return program
}

func (parser *Parser) parseStatement() ast.Statement {
	switch parser.curToken.Type {
	case token.LET:
		return parser.parseLetStatement()
	default:
		return nil
	}
}

func (parser *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: parser.curToken}

	if !parser.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: parser.curToken, Value: parser.curToken.Literal}


	if !parser.expectPeek(token.ASSIGN) {
		return nil
	}

	for !parser.expectPeek(token.SEMICOLON) {
		parser.nextToken()
	}

	return stmt
}

func (parser *Parser) curTokenIs(tokType token.TokenType) bool {
	return parser.curToken.Type == tokType
}

func (parser *Parser) peekTokenIs(tokType token.TokenType) bool {
	return parser.peekToken.Type == tokType
}

func (parser *Parser) expectPeek(tokType token.TokenType) bool {
	if parser.peekTokenIs(tokType) {
		parser.nextToken()
		return true
	} else {
		parser.peekError(tokType)
		return false
	}
}
