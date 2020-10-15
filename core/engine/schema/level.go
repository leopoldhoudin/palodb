package schema

type Level struct {
  Name string
  Attributes []*Attribute
}

func NewLevel(name string, attrs []*Attribute) *Level {
  return &Level{name, attrs}
}
