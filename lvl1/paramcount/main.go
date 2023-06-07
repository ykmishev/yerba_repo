package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := strconv.Itoa(len(os.Args[1:]))
	for _, chr := range args {
		z01.PrintRune(chr)
	}
	z01.PrintRune('\n')

}