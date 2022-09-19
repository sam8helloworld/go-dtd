package main

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	LeftAngleBracket     = "<"
	RightAngleBracket    = ">"
	Exclamation          = "!"
	Element              = "ELEMENT"
	Name                 = "Name"
	LeftBracket          = "("
	RightBracket         = ")"
	Comma                = ","
	Asterisk             = "*"
	TagNeed              = "-"
	TagUnNeed            = "O"
	Ampersand            = "&"
	VerticalLine         = "|"
	Plus                 = "+"
	Question             = "?"
	Empty                = "EMPTY"
	Minus                = "-"
	AttList              = "ATTLIST"
	DefaultValueImplied  = "#IMPLIED"
	DefaultValueRequired = "#REQUIRED"
	DefaultValueFixed    = "#FIXED"
	String               = "String"
	Entity               = "ENTITY"
	Percent              = "%"
)
