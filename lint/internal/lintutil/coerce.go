package lintutil

import (
	"go/ast"
	"go/token"
)

var (
	nilIdent        = &ast.Ident{}
	nilSelectorExpr = &ast.SelectorExpr{}
	nilUnaryExpr    = &ast.UnaryExpr{}
	nilBinaryExpr   = &ast.BinaryExpr{}
	nilCallExpr     = &ast.CallExpr{}
	nilParenExpr    = &ast.ParenExpr{}
	nilAssignStmt   = &ast.AssignStmt{}
)

// IsNil reports whether x is nil.
// Unlike simple nil check, also detects nil AST sentinels.
func IsNil(x ast.Node) bool {
	switch x := x.(type) {
	case *ast.Ident:
		return x == nilIdent || x == nil
	case *ast.SelectorExpr:
		return x == nilSelectorExpr || x == nil
	case *ast.UnaryExpr:
		return x == nilUnaryExpr || x == nil
	case *ast.BinaryExpr:
		return x == nilBinaryExpr || x == nil
	case *ast.CallExpr:
		return x == nilCallExpr || x == nil
	case *ast.ParenExpr:
		return x == nilParenExpr || x == nil
	case *ast.AssignStmt:
		return x == nilAssignStmt || x == nil

	default:
		return x == nil
	}
}

// AsUnaryExprOp is like AsUnaryExpr, but also checks for op token.
func AsUnaryExprOp(x ast.Node, op token.Token) *ast.UnaryExpr {
	e, ok := x.(*ast.UnaryExpr)
	if !ok || e.Op != op {
		return nilUnaryExpr
	}
	return e
}

// AsBinaryExprOp is like AsBinaryExpr, but also checks for op token.
func AsBinaryExprOp(x ast.Node, op token.Token) *ast.BinaryExpr {
	e, ok := x.(*ast.BinaryExpr)
	if !ok || e.Op != op {
		return nilBinaryExpr
	}
	return e
}
