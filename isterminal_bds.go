// +build darwin dragonfly freebsd netbsd openbsd

package terminfo

import (
	"syscall"
	"unsafe"
)

// IsFDTerminal returns true if the given file descriptor is a terminal.
// This is borrowed from golang/crypto/blob/master/ssh/terminal/util.go
func IsFDTerminal(fd uintptr) bool {
	var termios syscall.Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, fd, syscall.TIOCGETA, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
}
