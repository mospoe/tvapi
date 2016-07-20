package tvapi

import (
  "os"
  "strconv"
)

type DbSeries struct {
  api int64
  key string
  alias []string
}

type Dbase struct {
  file *os.File
  series []DbSeries
}

func NewDbase (file string) *Dbase {
  var d Dbase
  o_file, err := os.OpenFile(file, os.O_RDWR, os.ModeAppend)
  if err != nil && os.IsNotExist(err) {
    o_file, err = os.Create(file)
  }

  if err != nil {
    handle_err(err, 1)
  }

  d.file = o_file
  d.load(file)
  return &d
}

func (d *Dbase) load (file string) {
  hash := read_hash(file)
  for _, e := range hash {
    api, _ := strconv.Atoi(e.val)
    ds := DbSeries{int64(api), e.key, e.ext}
    ds.alias = append(ds.alias, ds.key)
    d.series = append(d.series, ds)
  }
}

func (d *Dbase) Close () {
  d.file.Close()
}
