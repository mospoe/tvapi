package tvapi

import (
  "fmt"
)

func Init() {
  config := NewConfig()
  if !config.ready {
    fmt.Println("config not ready")
    return
  }

  dbase := NewDbase(config.dbase)
  switch config.method {
  case "episode":
    fmt.Println("Method episode")
  break
  case "series":
    series := NewSeries(config.source[0], dbase)
    if series.api > 0 {
      fmt.Println("Found")
    } else {
      fmt.Println("No Joy :(")
    }
  break
  }
  dbase.Close()
}
