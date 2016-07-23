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

  if len(config.source) == 0 {
    fmt.Println("nothing to do :(")
    return
  }

  dbase := NewDbase(config.dbase)
  switch config.method {
  case "episode":
    for _, s := range config.source {
      episode := NewEpisode(s, dbase, config)
      if episode.title == "" && !config.quiet {
        fmt.Println(s, "not found")
      }
    }
  break
  case "series":
    series := NewSeries(config.source[0], dbase, config)

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
