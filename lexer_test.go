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
			name:  "成功",
			input: "<!ELEMENT person （name,age,license*）>",
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
					Type:    RightAngleBracket,
					Literal: ">",
				},
			},
			wantErr: nil,
		},
		{
			name:    "ELEMENT要素名が間違っていてエラーが発生する",
			input:   "<!ELEMINT person （name,age,license*）>",
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
