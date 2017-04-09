// +build windows

package terminfo

// Unfortunatly any other OS will just return false here. Maybe at some point we can
// propertly detect Windows cygwin terminals, but for now we'll stick with false.

// IsFDTTY returns true if the given file descriptor is a terminal.
func IsFDTTY(fd uintptr) bool {
	return false
}
