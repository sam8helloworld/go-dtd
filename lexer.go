package main

import (
	"github.com/pkg/errors"
)

var ErrElementTokenize = errors.New("failed to element tokenize")
var ErrEmptyTokenize = errors.New("failed to empty tokenize")
var ErrAttListTokenize = errors.New("failed to attlist tokenize")
var ErrDefaultValueTokenize = errors.New("failed to default value tokenize")
var ErrStringTokenize = errors.New("failed to string tokenize")
var ErrTagNecessityTokenize = errors.New("failed to tag necessity tokenize")

const (
	ExclamationSymbol       = '!'
	LeftAngleBracketSymbol  = '<'
	RightAngleBracketSymbol = '>'
	ElementOrEmptySymbol    = 'E'
	WhiteSpaceSymbol        = ' '
	WhiteSpaceTabSymbol     = '\t'
	WhiteSpaceCRSymbol      = '\r'
	WhiteSpaceLFSymbol      = '\n'
	LeftBracketSymbol       = '('
	RightBracketSymbol      = ')'
	CommaSymbol             = ','
	AsteriskSymbol          = '*'
	TagNeedSymbol           = '-'
	TagUnNeedSymbol         = 'O'
	AmpersandSymbol         = '&'
	VerticalLineSymbol      = '|'
	PlusSymbol              = '+'
	QuestionSymbol          = '?'
	MinusSymbol             = '-'
	AttListSymbol           = 'A'
	SharpSymbol             = '#'
	QuoteSymbol             = '\''
	DoubleQuoteSymbol       = '"'
)

type lexer struct {
	input        string
	position     int  // 読み込んでる文字のインデックス
	readPosition int  // 次に読み込む文字のインデックス
	ch           byte // 検査中の文字
}

func NewLexer(input string) *lexer {
	return &lexer{input: input}
}

func (l *lexer) Execute() ([]Token, error) {
	tokens := []Token{}
	for ch := l.readChar(); l.readPosition <= len(l.input); ch = l.readChar() {
		switch {
		case ch == LeftAngleBracketSymbol:
			tokens = append(tokens, Token{
				Type:    TokenType(LeftAngleBracket),
				Literal: string(ch),
			})
		case ch == RightAngleBracketSymbol:
			tokens = append(tokens, Token{
				Type:    TokenType(RightAngleBracket),
				Literal: string(ch),
			})
		case ch == ExclamationSymbol:
			tokens = append(tokens, Token{
				Type:    TokenType(Exclamation),
				Literal: string(ch),
			})
		case ch == ElementOrEmptySymbol:
			nextChar := l.peakChar()
			if nextChar == 'L' {
				token, err := l.elementTokenize()
				if err != nil {
					return nil, err
				}
				tokens = append(tokens, *token)
			}
			if nextChar == 'M' {
				token, err := l.emptyTokenize()
				if err != nil {
					return nil, err
				}
				tokens = append(tokens, *token)
			}
		case ch == WhiteSpaceSymbol || ch == WhiteSpaceTabSymbol || ch == WhiteSpaceCRSymbol || ch == WhiteSpaceLFSymbol:
			continue
		case ch == AttListSymbol:
			token, err := l.attListTokenize()
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, *token)
		case ch == LeftBracketSymbol:
			tokens = append(tokens, Token{
				Type:    LeftBracket,
				Literal: string(ch),
			})
		case ch == RightBracketSymbol:
			tokens = append(tokens, Token{
				Type:    RightBracket,
				Literal: string(ch),
			})
		case ch == CommaSymbol:
			tokens = append(tokens, Token{
				Type:    Comma,
				Literal: string(ch),
			})
		case ch == AmpersandSymbol:
			tokens = append(tokens, Token{
				Type:    Ampersand,
				Literal: string(ch),
			})
		case ch == AsteriskSymbol:
			tokens = append(tokens, Token{
				Type:    Asterisk,
				Literal: string(ch),
			})
		case ch == VerticalLineSymbol:
			tokens = append(tokens, Token{
				Type:    VerticalLine,
				Literal: string(ch),
			})
		case ch == PlusSymbol:
			tokens = append(tokens, Token{
				Type:    Plus,
				Literal: string(ch),
			})
		case ch == MinusSymbol:
			tokens = append(tokens, Token{
				Type:    Minus,
				Literal: string(ch),
			})
		case ch == QuoteSymbol || ch == DoubleQuoteSymbol:
			token, err := l.stringTokenize(ch)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, *token)
		case ch == SharpSymbol:
			token, err := l.defaulValueTokenize()
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, *token)
		case ch == QuestionSymbol:
			tokens = append(tokens, Token{
				Type:    Question,
				Literal: string(ch),
			})
		case ch == TagNeedSymbol:
			tokens = append(tokens, Token{
				Type:    TagNeed,
				Literal: string(ch),
			})
		case ch == TagUnNeedSymbol:
			tokens = append(tokens, Token{
				Type:    TagUnNeed,
				Literal: string(ch),
			})
		default:
			token, err := l.nameTokenize()
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, *token)
			continue
		}
	}
	return tokens, nil
}

