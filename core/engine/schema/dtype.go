package schema

type DataType int

const (
  STRING DataType = iota
  INTEGER
)

var dtypesMap = map[string]DataType {
  "STRING": STRING,
  "INTEGER": INTEGER,
}

func GetDataTypeFromName(dtypeName string) DataType {
  return dtypesMap[dtypeName]
}
