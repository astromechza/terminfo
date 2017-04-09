// +build windows

package terminfo

import (
	"fmt"
)

// GetFDDimensions returns the dimensions of the given terminal.
func GetFDDimensions(fd uintptr) (width, height int, err error) {
	return -1, -1, fmt.Errorf("not implemented")
}
