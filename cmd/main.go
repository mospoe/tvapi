package main

import (
  "fmt"
  "mospoe.com/tvapi"
)

func main () {
  fmt.Println("main")
  config := tvapi.NewConfig()
  fmt.Println(*config)
}
