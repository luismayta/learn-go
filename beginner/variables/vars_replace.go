//-*- coding: utf-8 -*-
package main

import (
	"fmt"
)

var (
	name = "itachi uchiha"
)

func main() {
	message := fmt.Sprintf("Hello %s", name)
	fmt.Println(message)
}
