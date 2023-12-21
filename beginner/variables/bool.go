package variables

import (
	"fmt"
)

var casiQueSi bool
var casiQueNo bool

func init() {
	casiQueSi = true
	casiQueNo = false
	fmt.Println("casiQueSi:", casiQueSi)
	fmt.Println("casiQueNo:", casiQueNo)
	fmt.Println("Sí  Y  No:", casiQueSi && casiQueNo)
	fmt.Println("Sí  O  No:", casiQueSi || casiQueNo)
	fmt.Println("NO(Sí)   :", !casiQueSi)
}
