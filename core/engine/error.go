package engine

import "fmt"

type RuntimeError struct {
  msg string
}

func (this *RuntimeError) Error() string {
  return fmt.Sprintf("[RuntimeError]: %s", this.msg)
}
