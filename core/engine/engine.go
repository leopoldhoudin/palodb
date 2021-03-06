package engine

import (
  "fmt"
  "strings"

  "github.com/leopoldhoudin/palodb/core/lang"
  "github.com/leopoldhoudin/palodb/core/engine/schema"
)

type Engine struct {
  config *Config
  schema *schema.Schema
}

func NewEngine(config *Config) *Engine {
  schema := schema.Load(config.DataPath)

  fmt.Println("=== Dimensions ===")
  for _, dim := range schema.Dims {
    fmt.Printf("- %s\n", dim.Name)
  }

  return &Engine{config, schema}
}

func (this *Engine) ExecuteString(text string) error {
  lexer := lang.NewLexer(strings.NewReader(text))
  parser := lang.NewParser(lexer)

  stmts, err := parser.Parse()
  if err != nil {
    return err
  }

  for _, stmt := range stmts {
    if err := this.ExecuteStatement(stmt); err != nil {
      return err
    }
  }

  return nil
}

func (this *Engine) ExecuteStatement(stmt lang.Statement) error {
  switch typedStmt := stmt.(type) {
  case *lang.CreateDimension:
    return this.executeCreateDimension(typedStmt)

  default:
    fmt.Printf("Unknown type: %s\n", typedStmt)
    break

  }

  return nil
}

func (this *Engine) executeCreateDimension(stmt *lang.CreateDimension) error {
  for _, dim := range this.schema.Dims {
    if dim.Name == stmt.Name {
      return &RuntimeError{fmt.Sprintf(
        "A dimension with the same name already exists ('%s')",
        dim.Name,
      )}
    }
  }

  var lvls = []*schema.Level{}
  for _, lvlStmt := range stmt.Levels {
    var attrs = []*schema.Attribute{}
    for _, attrStmt := range lvlStmt.Attributes {
      attrs = append(
        attrs,
        schema.NewAttribute(
          attrStmt.Name,
          schema.GetDataTypeFromName(attrStmt.Dtype),
        ),
      )
    }

    lvls = append(
      lvls,
      schema.NewLevel(
        lvlStmt.Name,
        attrs,
      ),
    )
  }

  dim := schema.NewDimension(stmt.Name, lvls)

  this.schema.Dims = append(this.schema.Dims, dim)

  schema.Save(this.schema, this.config.DataPath)

  return nil
}
