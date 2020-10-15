package lang

import "fmt"

type Callback func(Statement, *Token) (Statement, error)

func dummy(stmt Statement, token *Token) (Statement, error) {
  return stmt, nil
}

func createDimension(stmt Statement, token *Token) (Statement, error) {
  fmt.Printf(">>> createDimension(%s)\n", token.Literal)
  return &CreateDimension{token.Literal, []*CreateLevel{}}, nil
}

func createFirstLevel(stmt Statement, token *Token) (Statement, error) {
  fmt.Printf(">>> createFirstLevel(%s)\n", token.Literal)
  dim := stmt.(*CreateDimension)
  lvl := &CreateLevel{dim, token.Literal, []*CreateAttribute{}}
  dim.Levels = append(dim.Levels, lvl)
  return lvl, nil
}

func createNextLevel(stmt Statement, token *Token) (Statement, error) {
  fmt.Printf(">>> createNextLevel(%s)\n", token.Literal)
  oth := stmt.(*CreateLevel)
  lvl := &CreateLevel{oth.Dim, token.Literal, []*CreateAttribute{}}
  oth.Dim.Levels = append(oth.Dim.Levels, lvl)
  return lvl, nil
}

func createDimensionLevelsEnd(stmt Statement, token *Token) (Statement, error) {
  fmt.Printf(">>> createDimensionLevelsEnd\n")
  lvl := stmt.(*CreateLevel)
  return lvl.Dim, nil
}

func createFirstAttribute(stmt Statement, token *Token) (Statement, error) {
  fmt.Printf(">>> createFirstAttribute(%s)\n", token.Literal)
  lvl := stmt.(*CreateLevel)
  att := &CreateAttribute{lvl, token.Literal, "", false}
  lvl.Attributes = append(lvl.Attributes, att)
  return att, nil
}

func createNextAttribute(stmt Statement, token *Token) (Statement, error) {
  fmt.Printf(">>> createNextAttribute(%s)\n", token.Literal)
  oth := stmt.(*CreateAttribute)
  att := &CreateAttribute{oth.Lvl, token.Literal, "", false}
  oth.Lvl.Attributes = append(oth.Lvl.Attributes, att)
  return att, nil
}

func createAttributeSetType(stmt Statement, token *Token) (Statement, error) {
  fmt.Printf(">>> createAttributeSetType(%s)\n", token.Literal)
  att := stmt.(*CreateAttribute)
  att.Type = token.Literal
  return att, nil
}

func createAttributeSetDefault(stmt Statement, token *Token) (Statement, error) {
  fmt.Printf(">>> createAttributeSetDefault\n")
  att := stmt.(*CreateAttribute)
  att.Default = true
  return att, nil
}

func createLevelAttributesEnd(stmt Statement, token *Token) (Statement, error) {
  fmt.Printf(">>> createLevelAttributesEnd\n")
  att := stmt.(*CreateAttribute)
  return att.Lvl, nil
}
