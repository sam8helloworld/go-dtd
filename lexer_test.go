package main

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestElementLexer(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []Token
		wantErr error
	}{
		{
			name:  "成功ケース_子要素の数が1つ",
			input: "<!ELEMENT person - O (name)>",
			want: []Token{
				{
					Type:    LeftAngleBracket,
					Literal: "<",
				},
				{
					Type:    Exclamation,
					Literal: "!",
				},
				{
					Type:    Element,
					Literal: "ELEMENT",
				},
				{
					Type:    Name,
					Literal: "person",
				},
				{
					Type:    TagNeed,
					Literal: "-",
				},
				{
					Type:    TagUnNeed,
					Literal: "O",
				},
				{
					Type:    LeftBracket,
					Literal: "(",
				},
				{
					Type:    Name,
					Literal: "name",
				},
				{
					Type:    RightBracket,
					Literal: ")",
				},
				{
					Type:    RightAngleBracket,
					Literal: ">",
				},
			},
			wantErr: nil,
		},
		{
			name:  "成功ケース_子要素の数が2つ以上かつカンマで区切り",
			input: "<!ELEMENT person - O (name,age)>",
			want: []Token{
				{
					Type:    LeftAngleBracket,
					Literal: "<",
				},
				{
					Type:    Exclamation,
					Literal: "!",
				},
				{
					Type:    Element,
					Literal: "ELEMENT",
				},
				{
					Type:    Name,
					Literal: "person",
				},
				{
					Type:    TagNeed,
					Literal: "-",
				},
				{
					Type:    TagUnNeed,
					Literal: "O",
				},
				{
					Type:    LeftBracket,
					Literal: "(",
				},
				{
					Type:    Name,
					Literal: "name",
				},
				{
					Type:    Comma,
					Literal: ",",
				},
				{
					Type:    Name,
					Literal: "age",
				},
				{
					Type:    RightBracket,
					Literal: ")",
				},
				{
					Type:    RightAngleBracket,
					Literal: ">",
				},
			},
			wantErr: nil,
		},
		{
			name:  "成功ケース_子要素の数が2つ以上かつアンパサンドで区切り",
			input: "<!ELEMENT person - O (name&age)>",
			want: []Token{
				{
					Type:    LeftAngleBracket,
					Literal: "<",
				},
				{
					Type:    Exclamation,
					Literal: "!",
				},
				{
					Type:    Element,
					Literal: "ELEMENT",
				},
				{
					Type:    Name,
					Literal: "person",
				},
				{
					Type:    TagNeed,
					Literal: "-",
				},
				{
					Type:    TagUnNeed,
					Literal: "O",
				},
				{
					Type:    LeftBracket,
					Literal: "(",
				},
				{
					Type:    Name,
					Literal: "name",
				},
				{
					Type:    Ampersand,
					Literal: "&",
				},
				{
					Type:    Name,
					Literal: "age",
				},
				{
					Type:    RightBracket,
					Literal: ")",
				},
				{
					Type:    RightAngleBracket,
					Literal: ">",
				},
			},
			wantErr: nil,
		},
		{
			name:  "成功ケース_子要素の数が1つでアスタリスクで修飾",
			input: "<!ELEMENT person - O (name)*>",
			want: []Token{
				{
					Type:    LeftAngleBracket,
					Literal: "<",
				},
				{
					Type:    Exclamation,
					Literal: "!",
				},
				{
					Type:    Element,
					Literal: "ELEMENT",
				},
				{
					Type:    Name,
					Literal: "person",
				},
				{
					Type:    TagNeed,
					Literal: "-",
				},
				{
					Type:    TagUnNeed,
					Literal: "O",
				},
				{
					Type:    LeftBracket,
					Literal: "(",
				},
				{
					Type:    Name,
					Literal: "name",
				},
				{
					Type:    RightBracket,
					Literal: ")",
				},
				{
					Type:    Asterisk,
					Literal: "*",
				},
				{
					Type:    RightAngleBracket,
					Literal: ">",
				},
			},
			wantErr: nil,
		},
		{
			name:  "成功ケース_子要素の数が2つ以上かつ縦線で区切り",
			input: "<!ELEMENT person - O (name|age)>",
			want: []Token{
				{
					Type:    LeftAngleBracket,
					Literal: "<",
				},
				{
					Type:    Exclamation,
					Literal: "!",
				},
				{
					Type:    Element,
					Literal: "ELEMENT",
				},
				{
					Type:    Name,
					Literal: "person",
				},
				{
					Type:    TagNeed,
					Literal: "-",
				},
				{
					Type:    TagUnNeed,
					Literal: "O",
				},
				{
					Type:    LeftBracket,
					Literal: "(",
				},
				{
					Type:    Name,
					Literal: "name",
				},
				{
					Type:    VerticalLine,
					Literal: "|",
				},
				{
					Type:    Name,
					Literal: "age",
				},
				{
					Type:    RightBracket,
					Literal: ")",
				},
				{
					Type:    RightAngleBracket,
					Literal: ">",
				},
			},
			wantErr: nil,
		},
		{
			name:  "成功ケース_子要素の数が1つかつプラスで装飾",
			input: "<!ELEMENT person - O (name)+>",
			want: []Token{
				{
					Type:    LeftAngleBracket,
					Literal: "<",
				},
				{
					Type:    Exclamation,
					Literal: "!",
				},
				{
					Type:    Element,
					Literal: "ELEMENT",
				},
				{
					Type:    Name,
					Literal: "person",
				},
				{
					Type:    TagNeed,
					Literal: "-",
				},
				{
					Type:    TagUnNeed,
					Literal: "O",
				},
				{
					Type:    LeftBracket,
					Literal: "(",
				},
				{
					Type:    Name,
					Literal: "name",
				},
				{
					Type:    RightBracket,
					Literal: ")",
				},
				{
					Type:    Plus,
					Literal: "+",
				},
				{
					Type:    RightAngleBracket,
					Literal: ">",
				},
			},
			wantErr: nil,
		},
		{
			name:  "成功ケース_子要素の数が1つかつはてなで装飾",
			input: "<!ELEMENT person - O (name)?>",
			want: []Token{
				{
					Type:    LeftAngleBracket,
					Literal: "<",
				},
				{
					Type:    Exclamation,
					Literal: "!",
				},
				{
					Type:    Element,
					Literal: "ELEMENT",
				},
				{
					Type:    Name,
					Literal: "person",
				},
				{
					Type:    TagNeed,
					Literal: "-",
				},
				{
					Type:    TagUnNeed,
					Literal: "O",
				},
				{
					Type:    LeftBracket,
					Literal: "(",
				},
				{
					Type:    Name,
					Literal: "name",
				},
				{
					Type:    RightBracket,
					Literal: ")",
				},
				{
					Type:    Question,
					Literal: "?",
				},
				{
					Type:    RightAngleBracket,
					Literal: ">",
				},
			},
			wantErr: nil,
		},
		{
			name:  "成功ケース_子要素がEMPTY",
			input: "<!ELEMENT person - O EMPTY>",
			want: []Token{
				{
					Type:    LeftAngleBracket,
					Literal: "<",
				},
				{
					Type:    Exclamation,
					Literal: "!",
				},
				{
					Type:    Element,
					Literal: "ELEMENT",
				},
				{
					Type:    Name,
					Literal: "person",
				},
				{
					Type:    TagNeed,
					Literal: "-",
				},
				{
					Type:    TagUnNeed,
					Literal: "O",
				},
				{
					Type:    Empty,
					Literal: "EMPTY",
				},
				{
					Type:    RightAngleBracket,
					Literal: ">",
				},
			},
			wantErr: nil,
		},
		{
			name:  "成功ケース_子要素の種類が2つ以上かつ1種類は0個以上を許容",
			input: "<!ELEMENT person - O (name) +(age)>",
			want: []Token{
				{
					Type:    LeftAngleBracket,
					Literal: "<",
				},
				{
					Type:    Exclamation,
					Literal: "!",
				},
				{
					Type:    Element,
					Literal: "ELEMENT",
				},
				{
					Type:    Name,
					Literal: "person",
				},
				{
					Type:    TagNeed,
					Literal: "-",
				},
				{
					Type:    TagUnNeed,
					Literal: "O",
				},
				{
					Type:    LeftBracket,
					Literal: "(",
				},
				{
					Type:    Name,
					Literal: "name",
				},
				{
					Type:    RightBracket,
					Literal: ")",
				},
				{
					Type:    Plus,
					Literal: "+",
				},
				{
					Type:    LeftBracket,
					Literal: "(",
				},
				{
					Type:    Name,
					Literal: "age",
				},
				{
					Type:    RightBracket,
					Literal: ")",
				},
				{
					Type:    RightAngleBracket,
					Literal: ">",
				},
			},
			wantErr: nil,
		},
		{
			name:  "成功ケース_子要素の種類が2つ以上かつ1種類は0個以上を許容しない",
			input: "<!ELEMENT person - O (name) -(age)>",
			want: []Token{
				{
					Type:    LeftAngleBracket,
					Literal: "<",
				},
				{
					Type:    Exclamation,
					Literal: "!",
				},
				{
					Type:    Element,
					Literal: "ELEMENT",
				},
				{
					Type:    Name,
					Literal: "person",
				},
				{
					Type:    TagNeed,
					Literal: "-",
				},
				{
					Type:    TagUnNeed,
					Literal: "O",
				},
				{
					Type:    LeftBracket,
					Literal: "(",
				},
				{
					Type:    Name,
					Literal: "name",
				},
				{
					Type:    RightBracket,
					Literal: ")",
				},
				{
					Type:    Minus,
					Literal: "-",
				},
				{
					Type:    LeftBracket,
					Literal: "(",
				},
				{
					Type:    Name,
					Literal: "age",
				},
				{
					Type:    RightBracket,
					Literal: ")",
				},
				{
					Type:    RightAngleBracket,
					Literal: ">",
				},
			},
			wantErr: nil,
		},
		{
			name:    "ELEMENT要素名が間違っていてエラーが発生する",
			input:   "<!ELEMINT person - O (name,age,license*）>",
			want:    nil,
			wantErr: ErrElementTokenize,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := NewLexer(tt.input)
			got, err := sut.Execute()
			if err != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("error mismatch want: %v, but got %v", tt.wantErr, err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestAttListLexer(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []Token
		wantErr error
	}{
		{
			name: "成功ケース_属性の数が1つ",
			input: `
<!ATTLIST HTML
lang    NAME
>
			`,
			want: []Token{
				{
					Type:    LeftAngleBracket,
					Literal: "<",
				},
				{
					Type:    Exclamation,
					Literal: "!",
				},
				{
					Type:    AttList,
					Literal: "ATTLIST",
				},
				{
					Type:    Name,
					Literal: "HTML",
				},
				{
					Type:    Name,
					Literal: "lang",
				},
				{
					Type:    Name,
					Literal: "NAME",
				},
				{
					Type:    RightAngleBracket,
					Literal: ">",
				},
			},
			wantErr: nil,
		},
		{
			name: "成功ケース_属性の数が1つで属性の説明付き",
			input: `
<!ATTLIST HTML
lang    NAME      #IMPLIED
>
			`,
			want: []Token{
				{
					Type:    LeftAngleBracket,
					Literal: "<",
				},
				{
					Type:    Exclamation,
					Literal: "!",
				},
				{
					Type:    AttList,
					Literal: "ATTLIST",
				},
				{
					Type:    Name,
					Literal: "HTML",
				},
				{
					Type:    Name,
					Literal: "lang",
				},
				{
					Type:    Name,
					Literal: "NAME",
				},
				{
					Type:    DefaultValueImplied,
					Literal: "#IMPLIED",
				},
				{
					Type:    RightAngleBracket,
					Literal: ">",
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := NewLexer(tt.input)
			got, err := sut.Execute()
			if err != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("error mismatch want: %v, but got %v", tt.wantErr, err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("mismatch (-got +want):\n%s", diff)
			}
		})
	}

}
