package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sarahmr/lox/scanner"
)

// don't call main function with arguments since the program is calling it
func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Usage: lox [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		// runPrompt()
	}
}

func runFile(path string) {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	// fmt.Print(string(file))
	run(string(file))
}

func run(source string) {
	loxscanner := scanner.NewScanner(source)

	tokens := loxscanner.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}
}
