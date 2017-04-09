// +build linux android

package terminfo

import (
	"syscall"
	"unsafe"
)

// IsFDTTY returns true if the given file descriptor is a terminal.
// This is borrowed from golang/crypto/blob/master/ssh/terminal/util.go
func IsFDTTY(fd uintptr) bool {
	var termios syscall.Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, fd, 0x5401, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
}
