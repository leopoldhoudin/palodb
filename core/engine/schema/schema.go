package schema

import (
  "os"
  "path"
  "encoding/gob"
)

const FILENAME = "palo_schema"

type Schema struct {
  Dims []*Dimension
}

func newEmptySchema() *Schema {
  return &Schema{
    []*Dimension{},
  }
}

func Load(dataPath string) *Schema {
  file, err := os.Open(path.Join(dataPath, FILENAME))
  if err != nil {
    if os.IsNotExist(err) {
      return newEmptySchema()
    }

    panic(err)
  }

  defer file.Close()

  schema := &Schema{}
  if err := gob.NewDecoder(file).Decode(schema); err != nil {
    panic(err)
  }

  return schema
}

func Save(schema *Schema, dataPath string) {
  file, err := os.OpenFile(
    path.Join(dataPath, FILENAME),
    os.O_CREATE | os.O_WRONLY,
    0644,
  )
  if err != nil {
    panic(err)
  }

  defer file.Close()

  if err := gob.NewEncoder(file).Encode(schema); err != nil {
    panic(err)
  }
}
