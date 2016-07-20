package tvapi

import (
  "fmt"
  "os"
  "strings"
)

type Hash struct {
  key string
  val string
  ext []string
}

func handle_err(err error, exit int, args ...interface{}) {
  if err != nil {
    fmt.Println(err.Error())
  }

  if len(args) > 0 {
    for _, msg := range args {
      fmt.Println("msg:", msg)
    }
  }

  if exit > 0 {
    os.Exit(exit)
  }
}

func read_hash(file string) []Hash {
  var hash []Hash

  o_file, err := os.Open(file)
  defer o_file.Close()
  if err != nil {
    handle_err(err, 1)
  }

  info, err := o_file.Stat()
  if err != nil {
    handle_err(err, 1)
  }

  data := make([]byte, info.Size())
  _, err = o_file.Read(data)
  if err != nil {
    handle_err(err, 1)
  }

  raw := strings.Split(string(data), "\n")

  for _, line := range raw {
    if len(line) == 0 {
      continue
    }

    if line[0] == '#' {
      continue
    }

    kv := strings.Split(line, " ")
    if len(kv) < 2 {
      continue
    }
    var h Hash
    h = Hash{kv[0], kv[1], kv[2:]}
    hash = append(hash, h)
  }

  return hash
}

func util_format (s string, delim string) string {
  var tmp string
  parts := make([]string, 0)
  source := strings.ToLower(s)
  for _, c := range source {
    if (util_isan(c)) {
      tmp += string(c)
    } else if (len(tmp) > 0) {
      parts = append(parts, tmp)
      tmp = ""
    }
  }

  if (len(tmp) > 0) {
    parts = append(parts, tmp)
  }

  r := strings.Join(parts, delim)
  return r
}

func util_isan(c rune) bool {
  r := false
  if (c >= rune('a') && c <= rune('z')) {
    r = true
  }

  if (c >= rune('0') && c <= rune('9')) {
    r = true
  }

  return r
}

func util_print (args ...interface{}) {
  for _, v := range args {
    fmt.Println(v)
  }
}
