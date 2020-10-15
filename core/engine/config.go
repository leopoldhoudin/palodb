package engine

type Config struct {
  DataPath string
}

func NewConfig() *Config {
  return &Config{""}
}
