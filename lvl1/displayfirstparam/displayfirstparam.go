package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		return
	}
	for _, chr := range args[1] {
		z01.PrintRune(chr)
	} 
	z01.PrintRune('\n')
}