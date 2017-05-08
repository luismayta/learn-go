package main

import "fmt"
import "time"

// test goroutine
func main() {
  var message = "hello"
  go func() {
    fmt.Println(message)
  }()
  message = "Goodbye"
  time.Sleep(100 * time.Millisecond)
}
