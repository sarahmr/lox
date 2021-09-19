package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"bufio"

	"github.com/sarahmr/lox/scanner"
)

// don't call main function with arguments since the program is calling it
func main() {
	testPrinter()

	args := os.Args[1:]
	loxInterpreter := Lox{}

	if len(args) > 1 {
		fmt.Println("Usage: lox [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		loxInterpreter.runFile(args[0])
	} else {
		loxInterpreter.runPrompt()
	}
}

type Lox struct {
	hadError bool
}

func (l *Lox) runFile(path string) {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}
	// fmt.Print(string(file))
	l.run(string(file))
	if l.hadError {
		os.Exit(64)
	}
}

func (l *Lox) run(source string) {
	loxscanner := scanner.NewScanner(source, l.error)

	tokens := loxscanner.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}
}

func (l *Lox) runPrompt() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		l.run(text)
		l.hadError = false
	}
}

func (l *Lox) error(lineNumber int, message string) {
	l.report(lineNumber, "", message)
}

func (l *Lox) report(lineNumber int, location string, message string) {
	fmt.Printf("[line %d ] Error %s: %s \n", lineNumber, location, message)
	l.hadError = true
}

func testPrinter() {
	expr := BinaryExpr{
		Operator: scanner.Token{
			Type:    scanner.Star,
			Lexeme:  "*",
			Literal: nil,
			Line:    1,
		},
		Left: UnaryExpr{
			Operator: scanner.Token{
				Type:    scanner.Minus,
				Lexeme:  "-",
				Literal: nil,
				Line:    1,
			},
			Right: LiteralExpr{
				Value: scanner.FloatLiteral(123),
			},
		},
		Right: GroupingExpr{
			Expression: LiteralExpr{
				Value: scanner.FloatLiteral(45.67),
			},
		},
	}

	astPrinter := AstPrinter{}
	fmt.Println(astPrinter.Print(expr))
}
