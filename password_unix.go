// +build linux darwin

package passwdprompt

import (
	"code.google.com/p/gopass"
)

// GetPassword asks the user to enter her password.
func GetPassword(prompt string) (pw string, err error) {
	return gopass.GetPass(prompt)
}
