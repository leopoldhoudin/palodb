package lang

type Token struct {
  Position Position
  Type TokenType
  Literal string
}

type Position struct {
  Line int
  Column int
}

type TokenType string

const (
  TOKEN_ERROR TokenType = "ERROR"
  TOKEN_EOF             = "EOF"

  TOKEN_IDENTIFIER      = "IDENT"
  TOKEN_SEMI            = "SEMI"
  TOKEN_COMA            = "COMA"
  TOKEN_PAR_OPEN        = "PAR_OPEN"
  TOKEN_PAR_CLOSE       = "PAR_CLOSE"

  TOKEN_STRING          = "CST_STRING"
  TOKEN_INTEGER         = "CST_INTEGER"
  TOKEN_FLOAT           = "CST_FLOAT"

  TOKEN_KW_STRING       = "KW_STRING"
  TOKEN_KW_INTEGER      = "KW_INTEGER"
  TOKEN_KW_CREATE       = "KW_CREATE"
  TOKEN_KW_DIMENSION    = "KW_DIM"
  TOKEN_KW_DEFAULT      = "KW_DEFAULT"
)

var keywords = map[string]TokenType{
  "STRING": TOKEN_KW_STRING,
  "INTEGER": TOKEN_KW_INTEGER,
  "CREATE": TOKEN_KW_CREATE,
  "DIMENSION": TOKEN_KW_DIMENSION,
  "DEFAULT": TOKEN_KW_DEFAULT,
}

func NewToken(position Position, typ TokenType, literal string) (*Token, error) {
  return &Token{position, typ, literal}, nil
}
