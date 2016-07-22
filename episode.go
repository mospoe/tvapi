package tvapi

import (
//  "encoding/json"
//  "fmt"
  "regexp"
  "strings"
)

type Episode struct {
  source string
  title string
  path string
  series *Series
  split []string
  season string
  episode string
}

func NewEpisode (source string, dbase *Dbase) *Episode {
  var e Episode
  e.source = source
  series_index := e.digest()
  if series_index > 0 {
    tmp := strings.Join(e.split[0:series_index], "-")
    e.series = NewSeries(tmp, dbase)
    if e.series.api > 0 {
      util_print(e.series.api)
    }
  }
  return &e
}

func (e *Episode) digest () int {
  tmp := util_format(e.source, "episode")
  e.split = strings.Split(tmp, ".")

  for i, p := range e.split {
    if se, t := e.is_se(p); se {
      e.get_se(p, t)
      return i
      break
    }
  }

  return 0
}

func (e *Episode) get_se (se string, t int) {
  switch t {
  case 1:
    if len(se) == 3 {
      se = "0" + se
    }

    e.season = se[0:2]
    e.episode = se[2:4]
  break

  case 2:
    e.season = se[1:3]
    e.episode = se[4:6]
  break
  }
}

func (e *Episode) is_se (p string) (bool, int) {
  var r string
  var m bool

  r = "[0-9]{3,4}"
  if m, _ = regexp.MatchString(r, p); m {
    if len(p) == 3 {
      p = "0" + p
    }
    return true, 1
  }

  r = "s[0-9]{2}e[0-9]{2}"
  if m, _ = regexp.MatchString(r, p); m {
    return true, 2
  }

  return false, 0
}
