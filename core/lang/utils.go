package lang

func isWhitespace(chr rune) bool {
  return chr == ' ' || chr == '\t' || chr == '\n' || chr == '\r'
}

func isDigit(chr rune) bool {
  return chr >= '0' && chr <= '9'
}

func isLetter(chr rune) bool {
  return (chr >= 'a' && chr <= 'z') || (chr >= 'A' && chr <= 'Z')
}

func isAlphanum(chr rune) bool {
  return isDigit(chr) || isLetter(chr)
}

func isValidIdentifierStart(chr rune) bool {
  return isLetter(chr) || chr == '_'
}

func isValididentifierPart(chr rune) bool {
  return isAlphanum(chr) || chr == '_'
}
