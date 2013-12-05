package main

import (
	"fmt"
	"github.com/tajtiattila/passwdprompt"
)

func main() {
	fmt.Println("Password demo")
	u, p, err := passwdprompt.GetUserPassword("Username: ", "Password: ")
	if err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("Got username:", u)
		fmt.Println("Got password:", p)
	}
}
