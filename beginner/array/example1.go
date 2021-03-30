package array

import (
	"fmt"
)

var users = [...]string{"Luis", "Jorge", "Oscar"}

func main() {
	fmt.Println(users)
	for _, value := range users {
		fmt.Println(value)
	}
}
