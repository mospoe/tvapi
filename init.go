package tvapi

import (
  "fmt"
)

func Init() {
  config := NewConfig()
  fmt.Println(*config)
}