func (l *lexer) elementTokenize() (*Token, error) {
	el := string(l.ch)
	for i := 0; i < 6; i++ {
		el += string(l.readChar())
	}
	if el == "ELEMENT" {
		return &Token{
			Type:    Element,
			Literal: el,
		}, nil
	}
	return nil, ErrElementTokenize
}

func (l *lexer) nameTokenize() (*Token, error) {
	name := string(l.ch)
	for {
		ch := l.peakChar()
		if ch == WhiteSpaceSymbol || ch == WhiteSpaceTabSymbol || ch == WhiteSpaceCRSymbol || ch == WhiteSpaceLFSymbol || ch == CommaSymbol || ch == RightBracketSymbol || ch == AsteriskSymbol || ch == AmpersandSymbol || ch == VerticalLineSymbol || ch == PlusSymbol || ch == QuestionSymbol {
			break
		}
		name += string(ch)
		l.readChar()
	}
	return &Token{
		Type:    Name,
		Literal: name,
	}, nil
}

func (l *lexer) emptyTokenize() (*Token, error) {
	em := string(l.ch)
	for i := 0; i < 4; i++ {
		em += string(l.readChar())
	}
	if em == "EMPTY" {
		return &Token{
			Type:    Empty,
			Literal: em,
		}, nil
	}
	return nil, ErrEmptyTokenize
}

func (l *lexer) attListTokenize() (*Token, error) {
	att := string(l.ch)
	for i := 0; i < 6; i++ {
		att += string(l.readChar())
	}
	if att == "ATTLIST" {
		return &Token{
			Type:    AttList,
			Literal: att,
		}, nil
	}
	return nil, ErrAttListTokenize
}

func (l *lexer) defaulValueTokenize() (*Token, error) {
	switch l.peakChar() {
	case 'I':
		imp := string(l.readChar())
		for i := 0; i < 6; i++ {
			imp += string(l.readChar())
		}
		if imp == "IMPLIED" {
			return &Token{
				Type:    DefaultValueImplied,
				Literal: "#IMPLIED",
			}, nil
		}
		return nil, ErrDefaultValueTokenize
	case 'R':
		req := string(l.readChar())
		for i := 0; i < 7; i++ {
			req += string(l.readChar())
		}
		if req == "REQUIRED" {
			return &Token{
				Type:    DefaultValueRequired,
				Literal: "#REQUIRED",
			}, nil
		}
		return nil, ErrDefaultValueTokenize
	case 'F':
		fix := string(l.readChar())
		for i := 0; i < 4; i++ {
			fix += string(l.readChar())
		}
		if fix == "FIXED" {
			return &Token{
				Type:    DefaultValueFixed,
				Literal: "#FIXED",
			}, nil
		}
		return nil, ErrDefaultValueTokenize
	default:
		return nil, ErrDefaultValueTokenize
	}
}

func (l *lexer) stringTokenize(quoteSymbol byte) (*Token, error) {
	str := ""
	for ch := l.readChar(); ch != 0; ch = l.readChar() {
		switch ch {
		case quoteSymbol:
			return &Token{
				Type:    String,
				Literal: str,
			}, nil
		default:
		}
		str += string(ch)
	}
	return nil, ErrStringTokenize
}

func (l *lexer) readChar() byte {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
	return l.ch
}

func (l *lexer) peakChar() byte {
	// 入力が終わったらchを0に
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
