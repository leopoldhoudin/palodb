package schema

type Attribute struct {
  Name string
  Dtype DataType
}

func NewAttribute(name string, dtype DataType) *Attribute {
  return &Attribute{name, dtype}
}
