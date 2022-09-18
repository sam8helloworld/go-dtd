package main

import (
	"github.com/pkg/errors"
)

var ErrElementTokenize = errors.New("failed to element tokenize")

const (
	ExclamationSymbol       = '!'
	LeftAngleBracketSymbol  = '<'
	RightAngleBracketSymbol = '>'
	ElementSymbol           = 'E'
	WhiteSpaceSymbol        = ' '
	LeftBracketSymbol       = '('
	RightBracketSymbol      = ')'
	CommaSymbol             = ','
	AsteriskSymbol          = '*'
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
		case ch == ElementSymbol:
			token, err := l.elementTokenize()
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, *token)
		case ch == WhiteSpaceSymbol:
			continue
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
		case ch == AsteriskSymbol:
			tokens = append(tokens, Token{
				Type:    Asterisk,
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
		if ch == WhiteSpaceSymbol || ch == CommaSymbol || ch == RightBracketSymbol || ch == AsteriskSymbol {
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
