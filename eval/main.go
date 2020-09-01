package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strconv"
)

func main() {
	v, err := eval("1 + 2 * 3")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// 7
	fmt.Println(v)
}

func eval(s string) (float64, error) {
	expr, err := parser.ParseExpr(s)
	if err != nil {
		return 0, err
	}
	return evalExpr(expr)
}

func evalExpr(expr ast.Expr) (float64, error) {
	switch expr := expr.(type) {
	case *ast.BasicLit:
		v, err := strconv.ParseFloat(expr.Value, 64)
		if err != nil {
			return 0, nil
		}
		return v, nil
	case *ast.UnaryExpr: // 単項演算子
		if expr.Op != token.SUB {
			return 0, errors.New("unexpected op")
		}
		x, err := evalExpr(expr.X)
		if err != nil {
			return 0, err
		}
		return -x, nil
	}
	return 0, errors.New("unexpected expr")
}

func evalUnaryExpr(expr *ast.UnaryExpr) (float64, error) {
	if expr.Op != token.SUB {
		return 0, errors.New("unexpected op")
	}
	x, err := evalExpr(expr.X)
	if err != nil {
		return 0, err
	}
	return -x, nil
}

func evalBinaryExpr(expr *ast.BinaryExpr) (float64, error) {
	x, err := evalExpr(expr.X)
	if err != nil {
		return 0, nil
	}
	y, err := evalExpr(expr.Y)
	if err != nil {
		return 0, nil
	}
	switch expr.Op {
	case token.ADD:
		return x + y, nil
	case token.SUB:
		return x - y, nil
	case token.MUL:
		return x * y, nil
	case token.QUO:
		return x / y, nil
	}
	return 0, errors.New("unexpected expr")
}
