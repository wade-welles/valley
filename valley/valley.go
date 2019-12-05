package valley

import (
	"go/ast"
	"unsafe"
)

// ConstraintViolation ...
type ConstraintViolation struct {
	Field   string                 `json:"field"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details"`
}

// Constraint ...
type Constraint func(value Value, fieldType ast.Expr, opts interface{}) (ConstraintOutput, error)

// ConstraintOutput ...
type ConstraintOutput struct {
	Imports []Import
	Code    string
}

// Import represents information about a Go import that Valley uses to generate code.
type Import struct {
	Path  string
	Alias string
}

// NewImport returns a new Import value.
func NewImport(path, alias string) Import {
	return Import{
		Path:  path,
		Alias: alias,
	}
}

// Value ...
// TODO: Move to a generator package or something? Generator is maybe a poor name, because it makes
// a pretty code type name. Using something like `code` feels a bit crap though too.
type Value struct {
	TypeName  string
	Receiver  string
	FieldName string
	VarName   string
	Path      string
}

// Package ...
type Package struct {
	Name    string
	Methods Methods
	Structs Structs
}

// Methods is a map from struct name to Method.
type Methods map[string][]Method

// Method represents the information we need about a method in some Go source code.
type Method struct {
	Receiver string
	Name     string
}

// Structs is a map from struct name to Struct.
type Structs map[string]Struct

// Struct represents the information we need about a struct in some Go source code.
type Struct struct {
	Name   string
	Node   *ast.StructType
	Fields Fields
}

// Fields is a map from struct field name to Field.
type Fields map[string]Field

// Field represents the information we need about a struct field in some Go source code.
type Field struct {
	Name string
	Type ast.Expr
}

// Btos returns the given byte slice as a string, without allocating any new memory.
func Btos(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}
