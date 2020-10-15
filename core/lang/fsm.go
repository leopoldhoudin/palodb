package lang

type State string

type Transition struct {
  // Implicit Source from FSM's key
  NextState State
  TokenType TokenType
  Callback Callback
}

const (
  STATE_STMT_START State      = "STMT_START"
  STATE_CREATE_START          = "CREATE_START"
  STATE_CREATE_DIM_START      = "CREATE_DIM_START"
  STATE_CREATE_DIM_IDENTIFIER = "CREATE_DIM_IDENT"
  STATE_LEVELS_START          = "LEVELS_START"
  STATE_LEVEL_IDENTIFIER      = "LEVEL_IDENTIFIER"
  STATE_ATTRIBUTES_START      = "ATTRIBUTES_START"
  STATE_ATTRIBUTE_IDENTIFIER  = "ATTRIBUTE_IDENTIFIER"
  STATE_ATTRIBUTE_TYPE        = "ATTRIBUTE_TYPE"
  STATE_ATTRIBUTE_DEFAULT     = "ATTRIBUTE_DEFAULT"
  STATE_ATTRIBUTES_NEXT       = "ATTRIBUTES_NEXT"
  STATE_ATTRIBUTES_END        = "ATTRIBUTES_END"
  STATE_LEVELS_NEXT           = "LEVELS_NEXT"
  STATE_LEVELS_END            = "LEVELS_END"
)

var FSM = map[State]([]Transition){
  STATE_STMT_START: []Transition{
    {STATE_CREATE_START, TOKEN_KW_CREATE, passthrough},
  },

  STATE_CREATE_START: []Transition{
    {STATE_CREATE_DIM_START, TOKEN_KW_DIMENSION, passthrough},
  },

  STATE_CREATE_DIM_START: []Transition{
    {STATE_CREATE_DIM_IDENTIFIER, TOKEN_IDENTIFIER, createDimension},
  },

  STATE_CREATE_DIM_IDENTIFIER: []Transition{
    {STATE_LEVELS_START, TOKEN_PAR_OPEN, passthrough},
  },

  STATE_LEVELS_START: []Transition{
    {STATE_LEVEL_IDENTIFIER, TOKEN_IDENTIFIER, createFirstLevel},
  },

  STATE_LEVEL_IDENTIFIER: []Transition{
    {STATE_ATTRIBUTES_START, TOKEN_PAR_OPEN, passthrough},
  },

  STATE_ATTRIBUTES_START: []Transition{
    {STATE_ATTRIBUTE_IDENTIFIER, TOKEN_IDENTIFIER, createFirstAttribute},
  },

  STATE_ATTRIBUTE_IDENTIFIER: []Transition{
    {STATE_ATTRIBUTE_TYPE, TOKEN_KW_STRING, createAttributeSetType},
    {STATE_ATTRIBUTE_TYPE, TOKEN_KW_INTEGER, createAttributeSetType},
  },

  STATE_ATTRIBUTE_TYPE: []Transition{
    {STATE_ATTRIBUTE_DEFAULT, TOKEN_KW_DEFAULT, createAttributeSetDefault},
    {STATE_ATTRIBUTES_NEXT, TOKEN_COMA, passthrough},
    {STATE_ATTRIBUTES_END, TOKEN_PAR_CLOSE, passthrough},
  },

  STATE_ATTRIBUTE_DEFAULT: []Transition{
    {STATE_ATTRIBUTES_NEXT, TOKEN_COMA, passthrough},
    {STATE_ATTRIBUTES_END, TOKEN_PAR_CLOSE, createLevelAttributesEnd},
  },

  STATE_ATTRIBUTES_NEXT: []Transition{
    {STATE_ATTRIBUTE_IDENTIFIER, TOKEN_IDENTIFIER, createNextAttribute},
    {STATE_ATTRIBUTES_END, TOKEN_PAR_CLOSE, createLevelAttributesEnd},
  },

  STATE_ATTRIBUTES_END: []Transition{
    {STATE_LEVELS_NEXT, TOKEN_COMA, passthrough},
    {STATE_LEVELS_END, TOKEN_PAR_CLOSE, createDimensionLevelsEnd},
  },

  STATE_LEVELS_NEXT: []Transition{
    {STATE_LEVEL_IDENTIFIER, TOKEN_IDENTIFIER, createNextLevel},
    {STATE_LEVELS_END, TOKEN_PAR_CLOSE, createDimensionLevelsEnd},
  },

  STATE_LEVELS_END: []Transition{
    {STATE_STMT_START, TOKEN_SEMI, passthrough},
  },
}
