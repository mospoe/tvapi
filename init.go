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
    fmt.Println("method episode")
  break
  case "series":
    series := NewSeries(config.source[0], dbase)
    if series.api == 0 && !config.quiet {
      fmt.Println("search api for", series.key)
      series.Api(dbase)
    }

    if !config.quiet {
      if series.api > 0 {
        fmt.Println("found", series.key, "with api id", series.api)
      } else {
        fmt.Println(series.key, "not found")
      }
    }
  break
  }
  dbase.Close()
}
