package client

import (
  "os"
  "fmt"
  "net"
  "time"
  "sync"
  "bufio"
  "strings"
  "encoding/binary"

  "github.com/smallnest/goframe"

  "github.com/leopoldhoudin/palodb/protocol"
)

type Client struct {
  host string
  port int
  lock sync.Mutex
  conn goframe.FrameConn
  resetByRemote bool
}

func NewClient(host string, port int) *Client {
  var client Client
  client.host = host
  client.port = port
  return &client
}

func (this *Client) reset() {
  this.conn = createConn(this.host, this.port)
  this.resetByRemote = false
  go this.check()
}

func (this *Client) close() {
  this.conn.Close()
}

func (this *Client) check() {
  for {
    this.lock.Lock()
    if err := this.conn.WriteFrame(protocol.NewPingFrame()); err != nil {
      fmt.Println("Connection reset by remote")
      this.resetByRemote = true
      this.lock.Unlock()
      return
    }

    buf, err := this.conn.ReadFrame()
    if err != nil {
      fmt.Println("Connection reset by remote")
      this.resetByRemote = true
      this.lock.Unlock()
      return
    }

    this.resetByRemote = false
    this.lock.Unlock()

    if len(buf) > 1 || buf[0] != protocol.PONG {
      fmt.Println("Error while communicating with remote")
      return
    }

    time.Sleep(time.Second)
  }
}

func Start(host string, port int) {
  fmt.Printf("palodb (%s)\n", Version)

  client := NewClient(host, port)
  client.reset()

  fmt.Println(client.resetByRemote)

  defer client.conn.Close()

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
      if client.resetByRemote {
        client.reset()
      }

      cmd = strings.TrimRight(cmd, "\n\t ")
      if !strings.HasSuffix(cmd, ";") {
        cmd = fmt.Sprintf("%s;", cmd)
      }

      client.lock.Lock()

      fmt.Println("Sending:")
      fmt.Println(cmd)
      client.conn.WriteFrame(protocol.NewStmtFrame(cmd))

      buf, _ := client.conn.ReadFrame()

      client.lock.Unlock()

      fmt.Println(string(buf))
    }
  }
}

func createConn(host string, port int) goframe.FrameConn {
  tcpConn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
  if err != nil {
    fmt.Println("Could not connect to remote")
    os.Exit(1)
  }

  frameConn := goframe.NewLengthFieldBasedFrameConn(
    goframe.EncoderConfig{
  		ByteOrder:                       binary.BigEndian,
  		LengthFieldLength:               4,
  		LengthAdjustment:                0,
  		LengthIncludesLengthFieldLength: false,
  	},
    goframe.DecoderConfig{
  		ByteOrder:           binary.BigEndian,
  		LengthFieldOffset:   0,
  		LengthFieldLength:   4,
  		LengthAdjustment:    0,
  		InitialBytesToStrip: 4,
  	},
    tcpConn,
  )

  return frameConn
}
