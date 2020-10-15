package lang

import "fmt"

type SyntaxError struct {
  pos Position
  msg string
}

func (this *SyntaxError) Error() string {
  return fmt.Sprintf("(%d:%d): %s", this.pos.Line, this.pos.Column, this.msg)
}
