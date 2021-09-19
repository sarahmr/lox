package main

import "github.com/sarahmr/lox/scanner"

type ExprVisitor interface {
	VisitBinaryExpr(e BinaryExpr) interface{}
	VisitUnaryExpr(e UnaryExpr) interface{}
	VisitGroupingExpr(e GroupingExpr) interface{}
	VisitLiteralExpr(e LiteralExpr) interface{}
}

type AstPrinter struct{}

func (a AstPrinter) Print(expr Expression) string {
	result := expr.Accept(a)
	str, ok := result.(string)
	if !ok {
		// crash program
		panic("AstPrinter couldn't cast interface to string")
	}

	return str
}

func (a AstPrinter) VisitBinaryExpr(expr BinaryExpr) interface{} {
	return a.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}
func (a AstPrinter) VisitGroupingExpr(expr GroupingExpr) interface{} {
	return a.parenthesize("group", expr.Expression)
}
func (a AstPrinter) VisitLiteralExpr(expr LiteralExpr) interface{} {
	if expr.Value == nil {
		return "nil"
	}
	return expr.Value.ToString()
}
func (a AstPrinter) VisitUnaryExpr(expr UnaryExpr) interface{} {
	return a.parenthesize(expr.Operator.Lexeme, expr.Right)
}

func (a AstPrinter) parenthesize(name string, exprs ...Expression) string {
	s := "(" + name

	for _, expr := range exprs {
		result := expr.Accept(a)
		str, ok := result.(string)
		if !ok {
			// crash program
			panic("AstPrinter couldn't cast interface to string")
		}
		s += " " + str
	}

	s += ")"
	return s
}

type Expression interface {
	IsExpression()
	Accept(visitor ExprVisitor) interface{}
}

type BinaryExpr struct {
	Left     Expression
	Operator scanner.Token
	Right    Expression
}

func (BinaryExpr) IsExpression() {}
func (e BinaryExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitBinaryExpr(e)
}

type UnaryExpr struct {
	Operator scanner.Token
	Right    Expression
}

func (UnaryExpr) IsExpression() {}
func (e UnaryExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitUnaryExpr(e)
}

type GroupingExpr struct {
	Expression Expression
}

func (GroupingExpr) IsExpression() {}
func (e GroupingExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitGroupingExpr(e)
}

type LiteralExpr struct {
	Value scanner.Literal
}

func (LiteralExpr) IsExpression() {}
func (e LiteralExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitLiteralExpr(e)
}
