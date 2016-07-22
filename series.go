package tvapi

import (
  "encoding/json"
  "fmt"
)

type Series struct {
  api int64
  key string
  source string
}

func NewSeries (source string, dbase *Dbase) *Series {
  var s Series
  s.source = source
  s.key = util_format(source, "series")
  s.Populate(dbase)
  return &s
}

func (s *Series) Populate (dbase *Dbase) {
  entry := dbase.Search(s.key)
  if entry.api > 0 {
    s.api = entry.api
    s.key = entry.key
  }
}

func (s *Series) Api (dbase *Dbase) {
  url := fmt.Sprintf("search/shows?q=%s", s.source)
  raw := util_curl(url)

  type Show struct {
    Id float64
    Name string
    Premiered string
  }

  type Shows struct {
    Show Show
  }

  var entries []Shows
  err := json.Unmarshal(raw, &entries)
  if err != nil {
    fmt.Println("invalid json")
    return
  }

  if len(entries) == 0 {
    fmt.Println("no results")
    return
  }

  fmt.Println("0 None")
  for i, entry := range entries {
    e := entry.Show
    eline := fmt.Sprintf("%d %s (%s)", (i + 1), e.Name, e.Premiered)
    fmt.Println(eline)
  }

  fmt.Print("Select Show: ")
  var pick int
  rd, err := fmt.Scan(&pick)
  if err != nil || rd == 0 {
    pick = 0
  }

  if pick <= 0 {
    return
  }

  if pick > len(entries) {
    fmt.Println("invalid selection")
    return
  }

  //fmt.Println(s, dbase)
  s.api = int64(entries[pick - 1].Show.Id)
  dbase.Store(s.api, s.key)
}
