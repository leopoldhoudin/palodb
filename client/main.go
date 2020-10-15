package client

import (
  "os"
  "fmt"
  "net"
  // "time"
  "bufio"
)

func Start(port int) {
  fmt.Printf("palodb (%s)\n", Version)

  conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", port))
  if err != nil {
    panic(err)
  }

  defer conn.Close()

  input := bufio.NewReader(os.Stdin)
  for {
    fmt.Print(">>> ")
    cmd, err := input.ReadString('\n')
    if err != nil {
      fmt.Println("Ooops something went wrong!")
      fmt.Println(err)
      return
    }
    cmd = cmd[:len(cmd) - 1]

    if cmd == "exit" {
      fmt.Println("Bye bye :)")
      return
    } else {
      n, _ := conn.Write([]byte(cmd))
      fmt.Println(n)

      conn.SetReadDeadline(time.Now().Add(2 * time.Second))

      var b = make([]byte, 1024)
      m, err := conn.Read(b[:])
      if err != nil {
        panic(err)
      }

      fmt.Println(m)
      fmt.Println(string(b))
    }
  }
}
