package tvapi

import (
  "fmt"
  "os"
  "strconv"
)

type DbEntry struct {
  api int64
  key string
  alias []string
}

type Dbase struct {
  file *os.File
  series []DbEntry
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
  hash := util_hash(file)
  for _, e := range hash {
    api, _ := strconv.Atoi(e.val)
    ds := DbEntry{int64(api), e.key, e.ext}
    ds.alias = append(ds.alias, ds.key)
    d.series = append(d.series, ds)
  }
}

func (d *Dbase) Search (k string) *DbEntry {
  var e DbEntry
  for _, entry := range d.series {
    for _, alias := range entry.alias {
      if alias == k {
        e = entry
        return &e
      }
    }
  }
  return &e
}

func (d *Dbase) Store (api int64, key string) {
  e := DbEntry{api, key, []string{key}}
  d.series = append(d.series, e)
  line := fmt.Sprintf("%s %d\n", key, api)
  info, err := d.file.Stat()
  if err != nil {
    fmt.Println("stat error " + key)
    return
  }

  store := []byte(line)
  wr, err := d.file.WriteAt(store, info.Size())
  if err != nil || wr == 0 {
    fmt.Println("write error" + key)
    return
  }
}

func (d *Dbase) Close () {
  d.file.Close()
}
