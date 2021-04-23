package goroutines

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

func getFib() {
	for x := range fib(10) {
		fmt.Println(x)
	}
}

func fibBest(c chan int, n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
		c <- a
	}
	close(c)
}

func getFibBest() {
	c := make(chan int)
	go fibBest(c, 10)

	for x := range c {
		fmt.Println(x)
	}
}
