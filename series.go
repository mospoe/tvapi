package tvapi

type Series struct {
  api int64
  dbase string
}

func NewSeries (source string, dbase *Dbase) *Series {
  var series Series
  return &series
}
