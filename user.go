package passwdprompt

import (
	"fmt"
)

// GetUserPassword asks the user to enter her username and password.
func GetUserPassword(uprompt, pprompt string) (u, pw string, err error) {
	fmt.Print(uprompt)
	_, err = fmt.Scanln(&u)
	if err == nil {
		pw, err = GetPassword(pprompt)
	}
	return
}
