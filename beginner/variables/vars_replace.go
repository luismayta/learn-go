package variables

import (
	"fmt"
)

func getMessage() string {
	name := "itachi uchiha"

	message := fmt.Sprintf("Hello %s", name)

	return message
}
