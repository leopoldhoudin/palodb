package server

import (
  "fmt"

  "github.com/panjf2000/gnet"
)

type server struct {
  *gnet.EventServer
}

func (this *server) OnInitComplete(srv gnet.Server) (action gnet.Action) {
  fmt.Printf(
    "listening on %s (multi-cores: %t, loops: %d)\n",
		srv.Addr.String(),
    srv.Multicore,
    srv.NumEventLoop,
  )
  return
}

func (this *server) OnOpened(conn gnet.Conn) (out []byte, action gnet.Action) {
  fmt.Printf("opened connection from %s\n", conn.RemoteAddr().String())
  return
}

func (this *server) OnClosed(conn gnet.Conn, err error) (action gnet.Action) {
  fmt.Printf("closed connection from %s\n", conn.RemoteAddr().String())
  return
}

func (this *server) React(frame []byte, conn gnet.Conn) (out []byte, action gnet.Action) {
  // Echo synchronously.
	out = frame

  fmt.Printf("received %s\n", string(frame))

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

  srv := new(server)
  gnet.Serve(srv, fmt.Sprintf("tcp://:%d", port), gnet.WithMulticore(true))
}
