// +build windows

package terminfo

// Unfortunatly any other OS will just return false here. Maybe at some point we can
// propertly detect Windows cygwin terminals, but for now we'll stick with false.

// IsFDTerminal returns true if the given file descriptor is a terminal.
func IsFDTerminal(fd uintptr) bool {
	return false
}
