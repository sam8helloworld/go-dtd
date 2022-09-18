package main

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	LeftAngleBracket  = "<"
	RightAngleBracket = ">"
	Exclamation       = "!"
	Element           = "ELEMENT"
	Name              = "Name"
	LeftBracket       = "("
	RightBracket      = ")"
	Comma             = ","
	Asterisk          = "*"
	TagNeed           = "-"
	TagUnNeed         = "O"
)
