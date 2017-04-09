package terminfo

import "os"

// IsFileTerminal returns true if the given file's descriptor is a terminal
func IsFileTerminal(f *os.File) bool {
	return IsFDTerminal(f.Fd())
}

// IsStdoutTerminal returns true if os.Stdout is a terminal
func IsStdoutTerminal() bool {
	return IsFileTerminal(os.Stdout)
}

// IsStderrTerminal returns true if os.Stderr is a terminal
func IsStderrTerminal() bool {
	return IsFileTerminal(os.Stderr)
}

// GetFileDimensions returns the dimensions of the given file by looking at its
// file descriptor.
func GetFileDimensions(f *os.File) (width, height int, err error) {
	return GetFDDimensions(f.Fd())
}

// GetStdoutDimensions returns the dimensions of the stdout
func GetStdoutDimensions() (width, height int, err error) {
	return GetFileDimensions(os.Stdout)
}

// GetStderrDimensions returns the dimensions of the stderr
func GetStderrDimensions() (width, height int, err error) {
	return GetFileDimensions(os.Stderr)
}

// DimensionsOrDefault is a helper method to use default dimensions rather than the output dimensions
// in cases where there was an error reading the dimensions.
func DimensionsOrDefault(twidth, theight, defaultwidth, defaultheight int) (width, height int) {
	width, height = twidth, theight
	if width <= 0 {
		width = defaultwidth
	}
	if height <= 0 {
		height = defaultheight
	}
	return
}

// DimensionsOrMin is a helper method to use minimum dimensions rather than the output dimensions
// in cases where there was an error or the dimensions were smaller than expected.
func DimensionsOrMin(twidth, theight, minwidth, minheight int) (width, height int) {
	width, height = twidth, theight
	if width < minwidth {
		width = minwidth
	}
	if height < minheight {
		height = minheight
	}
	return
}
