package scanner

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
}

func NewScanner(source string) Scanner {
	return Scanner{
		Source: source,
		line:   1,
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
	}
}

func (s *Scanner) advance() string {
	s.current++
	return string(s.Source[s.current-1])
}

func (s *Scanner) addToken(t TokenType) {
	text := s.Source[s.start:s.current]

	s.Tokens = append(s.Tokens, Token{
		Type:   t,
		Lexeme: text,
		Line:   s.line,
	})
}
