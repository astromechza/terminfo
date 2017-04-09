package main

import (
	"fmt"

	"github.com/AstromechZA/terminfo"
)

func main() {
	fmt.Printf("Is stdout a tty? %t\n", terminfo.IsStdoutTerminal())
	fmt.Printf("Is stderr a tty? %t\n", terminfo.IsStderrTerminal())

	w, h, _ := terminfo.GetStdoutDimensions()
	w, h = terminfo.DimensionsOrDefault(w, h, 80, 30)
	fmt.Printf("Dimensions of stdout: %d x %d\n", w, h)

	w, h = terminfo.DimensionsOrMin(w, h, 80, 60)
	fmt.Printf("Rounded dimensions: %d x %d\n", w, h)
}
