package terminfo

import "os"

// IsFileTTY returns true if the given file's descriptor is a terminal
func IsFileTTY(f *os.File) bool {
	return IsFDTTY(f.Fd())
}

// IsStdoutTTY returns true if os.Stdout is a terminal
func IsStdoutTTY() bool {
	return IsFileTTY(os.Stdout)
}

// IsStderrTTY returns true if os.Stderr is a terminal
func IsStderrTTY() bool {
	return IsFileTTY(os.Stderr)
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
