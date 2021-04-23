package variables

import (
	"fmt"
)

func getInitializerVars() string {
	var (
		name     = "Luis Mayta"
		age      = 30
		location = "Lima - Peru"
		alias    = "Itachi uchiha"
	)

	return fmt.Sprintf("%s have %d from %s alias %s",
		name,
		age,
		location,
		alias,
	)
}
