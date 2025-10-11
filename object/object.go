package object

import (
	"bytes"
	"fmt"
	"monkey/ast"
	"strings"
)

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ = "ERROR"
	FUNCTION_OBJ = "FUNCTION"
	STRING_OBJ = "STRING"
	BUILTIN_OBJ = "BUILTIN"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}



type Integer struct {
	Value int64
}

type String struct {
	Value string
}

type Boolean struct {
	Value bool
}

type BuiltinFunction func(args ...Object) Object
type Builtin struct {
	Fn BuiltinFunction
}

type ReturnValue struct {
	Value Object
}

type Function struct {
	Parameters []*ast.Identifier
	Body *ast.BlockStatement
	Env *Environment
}

type Null struct{}

type Error struct {
	Message string
}

// Integer functions
func (integer *Integer) Inspect() string { return fmt.Sprintf("%d", integer.Value) }
func (integer *Integer) Type() ObjectType { return INTEGER_OBJ }

// String functions
func (str *String) Inspect() string { return str.Value }
func (str *String) Type() ObjectType { return STRING_OBJ }

// Bool functions
func (boolean *Boolean) Inspect() string { return fmt.Sprintf("%t", boolean.Value) }
func (boolean *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

// Built-in type functions
func (builtIn *Builtin) Inspect() string { return "builtin function" }
func (builtIn *Builtin) Type() ObjectType { return BUILTIN_OBJ }

// Return functions
func (retVal *ReturnValue) Inspect() string { return retVal.Inspect() }
func (retVal *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

// Null functions
func (null *Null) Inspect() string { return "null" }
func (null *Null) Type() ObjectType { return NULL_OBJ }

// Error functions
func (err *Error) Inspect() string { return "ERROR: " + err.Message }
func (err *Error) Type() ObjectType { return ERROR_OBJ }

// Function functions
func (fun *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}

	for _, p := range fun.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(fun.Body.String())
	out.WriteString("\n}")

	return out.String()
}
func (fun *Function) Type() ObjectType { return FUNCTION_OBJ }
