package lang

import (
  "testing"
  "strings"
)

func TestLexerNext(t *testing.T) {
  cazes := []struct{
    str string
    typ TokenType
    lit string
  }{
    {"", TOKEN_EOF, ""},
    {"test", TOKEN_IDENTIFIER, "test"},
    {";", TOKEN_SEMI, ";"},
    {")", TOKEN_PAR_CLOSE, ")"},

    {"DimenSion", TOKEN_KW_DIMENSION, "DIMENSION"},
  }

  for _, caze := range cazes {
    lex := NewLexer(strings.NewReader(caze.str))
    tok, _ := lex.Next()

    if tok.Type != caze.typ {
      t.Errorf(
        "Bad token type, expected: '%s', actual: '%s' (from input: '%s')",
        caze.typ,
        tok.Type,
        caze.str,
      )
    }

    if tok.Literal != caze.lit {
      t.Errorf(
        "Bad token literal, expected: '%s', actual: '%s' (from input: '%s')",
        caze.lit,
        tok.Literal,
        caze.str,
      )
    }
  }
}
