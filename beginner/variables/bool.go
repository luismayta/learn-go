//-*- coding: utf-8 -*-
package bevariables

import "fmt"

func main() {
	var casiQueSi = true
	var casiQueNo bool
	fmt.Println("casiQueSi:", casiQueSi)
	fmt.Println("casiQueNo:", casiQueNo)
	fmt.Println("Sí  Y  No:", casiQueSi && casiQueNo)
	fmt.Println("Sí  O  No:", casiQueSi || casiQueNo)
	fmt.Println("NO(Sí)   :", !casiQueSi)
}
