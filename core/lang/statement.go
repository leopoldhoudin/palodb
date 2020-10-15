package lang

import (
  "fmt"
  "strings"
)

type Statement interface {
}

type CreateDimension struct {
  Name string
  Levels []*CreateLevel
}

type CreateLevel struct {
  Dim *CreateDimension
  Name string
  Attributes []*CreateAttribute
}

type CreateAttribute struct {
  Lvl *CreateLevel
  Name string
  Type string
  Default bool
}

func (this *CreateDimension) String() string {
  var lvls []string
  for _, lvl := range this.Levels {
    lvls = append(lvls, fmt.Sprintf("   %s", lvl.String()))
  }

  return fmt.Sprintf(
    "CreateDimension(%s):\n%s",
    this.Name,
    strings.Join(lvls, "\n"),
  )
}

func (this *CreateLevel) String() string {
  var atts []string
  for _, att := range this.Attributes {
    atts = append(atts, fmt.Sprintf("      %s", att.String()))
  }

  return fmt.Sprintf(
    "CreateLevel(%s):\n%s",
    this.Name,
    strings.Join(atts, "\n"),
  )
}

func (this *CreateAttribute) String() string {
  def := ""
  if this.Default {
    def = " DEFAULT"
  }

  return fmt.Sprintf("CreateAttribute(%s, %s%s)", this.Name, this.Type, def)
}
