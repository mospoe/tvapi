package tvapi

import (
  "os"
  "strings"
)

type Config struct {
  ready bool

  method string // [episode, series]
  mode string // [copy, move]

  dbase string

  source []string
}

func NewConfig () *Config {
  var c Config
  c.ready = false
  c.ProcessArgs()
  c.UserConfig()
  return &c
}

func (c *Config) UserConfig () {
  c.dbase = "/home/mos/.local/tvapi.db"
}

func (c *Config) ProcessArgs () {
  c.mode = "move"
  c.method = "episode"

  for _, arg := range os.Args[1:] {
    if arg[0] == '-' {
      switch arg[1] {
      case 'c':
        c.mode = "copy"
      break

      case 'h':
        c.mode = "help"
      break
      }
    } else {
      if len(c.source) > 0 {
        continue
      }
      test, _ := os.Open(arg)
      if test == nil {
        c.method = "series"
      } else {
        test.Close()
      }
      c.source = append(c.source, arg)
    }
  }

  if len(c.source) == 0 && c.method == "episode" {
    c.ScanSource()
  }
}

func (c *Config) ScanSource () {
  extensions := []string{"avi", "mkv", "mp4"}

  pwd, _ := os.Getwd()
  o_dir, _ := os.Open(pwd)
  o_files, _ := o_dir.Readdir(0)
  o_dir.Close()

  for _, o_file := range o_files {
    if (o_file.IsDir()) {
      continue
    }

    name := o_file.Name()
    tmp := strings.Split(name, ".")
    extension := tmp[len(tmp) - 1]
    extension = strings.ToLower(extension)

    for _, e := range extensions {
      if e == extension {
        c.source = append(c.source, name)
        break
      }
    }
  }
}
