package main

import (
	// "fmt"
	// "os"
  // "io"
  // "bufio"

	"flag"

	"github.com/leopoldhoudin/palodb/server"

	// "github.com/leopoldhoudin/palodb/core/lang"
	// "github.com/leopoldhoudin/palodb/core/engine"
)

// func getFile(fileName string) io.Reader {
//   file, err := os.Open(fileName)
//   if err != nil {
//     panic(err)
//   }
//
//   return file
// }
//
// func runFile(fileName string) {
//   fmt.Println("===== FILE ====================================================")
//
//   reader := bufio.NewReader(getFile(fileName))
//   for {
//     line, err := reader.ReadString('\n')
//     if err != nil {
//       if err == io.EOF {
//         break
//       }
//
//       panic(err)
//     }
//
//     fmt.Print(line)
//   }
// }
//
// func runLexer(fileName string) {
//   fmt.Println("===== LEXER ===================================================")
//
// 	lexer := lang.NewLexer(getFile(fileName))
//
// 	for {
// 		token, err := lexer.Next()
//
//     if err != nil {
//       fmt.Println(err)
//       break
//     }
//
// 		if token.Type == lang.TOKEN_EOF {
// 			break
// 		}
//
// 		fmt.Printf(
//       "%d:%d\t%-10v\t%s\n",
//       token.Position.Line,
//       token.Position.Column,
//       token.Type,
//       token.Literal,
//     )
// 	}
// }
//
// func runParser(fileName string) {
//   fmt.Println("===== PARSER ==================================================")
//
//   lexer := lang.NewLexer(getFile(fileName))
//   parser := lang.NewParser(lexer)
//
//   stmts, err := parser.Parse()
//   if err != nil {
//     fmt.Println()
//     panic(err)
//   }
//
//   fmt.Printf("Got %d statements\n", len(stmts))
//   for _, stmt := range stmts {
//     fmt.Println(stmt)
//   }
// }
//
// func runExecute(fileName string) {
//   fmt.Println("===== EXEC ====================================================")
//
//   // lexer := lang.NewLexer(getFile(fileName))
//   // parser := lang.NewParser(lexer)
// 	//
//   // stmts, err := parser.Parse()
//   // if err != nil {
//   //   fmt.Println()
//   //   panic(err)
//   // }
//
//   config := engine.NewConfig()
//   config.DataPath = "data/"
//
// 	engine.NewEngine(config)
//
//   // for _, stmt := range stmts {
//   //   if err := eng.Execute(stmt); err != nil {
// 	// 		panic(err)
// 	// 	}
//   // }
// }

func main() {
	var port int

	flag.IntVar(&port, "port", 5050, "--port 5050")

	flag.Parse()

	// fmt.Println("palodb -", Version)
  //
  // fileName := "test.palo"
	//
  // runFile(fileName)
  // runLexer(fileName)
  // runParser(fileName)
	//
  // // runExecute(fileName)

	server.Start(port)
}
