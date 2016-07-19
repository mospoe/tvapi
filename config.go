package tvapi

type Config struct {
  dbase string
}

func NewConfig () *Config {
  var c Config

  return &c
}
