package lexer

import (
	"github.com/croese/golox/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `(`

	expected := []struct {
		tokenType token.TokenType
		literal   string
		line      int
		column    int
		length    int
	}{}
}
