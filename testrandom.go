package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  for i := 0; i < 25; i++ {
    fmt.Println(rand.Intn(12-1) + 1)
  }
}
