package main

import (
	"fmt"

	"github.com/AstromechZA/terminfo"
)

func main() {
	fmt.Printf("Is stdout a tty? %t\n", terminfo.IsStdoutTTY())
	fmt.Printf("Is stderr a tty? %t\n", terminfo.IsStdoutTTY())
	w, h, _ := terminfo.GetStderrDimensions()
	fmt.Printf("Dimensions of stdout: %d x %d\n", w, h)
}
