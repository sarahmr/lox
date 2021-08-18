package scanner

import (
	"fmt"
)

// declares a new type
type TokenType string

const (
	// single-character tokens
	LeftParen  TokenType = "LeftParen"
	RightParen TokenType = "RightParen"
	LeftBrace  TokenType = "LeftBrace"
	RightBrace TokenType = "RightBrace"
	Comma      TokenType = "Comma"
	Dot        TokenType = "Dot"
	Minus      TokenType = "Minus"
	Plus       TokenType = "Plus"
	SemiColon  TokenType = "SemiColon"
	Slash      TokenType = "Slash"
	Star       TokenType = "Star"

	// one or two character tokens
	Bang         TokenType = "Bang"
	BangEqual    TokenType = "BangEqual"
	Equal        TokenType = "Equal"
	EqualEqual   TokenType = "EqualEqual"
	Greater      TokenType = "Greater"
	GreaterEqual TokenType = "GreaterEqual"
	Less         TokenType = "Less"
	LessEqual    TokenType = "LessEqual"

	// Literals
	Identifier TokenType = "Identifier"
	String     TokenType = "String"
	Number     TokenType = "Number"

	// Keywords
	And    TokenType = "And"
	Class  TokenType = "Class"
	Else   TokenType = "Else"
	False  TokenType = "False"
	Fun    TokenType = "Fun"
	For    TokenType = "For"
	If     TokenType = "If"
	Nil    TokenType = "Nil"
	Or     TokenType = "Or"
	Print  TokenType = "Print"
	Return TokenType = "Return"
	Super  TokenType = "Super"
	This   TokenType = "This"
	True   TokenType = "True"
	Var    TokenType = "Var"
	While  TokenType = "While"

	EOF TokenType = "EOF"
)

type Literal interface {
	IsLiteral()
	ToString() string
}

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal Literal
	Line    int
}

func (t Token) ToString() string {
	return string(t.Type) + " " + t.Lexeme + " " + t.Literal.ToString()
}

type Scanner struct {
	Source  string
	Tokens  []Token
	start   int
	current int
	line    int
	onError func(lineNumber int, message string)
}

func NewScanner(source string, onError func(lineNumber int, message string)) Scanner {
	return Scanner{
		Source:  source,
		line:    1,
		onError: onError,
	}
}

func (s *Scanner) ScanTokens() []Token {

	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.Tokens = append(s.Tokens, Token{
		Type: EOF,
		Line: s.line,
	})
	return s.Tokens
}

func (s Scanner) isAtEnd() bool {
	return s.current >= len(s.Source)
}

func (s *Scanner) scanToken() {
	c := s.advance()

	switch c {
	case "(":
		s.addToken(LeftParen)
	case ")":
		s.addToken(RightParen)
	case "{":
		s.addToken(LeftBrace)
	case "}":
		s.addToken(RightBrace)
	case ",":
		s.addToken(Comma)
	case ".":
		s.addToken(Dot)
	case "-":
		s.addToken(Minus)
	case "+":
		s.addToken(Plus)
	case ";":
		s.addToken(SemiColon)
	case "*":
		s.addToken(Star)
	case "!":
		t := Bang
		if s.match('=') {
			t = BangEqual
		}
		s.addToken(t)
	case "=":
		t := Equal
		if s.match('=') {
			t = EqualEqual
		}
		s.addToken(t)
	case "<":
		t := Less
		if s.match('=') {
			t = LessEqual
		}
		s.addToken(t)
	case ">":
		t := Greater
		if s.match('=') {
			t = GreaterEqual
		}
		s.addToken(t)
	default:
		s.onError(s.line, fmt.Sprintf("Unexpected character: %s.", c))
	}
}

func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}
	if s.Source[s.current] != byte(expected) {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) advance() string {
	curr := string(s.Source[s.current])
	s.current++
	return curr
}

func (s *Scanner) addTokenWithLiteral(t TokenType, literal Literal) {
	text := s.Source[s.start:s.current]

	s.Tokens = append(s.Tokens, Token{
		Type:    t,
		Lexeme:  text,
		Literal: literal,
		Line:    s.line,
	})
}

func (s *Scanner) addToken(t TokenType) {
	s.addTokenWithLiteral(t, nil)
}
