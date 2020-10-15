package lang

import "fmt"

type Parser struct {
  lexer *Lexer

  currentState State

  currentStatement Statement
  statements []Statement
}

func NewParser(lexer *Lexer) *Parser {
  return &Parser{lexer, STATE_STMT_START, nil, []Statement{}}
}

func (this *Parser) Parse() ([]Statement, error) {
  for {
    token, err := this.lexer.Next()
    if err != nil {
      return nil, err
    }

    if token.Type == TOKEN_EOF {
      if this.currentState != STATE_STMT_START {
        return nil, &SyntaxError{token.Position, "Unexpected end of file"}
      }

      return this.statements, nil
    }

    if err := this.processToken(token); err != nil {
      return nil, err
    }
  }
}

func (this *Parser) processToken(token *Token) error {
  // fmt.Printf("Check token %-10v from current state %-25v", token.Type, this.currentState)

  for _, trans := range FSM[this.currentState] {
    if trans.TokenType == token.Type {
      // fmt.Printf("; Found next state: %s\n", trans.NextState)

      // Found appropriate transition
      nextStmt, err := trans.Callback(this.currentStatement, token)
      if err != nil {
        return err
      }
      this.currentStatement = nextStmt

      this.currentState = trans.NextState

      if this.currentState == STATE_STMT_START {
        this.statements = append(this.statements, this.currentStatement)
        this.currentStatement = nil
      }

      return nil
    }
  }

  return &SyntaxError{
    token.Position,
    fmt.Sprintf("Unexpected token '%s'", token.Literal),
  }
}
