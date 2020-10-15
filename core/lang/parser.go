package lang

import "fmt"

type Parser struct {
  lexer *Lexer
  currentState State
}

func NewParser(lexer *Lexer) *Parser {
  return &Parser{lexer, STATE_STMT_START}
}

func (this *Parser) Parse() error {
  for {
    token, err := this.lexer.Next()
    if err != nil {
      return err
    }

    if token.Type == TOKEN_EOF {
      if this.currentState != STATE_STMT_START {
        return &SyntaxError{token.Position, "Unexpected end of file"}
      }

      return nil
    }

    if err := this.processToken(token); err != nil {
      return err
    }
  }
}

func (this *Parser) processToken(token *Token) error {
  fmt.Printf("Check token %-10v from current state %-25v ", token.Type, this.currentState)

  for _, trans := range FSM[this.currentState] {
    if trans.TokenType == token.Type {
      // Found appropriate transition
      this.currentState = trans.NextState
      fmt.Printf("; Found next state: %s\n", trans.NextState)

      return nil
    }
  }

  return &SyntaxError{
    token.Position,
    fmt.Sprintf("Unexpected token '%s'", token.Literal),
  }
}
