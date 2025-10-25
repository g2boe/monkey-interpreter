package ast

import (
	"monkey/token"
	"bytes"
	"strings"
)

// Types

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type LetStatement struct {
	Token token.Token
	Name *Identifier
	Value Expression
}

type ReturnStatement struct {
	Token token.Token
	ReturnValue Expression
}

type ExpressionStatement struct {
	Token token.Token
	Expression Expression
}

type PrefixExpression struct {
	Token token.Token
	Operator string
	Right Expression
}

type InfixExpression struct {
	Token token.Token
	Left Expression
	Operator string
	Right Expression
}

type IfExpression struct {
	Token token.Token
	Condition Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

type BlockStatement struct {
	Token token.Token
	Statements []Statement
}

type FunctionLiteral struct {
	Token token.Token
	Parameters []*Identifier
	Body *BlockStatement
}

type CallExpression struct {
	Token token.Token
	Function Expression
	Arguments []Expression
}

type ArrayLiteral struct {
	Token token.Token
	Elements []Expression
}

type IndexExpression struct {
	Token token.Token
	Left Expression
	Index Expression
}

type StringLiteral struct {
	Token token.Token
	Value string
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

type Identifier struct {
	Token token.Token
	Value string
}

type Boolean struct {
	Token token.Token
	Value bool
}

// Program functions
func (program *Program) TokenLiteral() string {
	if len(program.Statements) > 0 {
		return program.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (program *Program) String() string {
	var out bytes.Buffer

	for _, stmt := range program.Statements {
		out.WriteString(stmt.String())
	}
	return out.String()
}

// let statement functions
func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// identifier functions
func (identifier *Identifier) expressionNode() {}
func (identifier *Identifier) TokenLiteral() string { return identifier.Token.Literal }
func (identifier *Identifier) String() string { return identifier.Value }

func (returnStmt *ReturnStatement) statementNode() {}
func (returnStmt *ReturnStatement) TokenLiteral() string { return returnStmt.Token.Literal }

func (returnStmt *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(returnStmt.TokenLiteral() + " ")
	if returnStmt.ReturnValue != nil {
		out.WriteString(returnStmt.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// Expression statement functions
func (expressionStmt *ExpressionStatement) statementNode() {}
func (expressionStmt *ExpressionStatement) TokenLiteral() string { return expressionStmt.Token.Literal }

func (expressionStmt *ExpressionStatement) String() string {
	if expressionStmt.Expression != nil {
		return expressionStmt.Expression.String()
	}

	return ""
}

// Prefix expression functions
func (prefixExpression *PrefixExpression) expressionNode() {}
func (prefixExpression *PrefixExpression) TokenLiteral() string { return prefixExpression.Token.Literal }
func (prefixExpression *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(prefixExpression.Operator)
	out.WriteString(prefixExpression.Right.String())
	out.WriteString(")")

	return out.String()
}

// Infix expression functions
func (infixExpression *InfixExpression) expressionNode() {}
func (infixExpression *InfixExpression) TokenLiteral() string { return infixExpression.Token.Literal }
func (infixExpression *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(infixExpression.Left.String())
	out.WriteString(" " + infixExpression.Operator + " ")
	out.WriteString(infixExpression.Right.String())
	out.WriteString(")")

	return out.String()
}

// If expression functions
func (ifExpression *IfExpression) expressionNode() {}
func (ifExpression *IfExpression) TokenLiteral() string { return ifExpression.Token.Literal }
func (ifExpression *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ifExpression.Condition.String())
	out.WriteString(" ")
	out.WriteString(ifExpression.Consequence.String())

	if ifExpression.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ifExpression.Alternative.String())
	}

	return out.String()
}

// If expression functions
func (blockStatement *BlockStatement) expressionNode() {}
func (blockStatement *BlockStatement) TokenLiteral() string { return blockStatement.Token.Literal }
func (blockStatement *BlockStatement) String() string {
	var out bytes.Buffer

	for _, stmt := range blockStatement.Statements {
		out.WriteString(stmt.String())
	}

	return out.String()
}

// Function literal functions
func (funcLiteral *FunctionLiteral) expressionNode() {}
func (funcLiteral *FunctionLiteral) TokenLiteral() string { return funcLiteral.Token.Literal }
func (funcLiteral *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range funcLiteral.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(funcLiteral.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(funcLiteral.Body.String())

	return out.String()
}


// Call expression functions
func (callExpression *CallExpression) expressionNode() {}
func (callExpression *CallExpression) TokenLiteral() string { return callExpression.Token.Literal }
func (callExpression *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range callExpression.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(callExpression.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

// Array literal functions
func (array *ArrayLiteral) expressionNode() {}
func (array *ArrayLiteral) TokenLiteral() string { return array.Token.Literal }
func (array *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range array.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

// Index expression functions
func (index *IndexExpression) expressionNode() {}
func (index *IndexExpression) TokenLiteral() string { return index.Token.Literal }
func (index *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(index.Left.String())
	out.WriteString("[")
	out.WriteString(index.Index.String())
	out.WriteString("])")

	return out.String()
}

// String literal functions
func (str *StringLiteral) expressionNode() {}
func (str *StringLiteral) TokenLiteral() string { return str.Token.Literal }
func (str *StringLiteral) String() string { return str.Token.Literal }

// Integer literal functions
func (intLiteral *IntegerLiteral) expressionNode() {}
func (intLiteral *IntegerLiteral) TokenLiteral() string { return intLiteral.Token.Literal }
func (intLiteral *IntegerLiteral) String() string {return intLiteral.Token.Literal}

// Bool functions
func (boolean *Boolean) expressionNode() {}
func (boolean *Boolean) TokenLiteral() string { return boolean.Token.Literal }
func (boolean *Boolean) String() string {return boolean.Token.Literal}
