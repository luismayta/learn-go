package variables

import "fmt"

func getSingleVar() string {
	name, location := "itachi", "konoha"
	return fmt.Sprintf("%s of %s", name, location)
}
