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
		switch ch {
		case LeftAngleBracketSymbol:
			tokens = append(tokens, Token{
				Type:    TokenType(LeftAngleBracket),
				Literal: string(ch),
			})
		case RightAngleBracketSymbol:
			tokens = append(tokens, Token{
				Type:    TokenType(RightAngleBracket),
				Literal: string(ch),
			})
		case ExclamationSymbol:
			tokens = append(tokens, Token{
				Type:    TokenType(Exclamation),
				Literal: string(ch),
			})
		case ElementSymbol:
			token, err := l.elementTokenize()
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, *token)
		default:
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
