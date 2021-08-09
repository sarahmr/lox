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
	loxscanner := scanner.NewScanner(source)

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
	fmt.Printf("[line %d ] Error %s: %s", lineNumber, location, message)
	l.hadError = true
}
