package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

/*
 *
 * Create Dimension product (
 *   category String
 *   name String
 * ) ;
 *
 * Create Dimension time (
 *   year Integer
 *   date Date
 * ) ;
 *
 * Create Facts sales (
 *   amount Integer
 * ) With (
 *   time
 *   product
 * );
 *
 * Describe sales ;
 * Select (Sum) amount From sales By product.category Where time.year = 2020 ;
 */

func main() {
  fmt.Println("palodb - 0.0.0")
  fmt.Println("--------------")

  isRunning := true

  reader := bufio.NewReader(os.Stdin)

  for isRunning {
    fmt.Print(">>> ")
    text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)

    if strings.Compare(text, "quit") == 0 {
      isRunning = false
      fmt.Println("Bye bye :)")
    } else {
      fmt.Println(text)
    }
  }
}
