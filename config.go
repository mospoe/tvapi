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
  video string

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
  conf := "$HOME/.config/tvapi.conf"
  conf = os.ExpandEnv(conf)

  o_conf, err := os.Open(conf)
  defer o_conf.Close()
  if err != nil {
    handle_err(err, 1)
  }
  opts := read_hash(conf)
  for _, hash := range opts {
    hash.val = os.ExpandEnv(hash.val)

    switch hash.key {
    case "dbase":
      c.dbase = hash.val
    break
    case "video":
      c.video = hash.val
    break
    }
  }

  if c.dbase != "" && c.video != "" {
    c.ready = true
  }
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
