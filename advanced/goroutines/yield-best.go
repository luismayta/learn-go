package main

import "fmt"

func fib(n int) chan int {
	C := make(chan int)

	go func() {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			a, b = b, a+b
			C <- a
		}
		close(C)
	}()
	return C
}

func main() {
	for x := range fib(10) {
		fmt.Println(x)
	}
}
