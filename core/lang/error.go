package lang

import "fmt"

type SyntaxError struct {
  pos Position
  msg string
}

func (this *SyntaxError) Error() string {
  return fmt.Sprintf("[SyntaxError@(%d:%d)]: %s", this.pos.Line, this.pos.Column, this.msg)
}
