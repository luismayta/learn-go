package arrays

import (
	"fmt"

	"./slices"
)

var users = [...]string{"Luis", "Chris", "Oscar"}

func main() {
	fmt.Println(users)
	for _, value := range users {
		fmt.Println(value)
	}
	slices.P
}
