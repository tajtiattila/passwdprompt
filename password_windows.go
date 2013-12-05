package passwdprompt

import (
	"fmt"
	"os"
	"syscall"
)

// GetUserPass asks the user to enter her password.
func GetPassword(prompt string) (pw string, err error) {
	fmt.Print(prompt)

	h := syscall.Handle(os.Stdin.Fd())
	var m uint32
	if err = syscall.GetConsoleMode(h, &m); err != nil {
		return
	}
	SetConsoleMode(h, m&(^ENABLE_ECHO_INPUT))
	defer func() {
		SetConsoleMode(h, m)
		fmt.Println()
	}()
	_, err = fmt.Scanln(&pw)
	return
}

const ENABLE_ECHO_INPUT uint32 = 0x04

var (
	modkernel32        = syscall.NewLazyDLL("kernel32.dll")
	procSetConsoleMode = modkernel32.NewProc("SetConsoleMode")
)

func SetConsoleMode(console syscall.Handle, mode uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procSetConsoleMode.Addr(), 2, uintptr(console), uintptr(mode), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
