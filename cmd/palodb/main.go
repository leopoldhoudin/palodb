package main

import (
  "flag"

  "github.com/leopoldhoudin/palodb/client"
)

func main() {
  var port int

  flag.IntVar(&port, "port", 5050, "--port 5050")

  flag.Parse()

  client.Start(port)
}
