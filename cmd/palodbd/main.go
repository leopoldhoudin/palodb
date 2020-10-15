package main

import (
	"fmt"
	"os"
  "io"
  "bufio"

	"github.com/leopoldhoudin/palodb/core/lang"
)

func getFile(fileName string) io.Reader {
  file, err := os.Open(fileName)
  if err != nil {
    panic(err)
  }

  return file
}

func runFile(fileName string) {
  fmt.Println("===== FILE ====================================================")

  reader := bufio.NewReader(getFile(fileName))
  for {
    line, err := reader.ReadString('\n')
    if err != nil {
      if err == io.EOF {
        break
      }

      panic(err)
    }

    fmt.Print(line)
  }
}

func runLexer(fileName string) {
  fmt.Println("===== LEXER ===================================================")

	lexer := lang.NewLexer(getFile(fileName))

	for {
		token, err := lexer.Next()

    if err != nil {
      fmt.Println(err)
      break
    }

		if token.Type == lang.TOKEN_EOF {
			break
		}

		fmt.Printf(
      "%d:%d\t%-10v\t%s\n",
      token.Position.Line,
      token.Position.Column,
      token.Type,
      token.Literal,
    )
	}
}

func runParser(fileName string) {
  fmt.Println("===== PARSER ==================================================")

  lexer := lang.NewLexer(getFile(fileName))
  parser := lang.NewParser(lexer)

  stmts, err := parser.Parse()
  if err != nil {
    fmt.Println()
    panic(err)
  }

  fmt.Printf("Got %d statements\n", len(stmts))
  for _, stmt := range stmts {
    fmt.Println(stmt)
  }
}

func main() {
	fmt.Println("palodb -", Version)

  fileName := "test.palo"

  runFile(fileName)
  runLexer(fileName)
  runParser(fileName)
}
