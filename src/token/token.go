package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const {
	// Special types
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifiers and literals
	IDENT = "IDENT" // var names like foo, bar, x, y
	INT = "INT" // numbers

	// Operators
	ASSIGN = "="
	PLUS = "+"

	// Delimeters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
}

