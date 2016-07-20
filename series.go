package tvapi

type Series struct {
  api int64
  key string
  source string
}

func NewSeries (source string, dbase *Dbase) *Series {
  var s Series
  s.source = source
  s.key = util_format(source, "-")
  s.Populate(dbase)
  return &s
}

func (s *Series) Populate (dbase *Dbase) {
  entry := dbase.Search(s.key)
  util_print(*entry)
  if entry.api > 0 {
    s.api = entry.api
    s.key = entry.key
  }
}

func (s *Series) Api (dbase *Dbase) {
  // look up api
}
