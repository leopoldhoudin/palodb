package protocol

type FrameType byte

const (
  PING = 0x00
  PONG = 0x01
  STMT = 0x02
  DATA = 0x03
)

func NewPingFrame() []byte {
  return []byte{PING}
}

func NewPongFrame() []byte {
  return []byte{PONG}
}

func NewStmtFrame(stmt string) []byte {
  return append([]byte{STMT}, []byte(stmt)...)
}
