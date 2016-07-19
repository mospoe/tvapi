package tvapi

import (
  "fmt"
  "os"
)

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
