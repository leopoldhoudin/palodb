package lang

import (
  "io"
  "bufio"
  "bytes"
  "strings"
)

var eof = rune(0)

type Lexer struct {
  currentPosition Position
  input *bufio.Reader
  eof bool
}

func NewLexer(input io.Reader) *Lexer {
  return &Lexer{
    currentPosition: Position{Line: 1, Column: 0},
    input: bufio.NewReader(input),
    eof: false,
  }
}

func (this *Lexer) Next() (*Token, error) {
  for {
    if chr, err := this.read(); err != nil {
      return nil, err
    } else {
      switch chr {
      case eof:
        if this.eof {
          return nil, io.EOF
        }

        this.eof = true
        return NewToken(this.currentPosition, TOKEN_EOF, "")

  		case '\n':
  			this.newLine()
        break

  		case ';':
  			return NewToken(this.currentPosition, TOKEN_SEMI, string(chr))

  		case ',':
  			return NewToken(this.currentPosition, TOKEN_COMA, string(chr))

  		case '(':
  			return NewToken(this.currentPosition, TOKEN_PAR_OPEN, string(chr))

  		case ')':
  			return NewToken(this.currentPosition, TOKEN_PAR_CLOSE, string(chr))

      case '/':
        if chr, err := this.read(); err != nil {
          return nil, err
        } else if chr == '/' {
          // Line comment
          if err := this.skipLine(); err != nil {
            return nil, err
          }
        } else {
          if err := this.unread(); err != nil {
            return nil, err
          }
        }

      default:
        if isWhitespace(chr) {
          continue
        } else if isValidIdentifierStart(chr) {
          startPosition := this.currentPosition

          if err := this.unread(); err != nil {
            return nil, err
          }

          if literal, err := this.lexLiteral(); err != nil {
            return nil, err
          } else {
            kw := this.lexKeyWord(literal)
            if kw == TOKEN_ERROR {
              // Not a keyword
              return NewToken(startPosition, TOKEN_IDENTIFIER, literal)
            }

            return NewToken(startPosition, kw, strings.ToUpper(literal))
          }
        }
      }
    }
  }

  return &Token{this.currentPosition, TOKEN_EOF, ""}, nil
}

func (this *Lexer) read() (rune, error) {
  chr, _, err := this.input.ReadRune()

  if err != nil {
    if err == io.EOF {
      return eof, nil
    }

    return eof, err
  }

  this.currentPosition.Column++
  return chr, nil
}

func (this *Lexer) unread() error {
  err := this.input.UnreadRune()
  if err != nil {
    return err
  }

  this.currentPosition.Column--
  return nil
}

func (this *Lexer) newLine() {
  this.currentPosition.Line++
  this.currentPosition.Column = 0
}

func (this *Lexer) skipLine() error {
  for {
    chr, err := this.read()
    if err != nil {
      return err
    }

    if chr == '\n' {
      this.newLine()
      return nil
    }
  }
}

func (this *Lexer) lexLiteral() (string, error) {
  var buffer bytes.Buffer

  for {
    if chr, err := this.read(); err != nil {
      return "", err
    } else {
      if chr == eof {
        return buffer.String(), nil
      }

      if isValididentifierPart(chr) {
        buffer.WriteRune(chr)
      } else {
        if err := this.unread(); err != nil {
          return "", err
        }

        return buffer.String(), nil
      }
    }
  }
}

func (this *Lexer) lexKeyWord(literal string) TokenType {
  typ, ok := keywords[strings.ToUpper(literal)]
  if !ok {
    return TOKEN_ERROR
  }

  return typ
}
