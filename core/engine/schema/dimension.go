package schema

type Dimension struct {
  Name string
  Levels []*Level
}

func NewDimension(name string, lvls []*Level) *Dimension {
  return &Dimension{name, lvls}
}
