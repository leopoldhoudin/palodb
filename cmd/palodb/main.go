package main

import (
  "flag"

  "github.com/leopoldhoudin/palodb/client"
)

func main() {
  var port int
  var host string

  flag.IntVar(&port, "port", 5050, "--port 5050")
  flag.StringVar(&host, "host", "127.0.0.1", "--host 127.0.0.1")

  flag.Parse()

  client.Start(host, port)
}
