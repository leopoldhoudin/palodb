package server

import (
  "fmt"
  "time"
  "encoding/binary"

  "github.com/panjf2000/gnet"

  "github.com/leopoldhoudin/palodb/protocol"
  "github.com/leopoldhoudin/palodb/core/engine"
)

type Server struct {
  *gnet.EventServer
  lastFrameTsByConn map[gnet.Conn]time.Time

  engine *engine.Engine
}

func (this *Server) OnInitComplete(server gnet.Server) (action gnet.Action) {
  fmt.Printf(
    "listening on %s (multi-cores: %t, loops: %d)\n",
		server.Addr.String(),
    server.Multicore,
    server.NumEventLoop,
  )
  return
}

func (this *Server) OnOpened(conn gnet.Conn) (out []byte, action gnet.Action) {
  fmt.Printf("opened connection from %s\n", conn.RemoteAddr().String())

  this.lastFrameTsByConn[conn] = time.Now()

  return
}

func (this *Server) OnClosed(conn gnet.Conn, err error) (action gnet.Action) {
  fmt.Printf("closed connection from %s\n", conn.RemoteAddr().String())

  delete(this.lastFrameTsByConn, conn)

  return
}

func (this *Server) Tick() (delay time.Duration, action gnet.Action) {
  now := time.Now()
  for conn, lastFrameTs := range this.lastFrameTsByConn {
    if now.Sub(lastFrameTs) > 60 * time.Second {
      conn.Close()
    }
  }

  delay = 1 * time.Second
  return
}

func (this *Server) React(frame []byte, conn gnet.Conn) (out []byte, action gnet.Action) {
  // Echo synchronously.

  if len(frame) == 1 && frame[0] == protocol.PING {
    out = protocol.NewPongFrame()
    return
  }

  this.lastFrameTsByConn[conn] = time.Now()

  // === begin: handle statement frame ===

  if len(frame) > 1 && frame[0] == protocol.STMT {
    text := string(frame[1:])
    fmt.Printf("received statement: '%s'\n", text)

    err := this.engine.ExecuteString(text)
    if err != nil {
      out = []byte(err.Error())
      return
    }

    out = []byte("OK")
    return
  }

  // === end: handle statement frame ===

	out = frame
  return

	/*
		// Echo asynchronously.
		data := append([]byte{}, frame...)
		go func() {
			time.Sleep(time.Second)
			c.AsyncWrite(data)
		}()
		return
	*/
}

func Start(port int) {
  fmt.Printf("palodbd (%s)\n", Version)
  fmt.Printf("starting up...\n")

  config := engine.NewConfig()
  config.DataPath = "data/"

  engine := engine.NewEngine(config)

  server := &Server{
    lastFrameTsByConn: map[gnet.Conn]time.Time{},
    engine: engine,
  }
  gnet.Serve(
    server,
    fmt.Sprintf("tcp://:%d", port),
    gnet.WithMulticore(true),
    gnet.WithTicker(true),
    gnet.WithCodec(createCodec()),
  )
}

func createCodec() gnet.ICodec {
  return gnet.NewLengthFieldBasedFrameCodec(
    gnet.EncoderConfig{
			ByteOrder:                       binary.BigEndian,
			LengthFieldLength:               4,
			LengthAdjustment:                0,
			LengthIncludesLengthFieldLength: false,
		},
    gnet.DecoderConfig{
			ByteOrder:           binary.BigEndian,
			LengthFieldOffset:   0,
			LengthFieldLength:   4,
			LengthAdjustment:    0,
			InitialBytesToStrip: 4,
		},
  )
}
