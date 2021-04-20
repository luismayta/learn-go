package main

import (
	"fmt"
	"time"
)

// test goroutine.
func main() {
	message := "hello"
	go func() {
		fmt.Println(message)
	}()
	message = "Goodbye"
	time.Sleep(100 * time.Millisecond)
}
